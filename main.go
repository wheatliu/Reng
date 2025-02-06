package main

import (
	"net/http"

	"github.com/wheatsliu/reng/pkg/logger"
	"github.com/wheatsliu/reng/pkg/twitter"
)

var defaultCookies = map[string]*http.Cookie{}

var defaultAccount = &twitter.Account{
	UserName:   "xxxxxxx",
	Password:   "xxxxxxx",
	Email:      "",
	PrivateKey: "xxxxxxx",
	Cookies:    defaultCookies,
}

func main() {
	cli, err := twitter.NewTwitterClient(defaultAccount)
	if err != nil {
		panic(err)
	}
	query := &twitter.TimeLineReqVariables{
		UserID: "172200213",
	}
	resp, err := cli.GetUserTweets(query)
	if err != nil {
		logger.Error("GetUserTweets", logger.Errorv(err))
	}
	logger.Info("GetUserTweets", logger.Any("resp", resp))
}
