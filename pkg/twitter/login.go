package twitter

import (
	"fmt"
	"net/http"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/go-resty/resty/v2"
	"github.com/pquerna/otp/totp"
	"github.com/wheatsliu/reng/pkg/logger"
)

const (
	RequestTimeout    = 30 * time.Second
	GuestActivateURL  = "https://api.x.com/1.1/guest/activate.json"
	OnboardingTaskURL = "https://api.x.com/1.1/onboarding/task.json"
)

type UserInfo struct {
	ID                  int64   `json:"id"`
	IDStr               string  `json:"id_str"`
	Name                string  `json:"name"`
	ScreenName          string  `json:"screen_name"`
	Location            string  `json:"location"`
	URL                 *string `json:"url"`
	Description         string  `json:"description"`
	Protected           bool    `json:"protected"`
	DefaultProfile      bool    `json:"default_profile"`
	DefaultProfileImage bool    `json:"default_profile_image"`
}

// Start Login Request / Response Data
type (
	StartLoginRequest struct {
		InputFlowData struct {
			FlowContext struct {
				DebugOverrides map[string]string `json:"debug_overrides"`
				StartLocation  struct {
					Location string `json:"location"`
				} `json:"start_location"`
			} `json:"flow_context"`
		} `json:"input_flow_data"`
		SubtaskVersions map[string]int `json:"subtask_versions"`
	}

	StartLoginResponse struct {
		FlowToken string `json:"flow_token"`
		Status    string `json:"status"`
		Subtasks  []struct {
			SubtaskId         string `json:"subtask_id"`
			JsInstrumentation struct {
				Url       string `json:"url"`
				TimeoutMs int    `json:"timeout_ms"`
				NextLink  struct {
					LinkType string `json:"link_type"`
					LinkId   string `json:"link_id"`
				} `json:"next_link"`
			} `json:"js_instrumentation"`
		} `json:"subtasks"`
	}
)

// Js Instrumentation Request / Response Data
type (
	Subtask struct {
		SubtaskID             string        `json:"subtask_id"`
		SettingsList          *SettingsList `json:"settings_list,omitempty"`
		SubtaskBackNavigation string        `json:"subtask_back_navigation"`
	}

	SettingsList struct {
		Settings []Setting `json:"settings"`
	}

	Setting struct {
		ValueType       string    `json:"value_type"`
		ValueIdentifier string    `json:"value_identifier"`
		ValueData       ValueData `json:"value_data"`
	}

	ValueData struct {
		Button *Button `json:"button,omitempty"`
	}

	Button struct {
		NavigationLink NavigationLink `json:"navigation_link"`
	}

	NavigationLink struct {
		LinkType  string `json:"link_type"`
		LinkID    string `json:"link_id"`
		Label     string `json:"label,omitempty"`
		SubtaskID string `json:"subtask_id,omitempty"`
		URL       string `json:"url,omitempty"`
	}

	JsInstrumentation struct {
		Response string `json:"response"`
		Link     string `json:"link"`
	}

	JsInstrumentationSubtaskInput struct {
		SubtaskId         string             `json:"subtask_id"`
		JsInstrumentation *JsInstrumentation `json:"js_instrumentation"`
	}

	JsInstrumentationRequest struct {
		FlowToken     string                           `json:"flow_token"`
		SubtaskInputs []*JsInstrumentationSubtaskInput `json:"subtask_inputs"`
	}

	JsInstrumentationResponse struct {
		FlowToken string    `json:"flow_token"`
		Status    string    `json:"status"`
		Subtasks  []Subtask `json:"subtasks"`
	}
)

