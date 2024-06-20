package handlers

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"golang.org/x/oauth2"
)

func Authorize(oAuthConfig OAuth2Mock) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		url := oAuthConfig.AuthCodeURL()
		http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	}
}

func OAuth2MockAuthorize() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u, err := url.Parse("http://localhost/callback")
		if err != nil {
			fmt.Println(err)
		}

		v := url.Values{}
		v.Set("code", "code")
		u.RawQuery = v.Encode()

		http.Redirect(w, r, u.String(), http.StatusTemporaryRedirect)
	}
}

func Callback(oAuthConfig OAuth2Mock) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u, err := url.Parse("http://localhost/callback")
		if err != nil {
			fmt.Println(err)
		}

		v := url.Values{}
		v.Set("code", "code")
		u.RawQuery = v.Encode()

		http.Redirect(w, r, u.String(), http.StatusTemporaryRedirect)
	}
}

type OAuth2Mock struct{}

func (o *OAuth2Mock) AuthCodeURL() string {
	u := url.URL{
		Scheme: "http",
		Host:   "localhost",
		Path:   "login/oauth/authorize",
	}

	return u.String()
}

func (o *OAuth2Mock) Exchange(ctx context.Context, code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error) {
	return &oauth2.Token{
		AccessToken: "AccessToken",
		Expiry:      time.Now().Add(1 * time.Hour),
	}, nil
}

var oauthConfig = &oauth2.Config{
	ClientID:     "your own client id here",
	ClientSecret: "your own client secret here",
	Endpoint: oauth2.Endpoint{
		AuthURL:  "/auth",
		TokenURL: "/token",
	},
	RedirectURL: "",
	Scopes:      []string{"user:email"},
}
