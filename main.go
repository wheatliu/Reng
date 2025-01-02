package main

import (
	"github.com/wheatsliu/reng/pkg/twitter"
)

func main() {
	twitterAccount := &twitter.Account{
		UserName:   "PhilemenW68622",
		Password:   "Cassava2023@#",
		Email:      "",
		PrivateKey: "GCJY4RSIVMVDG47E",
	}
	twitter.Login(twitterAccount)
}

// kdt=RweFe2aNTz3PI76Fwnw7o41TblbLSMAGKK25YHey; Path=/; Domain=x.com; Expires=Mon, 29 Jun 2026 14:23:35 GMT; Max-Age=47260800; HttpOnly; Secure twid="u=1722978745911267328"; Path=/; Domain=x.com; Expires=Fri, 28 Dec 2029 14:23:35 GMT; Max-Age=157680000; Secure; SameSite=None att=; Path=/; Domain=x.com; Expires=Sun, 29 Dec 2024 14:23:35 GMT; Max-Age=0; HttpOnly; Secure; SameSite=None ct0=5920adcc95389db734dfdf2ebbbb9246; Path=/; Domain=x.com; Expires=Sun, 29 Dec 2024 20:23:35 GMT; Max-Age=21600; Secure auth_token=04a184c956f8d97275e95c00e26f3dd1c4eb11be; Path=/; Domain=x.com; Expires=Fri, 28 Dec 2029 14:23:35 GMT; Max-Age=157680000; HttpOnly; Secure; SameSite=None