// Enter User Identifier Request / Response Data
type (
	TextData struct {
		Result string `json:"result"`
	}
	SettingResponseData struct {
		TextData *TextData `json:"text_data"`
	}
	SettingResponse struct {
		Key          string               `json:"key"`
		ResponseData *SettingResponseData `json:"response_data"`
	}

	EnterUserIdentifierSettingsList struct {
		SettingResponses []SettingResponse `json:"setting_responses"`
		Link             string            `json:"link"`
	}

	EnterUserIdentifierSubtaskInput struct {
		SubtaskId    string                           `json:"subtask_id"`
		SettingsList *EnterUserIdentifierSettingsList `json:"settings_list"`
	}

	EnterUserIdentifierRequest struct {
		FlowToken     string                            `json:"flow_token"`
		SubtaskInputs []EnterUserIdentifierSubtaskInput `json:"subtask_inputs"`
	}

	EnterUserIdentifierSubTask struct {
		SubtaskId string `json:"subtask_id"`
	}

	EnterUserIdentifierResponse struct {
		FlowToken string                        `json:"flow_token"`
		Status    string                        `json:"status"`
		Subtasks  []*EnterUserIdentifierSubTask `json:"subtasks"`
	}
)

// Enter Password Request / Response Data
type (
	EnterPassword struct {
		Password string `json:"password"`
		Link     string `json:"link"`
	}

	EnterPasswordSubtaskInput struct {
		SubtaskId     string         `json:"subtask_id"`
		EnterPassword *EnterPassword `json:"enter_password"`
	}

	EnterPasswordRequest struct {
		// {"flow_token":"g;173530831161593371:-1735311452571:uriVXkTkwRhIMO50ZlGr8XMf:7","subtask_inputs":[{"subtask_id":"LoginEnterPassword","enter_password":{"password":"Cassava2023@#","link":"next_link"}}]}
		FlowToken     string                       `json:"flow_token"`
		SubtaskInputs []*EnterPasswordSubtaskInput `json:"subtask_inputs"`
	}

	EnterPasswordSubTask struct {
		SubtaskId string    `json:"subtask_id"`
		User      *UserInfo `json:"user,omitempty"`
	}
	EnterPasswordResponse struct {
		FlowToken string                  `json:"flow_token"`
		Status    string                  `json:"status"`
		Subtasks  []*EnterPasswordSubTask `json:"subtasks"`
	}
)

// Two Factor Auth Challenge Request / Response Data
type (
	EnterText struct {
		Text string `json:"text"`
		Link string `json:"link"`
	}

	TwoFactorAuthChallengeSubtaskInput struct {
		SubtaskId string     `json:"subtask_id"`
		EnterText *EnterText `json:"enter_text"`
	}

	TwoFactorAuthChallengeRequest struct {
		FlowToken     string                                `json:"flow_token"`
		SubtaskInputs []*TwoFactorAuthChallengeSubtaskInput `json:"subtask_inputs"`
	}

	TwoFactorAuthChallengeResponse struct {
		FlowToken string     `json:"flow_token"`
		Status    string     `json:"status"`
		Subtasks  []struct{} `json:"subtasks"`
	}

	UpdateCsrfTokenRequest struct {
		FlowToken     string     `json:"flow_token"`
		SubtaskInputs []struct{} `json:"subtask_inputs"`
	}
)

type Flow struct {
	FlowToken string
	SubtaskID string
	Handler   func(p *LoginProgress) error
	Next      *Flow
}

type LoginProgress struct {
	FirstFlow   *Flow
	LastFlow    *Flow
	CurrentFlow *Flow
	httpCli     *resty.Client
	account     *Account
}

func NewLoginProcess(account *Account) *LoginProgress {
	httpCli := resty.New().SetHeaders(CommonRequestHeaders).SetTimeout(RequestTimeout).SetDebug(true)
	return &LoginProgress{
		httpCli: httpCli,
		account: account,
	}
}

func (c *LoginProgress) AddFlow(subtaskID string, f func(p *LoginProgress) error) {
	flow := &Flow{
		SubtaskID: subtaskID,
		Handler:   f,
		Next:      nil,
	}
	if c.FirstFlow == nil {
		c.FirstFlow = flow
		c.LastFlow = flow
	} else {
		c.LastFlow.Next = flow
		c.LastFlow = flow
	}
}

