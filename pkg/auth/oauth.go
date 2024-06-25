package auth

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"net/http"
)

var oAuthConfig = &oauth2.Config{
	RedirectURL: "http://127.0.0.1:8080/oauth2/callback",
	ClientID: "",
	ClientSecret: "",
	Scopes: []string{"https://www.googleapis.com/auth/userinfo.profile"},
	Endpoint: google.Endpoint,
}

func GetOAuthURL(state string) string {
	return oAuthConfig.AuthCodeURL(state)
}

func GetUser(code string) (*oauth2.Token, error) {
	return oAuthConfig.Exchange(oauth2.NoContext, code)
}

func GetUserInfo(token *oauth2.Token) (*http.Response, error) {
	client := oAuthConfig.Client(oauth2.NoContext, token)
	return client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
}