package twitter

import (
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/wheatsliu/reng/pkg/twitter/iface"

	"github.com/pquerna/otp/totp"
)

const (
	RequestTimeout    = 30 * time.Second
	GuestActivateURL  = "https://api.x.com/1.1/guest/activate.json"
	OnboardingTaskURL = "https://api.x.com/1.1/onboarding/task.json"
)

var (
	CommonRequestHeaders = map[string]string{
		"authorization":             "Bearer AAAAAAAAAAAAAAAAAAAAANRILgAAAAAAnNwIzUejRCOuH5E6I8xnZz4puTs%3D1Zv7ttfk8LF81IUq16cHjhLTvJu4FA33AGWWjCpTnA",
		"content-type":              "application/json",
		"x-twitter-active-user":     "yes",
		"x-twitter-client-language": "en",
		"user-agent":                "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36",
	}
)

type Process struct {
	httpClient *resty.Client
	account    *Account
	flow       *Flow
	handlers   []func(ctx *Process)
	index      int8
}

type Account struct {
	UserName   string
	Email      string
	Password   string
	PrivateKey string
}

type Flow struct {
	FlowToken string `json:"flow_token"`
	SubtaskID string `json:"subtask_id"`
}

func NewLoginProcess(account *Account) *Process {
	// cookiejar, _ := cookiejar.New(nil)
	httpClient := resty.New().SetHeaders(CommonRequestHeaders).SetTimeout(RequestTimeout).SetDebug(true)
	// .SetCookieJar(cookiejar)
	return &Process{
		httpClient: httpClient,
		account:    account,
		handlers:   make([]func(ctx *Process), 0),
		index:      0,
		flow:       &Flow{},
	}
}

func (c *Process) Use(f func(ctx *Process)) {
	c.handlers = append(c.handlers, f)
}

func (c *Process) Next() {
	c.index++
	if c.index < int8(len(c.handlers)) {
		c.handlers[c.index](c)
	}
}

func (c *Process) Run() {
	c.handlers[0](c)
}

func Login(account *Account) {
	loginProcess := NewLoginProcess(account)
	loginProcess.Use(initGuestToken())
	loginProcess.Use(startLogin())
	loginProcess.Use(jsInstrumentation())
	loginProcess.Use(enterUserIdentifier())
	loginProcess.Use(enterPassword())
	loginProcess.Use(twoFactorAuthChallenge())
	loginProcess.Run()
}

func initGuestToken() func(p *Process) {
	return func(p *Process) {
		time.Sleep(2 * time.Second)
		result := map[string]string{}
		resp, err := p.httpClient.R().SetResult(&result).Post(GuestActivateURL)
		if err != nil {
			fmt.Println("failed to get guest token:", err)
			return
		}
		if resp.StatusCode() == 200 {
			if token, ok := result["guest_token"]; ok {
				p.httpClient.SetHeader("x-guest-token", token)
				fmt.Println("guest token:", token)
				p.Next()
			} else {
				fmt.Printf("guest token not found in response: %v", result)
				return
			}
		} else {
			fmt.Printf("failed to get guest token, status code: %d, result: %s", resp.StatusCode(), resp.String())
			return
		}
	}
}

func startLogin() func(p *Process) {
	return func(p *Process) {
		time.Sleep(2 * time.Second)
		query := map[string]string{
			"flow_name": "login",
		}
		result := &iface.StartLoginResponse{}
		data := iface.StartLoginRequest{}
		data.WithDefaultValue()
		resp, err := p.httpClient.R().SetResult(result).SetBody(data).SetQueryParams(query).Post(OnboardingTaskURL)
		if err != nil {
			fmt.Println("failed to start login:", err)
			return
		}
		if resp.StatusCode() == 200 {
			if result.Status == "success" {
				p.flow.FlowToken = result.FlowToken
				p.flow.SubtaskID = result.Subtasks[0].SubtaskId
				fmt.Printf("Initiated login flow %s, flow token: %s, subtask id: %s", result.Status, p.flow.FlowToken, p.flow.SubtaskID)
				p.Next()
			} else {
				fmt.Printf("failed to start login, status: %s, result: %v", result.Status, result)
				return
			}
		} else {
			fmt.Printf("failed to start login, status code: %d, result: %s", resp.StatusCode(), resp.String())
			return
		}
	}
}