func createGuestToken() func(p *LoginProgress) error {
	return func(p *LoginProgress) error {
		time.Sleep(2 * time.Second)
		result := map[string]string{}
		resp, err := p.httpCli.R().SetResult(&result).Post(GuestActivateURL)
		if err != nil {
			return errors.Wrap(err, "failed to get guest token")
		}
		logger.Debug("get guest token response", logger.Any("status_code", resp.StatusCode()), logger.String("result", resp.String()))

		if resp.StatusCode() != http.StatusOK {
			logger.Error("failed to get guest token,", logger.Any("status_code", resp.StatusCode()), logger.String("result", resp.String()))
			return errors.New("failed to get guest token")
		}

		if token, ok := result["guest_token"]; ok {
			p.httpCli.SetHeader("x-guest-token", token)
			return nil
		}

		logger.Error("guest token not found in response,", logger.Any("result", result))
		return errors.New("guest token not found in response")
	}
}

func startLogin() func(p *LoginProgress) error {
	return func(p *LoginProgress) error {
		time.Sleep(2 * time.Second)
		query := map[string]string{
			"flow_name": "login",
		}
		result := &StartLoginResponse{}
		data := StartLoginRequest{}
		data.WithDefaultValue()
		resp, err := p.httpCli.R().SetResult(result).SetBody(data).SetQueryParams(query).Post(OnboardingTaskURL)
		if err != nil {
			return fmt.Errorf("failed to start login: %w", err)
		}

		logger.Debug("start login response", logger.Any("status_code", resp.StatusCode()), logger.String("result", resp.String()))
		if resp.StatusCode() != http.StatusOK || result.Status != "success" {
			return errors.Newf("failed to start login, status code: %d, result: %s", resp.StatusCode(), resp.String())
		}

		isNextSubtaskIDMatch := false
		subTaskIdList := make([]string, len(result.Subtasks))
		for i, subtask := range result.Subtasks {
			if subtask.SubtaskId == p.CurrentFlow.Next.SubtaskID {
				isNextSubtaskIDMatch = true
			}
			subTaskIdList[i] = subtask.SubtaskId
		}
		if isNextSubtaskIDMatch {
			p.CurrentFlow.Next.FlowToken = result.FlowToken
			return nil
		} else {
			return errors.Newf("subtask id not match, expect: %s, got: %v", p.CurrentFlow.Next.SubtaskID, subTaskIdList)
		}
	}
}

func jsInstrumentation() func(p *LoginProgress) error {
	return func(p *LoginProgress) error {
		time.Sleep(2 * time.Second)
		instrumentation := &JsInstrumentation{
			Response: "{}",
			Link:     "next_link",
		}
		input := &JsInstrumentationSubtaskInput{
			SubtaskId:         p.CurrentFlow.SubtaskID,
			JsInstrumentation: instrumentation,
		}
		data := &JsInstrumentationRequest{
			FlowToken:     p.CurrentFlow.FlowToken,
			SubtaskInputs: []*JsInstrumentationSubtaskInput{input},
		}
		result := &JsInstrumentationResponse{}
		resp, err := p.httpCli.R().SetResult(result).SetBody(data).Post(OnboardingTaskURL)
		if err != nil {
			return errors.Wrap(err, "failed to send js instrumentation request")
		}
		logger.Debug("js instrumentation response", logger.Any("status_code", resp.StatusCode()), logger.String("result", resp.String()))

		if resp.StatusCode() != http.StatusOK || result.Status != "success" {
			return errors.Newf("failed to send js instrumentation, status code: %d, result: %s", resp.StatusCode(), resp.String())
		}

		isNextSubtaskIDMatch := false
		subTaskIdList := make([]string, len(result.Subtasks))

		for i, subtask := range result.Subtasks {
			if subtask.SubtaskID == p.CurrentFlow.Next.SubtaskID {
				isNextSubtaskIDMatch = true
			}
			subTaskIdList[i] = subtask.SubtaskID
		}

		if isNextSubtaskIDMatch {
			if p.CurrentFlow.Next != nil {
				p.CurrentFlow.Next.FlowToken = result.FlowToken
			}
			return nil
		} else {
			return errors.Newf("subtask id not match, expect: %s, got: %v", p.CurrentFlow.Next.SubtaskID, subTaskIdList)
		}
	}
}

