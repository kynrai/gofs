package auth

import (
	"net/http"
	"net/url"

	"golang.org/x/oauth2"
)

var Endpoint = oauth2.Endpoint{
	AuthURL:  "http://localhost:8080/auth",
	TokenURL: "http://localhost:8080/token",
}

type MockOAuthServer struct {
	r *http.ServeMux
}

func (m *MockOAuthServer) handleLoginPage(w http.ResponseWriter, r *http.Request) {
	// Fake a login by posting values to the auth server
	var user string
	var password string

	_, err := http.PostForm("/login", url.Values{
		"user":     []string{user},
		"password": []string{password},
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	// do something with token in resp
}

func (m *MockOAuthServer) handleLogin(w http.ResponseWriter, r *http.Request) {
	// Fake a login by posting values to the auth server
	var user string
	var password string

	_, err := http.PostForm("/login", url.Values{
		"user":     []string{user},
		"password": []string{password},
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	// do something with token in resp
	// redirect to callback
}

func (m *MockOAuthServer) handleAuthorize(w http.ResponseWriter, r *http.Request) {
	// Get the token
	var token string
	// Return success
	w.Write([]byte("Callback successful: " + token))
}

func (m *MockOAuthServer) handleToken(w http.ResponseWriter, r *http.Request) {
	// Get the token
	var token string
	// Return success
	w.Write([]byte("Callback successful: " + token))
}

func (m *MockOAuthServer) Serve() error {
	m.r.HandleFunc("GET /login", m.handleLoginPage) //TODO skip, if authorize -> callback
	m.r.HandleFunc("POST /login", m.handleLogin)    //TODO skip, if authorize -> callback
	m.r.HandleFunc("/authorize", m.handleAuthorize)
	m.r.HandleFunc("/token", m.handleToken)
	return http.ListenAndServe("localhost:8080", m.r)
}