func jsInstrumentation() func(p *Process) {
	return func(p *Process) {
		time.Sleep(2 * time.Second)
		instrumentation := &iface.JsInstrumentation{
			Response: "{}",
			Link:     "next_link",
		}
		input := &iface.JsInstrumentationSubtaskInput{
			SubtaskId:         p.flow.SubtaskID,
			JsInstrumentation: instrumentation,
		}
		data := &iface.JsInstrumentationRequest{
			FlowToken:     p.flow.FlowToken,
			SubtaskInputs: []*iface.JsInstrumentationSubtaskInput{input},
		}
		result := &iface.JsInstrumentationResponse{}
		resp, err := p.httpClient.R().SetResult(result).SetBody(data).Post(OnboardingTaskURL)
		if err != nil {
			fmt.Println("failed to send js instrumentation:", err)
			return
		}
		if resp.StatusCode() == 200 {
			if result.Status == "success" {
				fmt.Printf("Js instrumentation success, flow token: %s", result.FlowToken)
				p.flow.FlowToken = result.FlowToken
				p.flow.SubtaskID = result.Subtasks[0].SubtaskID
				p.Next()
			} else {
				fmt.Printf("failed to send js instrumentation, status: %s, result: %v", result.Status, result)
				return
			}
		} else {
			fmt.Printf("failed to send js instrumentation, status code: %d, result: %s", resp.StatusCode(), resp.String())
			return
		}
	}
}

func enterUserIdentifier() func(p *Process) {
	return func(p *Process) {
		time.Sleep(2 * time.Second)
		settingsList := &iface.EnterUserIdentifierSettingsList{
			SettingResponses: []iface.SettingResponse{
				{
					Key: "user_identifier",
					ResponseData: &iface.SettingResponseData{
						TextData: &iface.TextData{
							Result: p.account.UserName,
						},
					},
				},
			},
			Link: "next_link",
		}
		input := &iface.EnterUserIdentifierSubtaskInput{
			SubtaskId:    p.flow.SubtaskID,
			SettingsList: settingsList,
		}
		data := &iface.EnterUserIdentifierRequest{
			FlowToken:     p.flow.FlowToken,
			SubtaskInputs: []iface.EnterUserIdentifierSubtaskInput{*input},
		}
		result := &iface.EnterUserIdentifierResponse{}
		resp, err := p.httpClient.R().SetResult(result).SetBody(data).Post(OnboardingTaskURL)
		if err != nil {
			fmt.Println("failed to enter user identifier:", err)
			return
		}
		if resp.StatusCode() == 200 {
			if result.Status == "success" {
				fmt.Printf("Entered user identifier success, flow token: %s", result.FlowToken)
				p.flow.FlowToken = result.FlowToken
				// todo: check if "LoginEnterPassword" subtask is exist
				p.flow.SubtaskID = "LoginEnterPassword"
				p.Next()
			} else {
				fmt.Printf("failed to enter user identifier, status: %s, result: %v", result.Status, result)
				return
			}
		} else {
			fmt.Printf("failed to enter user identifier, status code: %d, result: %s", resp.StatusCode(), resp.String())
			return
		}
	}
}

func enterPassword() func(p *Process) {
	return func(p *Process) {
		time.Sleep(2 * time.Second)
		password := &iface.EnterPassword{
			Password: p.account.Password,
			Link:     "next_link",
		}
		input := &iface.EnterPasswordSubtaskInput{
			SubtaskId:     p.flow.SubtaskID,
			EnterPassword: password,
		}
		data := &iface.EnterPasswordRequest{
			FlowToken:     p.flow.FlowToken,
			SubtaskInputs: []*iface.EnterPasswordSubtaskInput{input},
		}
		result := &iface.EnterPasswordResponse{}
		resp, err := p.httpClient.R().SetResult(result).SetBody(data).Post(OnboardingTaskURL)
		if err != nil {
			fmt.Println("failed to enter password:", err)
			return
		}
		if resp.StatusCode() == 200 {
			if result.Status == "success" {
				fmt.Printf("Entered password success, flow token: %s", result.FlowToken)
				p.flow.FlowToken = result.FlowToken
				p.flow.SubtaskID = "LoginTwoFactorAuthChallenge"
				p.Next()
			} else {
				fmt.Printf("failed to enter password, status: %s, result: %v", result.Status, result)
				return
			}
		} else {
			fmt.Printf("failed to enter password, status code: %d, result: %s", resp.StatusCode(), resp.String())
			return
		}
	}
}

func twoFactorAuthChallenge() func(p *Process) {
	return func(p *Process) {
		time.Sleep(2 * time.Second)
		code, _ := totp.GenerateCode(p.account.PrivateKey, time.Now())
		input := &iface.TwoFactorAuthChallengeSubtaskInput{
			SubtaskId: p.flow.SubtaskID,
			EnterText: &iface.EnterText{
				Text: code,
				Link: "next_link",
			},
		}
		request := &iface.TwoFactorAuthChallengeRequest{
			FlowToken:     p.flow.FlowToken,
			SubtaskInputs: []*iface.TwoFactorAuthChallengeSubtaskInput{input},
		}
		resp, err := p.httpClient.R().SetBody(request).Post(OnboardingTaskURL)

		if err != nil {
			fmt.Println("failed to send two factor auth challenge:", err)
			return
		}

		if resp.StatusCode() == 200 {
			cookies := resp.Cookies()
			fmt.Printf("Two factor auth challenge success, cookies: %v", cookies)
			return
		} else {
			fmt.Printf("failed to send two factor auth challenge, status code: %d, result: %s", resp.StatusCode(), resp.String())
			return
		}
	}
}