func enterUserIdentifier() func(p *LoginProgress) error {
	return func(p *LoginProgress) error {
		time.Sleep(2 * time.Second)
		settingsList := &EnterUserIdentifierSettingsList{
			SettingResponses: []SettingResponse{
				{
					Key: "user_identifier",
					ResponseData: &SettingResponseData{
						TextData: &TextData{
							Result: p.account.UserName,
						},
					},
				},
			},
			Link: "next_link",
		}
		input := &EnterUserIdentifierSubtaskInput{
			SubtaskId:    p.CurrentFlow.SubtaskID,
			SettingsList: settingsList,
		}
		data := &EnterUserIdentifierRequest{
			FlowToken:     p.CurrentFlow.FlowToken,
			SubtaskInputs: []EnterUserIdentifierSubtaskInput{*input},
		}
		result := &EnterUserIdentifierResponse{}
		resp, err := p.httpCli.R().SetResult(result).SetBody(data).Post(OnboardingTaskURL)
		if err != nil {
			return errors.Wrap(err, "failed to enter user identifier")
		}
		logger.Debug("enter user identifier response", logger.Any("status_code", resp.StatusCode()), logger.String("result", resp.String()))

		if resp.StatusCode() != http.StatusOK || result.Status != "success" {
			return errors.Newf("failed to enter user identifier, status code: %d, result: %s", resp.StatusCode(), resp.String())
		}

		matchNextSubtaskID := false
		subTaskIdList := make([]string, len(result.Subtasks))
		for i, subtask := range result.Subtasks {
			if subtask.SubtaskId == p.CurrentFlow.Next.SubtaskID {
				matchNextSubtaskID = true
			}
			subTaskIdList[i] = subtask.SubtaskId
		}
		if matchNextSubtaskID {
			if p.CurrentFlow.Next != nil {
				p.CurrentFlow.Next.FlowToken = result.FlowToken
			}
			return nil
		} else {
			return errors.Newf("subtask id not match, expect: %s, got: %v", p.CurrentFlow.Next.SubtaskID, subTaskIdList)
		}
	}
}

func enterPassword() func(p *LoginProgress) error {
	return func(p *LoginProgress) error {
		time.Sleep(2 * time.Second)
		password := &EnterPassword{
			Password: p.account.Password,
			Link:     "next_link",
		}
		input := &EnterPasswordSubtaskInput{
			SubtaskId:     p.CurrentFlow.SubtaskID,
			EnterPassword: password,
		}
		data := &EnterPasswordRequest{
			FlowToken:     p.CurrentFlow.FlowToken,
			SubtaskInputs: []*EnterPasswordSubtaskInput{input},
		}
		result := &EnterPasswordResponse{}
		resp, err := p.httpCli.R().SetResult(result).SetBody(data).Post(OnboardingTaskURL)
		if err != nil {
			return errors.Wrap(err, "failed to enter password")
		}
		logger.Debug("enter password response", logger.Any("status_code", resp.StatusCode()), logger.String("result", resp.String()))

		if resp.StatusCode() != http.StatusOK || result.Status != "success" {
			return errors.Newf("failed to enter password, status code: %d, result: %s", resp.StatusCode(), resp.String())
		}

		matchNextSubtaskID := false
		subTaskIdList := make([]string, len(result.Subtasks))
		for i, subtask := range result.Subtasks {
			if subtask.SubtaskId == p.CurrentFlow.Next.SubtaskID {
				matchNextSubtaskID = true
			}
			subTaskIdList[i] = subtask.SubtaskId
		}
		if matchNextSubtaskID {
			if p.CurrentFlow.Next != nil {
				p.CurrentFlow.Next.FlowToken = result.FlowToken
			}
			return nil
		} else {
			return errors.Newf("subtask id not match, expect: %s, got: %v", p.CurrentFlow.Next.SubtaskID, subTaskIdList)
		}
	}
}

