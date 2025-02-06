package twitter

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/cockroachdb/errors"
	"github.com/go-resty/resty/v2"
	"github.com/wheatsliu/reng/pkg/logger"
)

type Account struct {
	UserName   string
	Email      string
	Password   string
	PrivateKey string
	Cookies    map[string]*http.Cookie
}

func (a *Account) isAccountLogin() bool {
	return len(a.Cookies) > 0
}

func (a *Account) GetRequestCookie() string {
	cookies := make([]string, len(a.Cookies))
	for name, cookie := range a.Cookies {
		format := "%s=%s"
		if cookie.Quoted {
			format = "%s=\"%s\""
		}
		cookieItem := fmt.Sprintf(format, name, cookie.Value)
		cookies = append(cookies, cookieItem)
	}
	reqCookie := strings.Join(cookies, "; ")
	return reqCookie
}

func (a *Account) Login(httpCli *resty.Client) error {
	loginProcess := NewLoginProcess(a)
	loginProcess.AddFlow("CreateGetGuestToken", createGuestToken())
	loginProcess.AddFlow("InitLoginProgress", startLogin())
	loginProcess.AddFlow("LoginJsInstrumentationSubtask", jsInstrumentation())
	loginProcess.AddFlow("LoginEnterUserIdentifierSSO", enterUserIdentifier())
	loginProcess.AddFlow("LoginEnterPassword", enterPassword())
	loginProcess.AddFlow("LoginTwoFactorAuthChallenge", twoFactorAuthChallenge())
	loginProcess.AddFlow("UpdateCsrfToken", updateCsrfToken())

	currentFlow := loginProcess.FirstFlow
	for currentFlow != nil {
		loginProcess.CurrentFlow = currentFlow
		err := currentFlow.Handler(loginProcess)
		if err != nil {
			return errors.Wrapf(err, "Login flow failed, flow: %s", currentFlow.SubtaskID)
		}
		currentFlow = currentFlow.Next
	}

	logger.Info("Login success", logger.Any("headers", httpCli.Header), logger.Any("cookies", httpCli.Cookies))
	return nil
}