func twoFactorAuthChallenge() func(p *LoginProgress) error {
	return func(p *LoginProgress) error {
		time.Sleep(2 * time.Second)
		code, _ := totp.GenerateCode(p.account.PrivateKey, time.Now())
		input := &TwoFactorAuthChallengeSubtaskInput{
			SubtaskId: p.CurrentFlow.SubtaskID,
			EnterText: &EnterText{
				Text: code,
				Link: "next_link",
			},
		}
		data := &TwoFactorAuthChallengeRequest{
			FlowToken:     p.CurrentFlow.FlowToken,
			SubtaskInputs: []*TwoFactorAuthChallengeSubtaskInput{input},
		}
		result := &TwoFactorAuthChallengeResponse{}
		resp, err := p.httpCli.R().SetBody(data).SetResult(result).Post(OnboardingTaskURL)
		if err != nil {
			return errors.Wrap(err, "failed to send two factor auth challenge")
		}

		logger.Debug("two factor auth challenge response", logger.Any("status_code", resp.StatusCode()), logger.String("result", resp.String()))
		if resp.StatusCode() != http.StatusOK || result.Status != "success" {
			return errors.Newf("failed to send two factor auth challenge, status code: %d, result: %s", resp.StatusCode(), resp.String())
		}

		logger.Debug("Response", logger.Any("headers", resp.RawResponse.Header))
		if p.CurrentFlow.Next != nil {
			p.CurrentFlow.Next.FlowToken = result.FlowToken
		}

		setCookies := resp.Cookies()
		for _, cookie := range setCookies {
			p.account.Cookies[cookie.Name] = cookie
		}
		logger.Debug("finish user login", logger.Any("cookies", p.account.Cookies))
		return nil
	}
}

func updateCsrfToken() func(p *LoginProgress) error {
	return func(p *LoginProgress) error {
		reqCookie := p.account.GetRequestCookie()
		appendHeaders := map[string]string{
			"Cookie":       reqCookie,
			"x-csrf-token": p.account.Cookies["ct0"].Value,
		}
		p.httpCli.SetHeaders(appendHeaders)

		data := &UpdateCsrfTokenRequest{
			FlowToken:     p.CurrentFlow.FlowToken,
			SubtaskInputs: []struct{}{},
		}

		resp, err := p.httpCli.R().SetBody(data).Post(OnboardingTaskURL)
		if err != nil {
			return errors.Wrap(err, "failed to update csrf token")
		}
		logger.Debug("update csrf token response", logger.Any("status_code", resp.StatusCode()), logger.String("result", resp.String()))

		if resp.StatusCode() != http.StatusOK {
			return errors.Newf("failed to update csrf token, status code: %d, result: %s", resp.StatusCode(), resp.String())
		}

		setCookies := resp.Cookies()
		for _, cookie := range setCookies {
			p.account.Cookies[cookie.Name] = cookie
		}
		logger.Debug("finish update csrf token", logger.Any("new cookies", setCookies))
		return nil
	}
}
func (s *StartLoginRequest) WithDefaultValue() {
	s.InputFlowData.FlowContext.DebugOverrides = make(map[string]string)
	s.InputFlowData.FlowContext.StartLocation.Location = "manual_link"
	s.SubtaskVersions = map[string]int{"action_list": 2, "alert_dialog": 1, "app_download_cta": 1, "check_logged_in_account": 1, "choice_selection": 3, "contacts_live_sync_permission_prompt": 0, "cta": 7, "email_verification": 2, "end_flow": 1, "enter_date": 1, "enter_email": 2, "enter_password": 5, "enter_phone": 2, "enter_recaptcha": 1, "enter_text": 5, "enter_username": 2, "generic_urt": 3, "in_app_notification": 1, "interest_picker": 3, "js_instrumentation": 1, "menu_dialog": 1, "notifications_permission_prompt": 2, "open_account": 2, "open_home_timeline": 1, "open_link": 1, "phone_verification": 4, "privacy_options": 1, "security_key": 3, "select_avatar": 4, "select_banner": 2, "settings_list": 7, "show_code": 1, "sign_up": 2, "sign_up_review": 4, "tweet_selection_urt": 1, "update_users": 1, "upload_media": 1, "user_recommendations_list": 4, "user_recommendations_urt": 1, "wait_spinner": 3, "web_modal": 1}
}
