package handlers

import "net/http"

func HandleCallback() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO
		// Get the token
		var token string
		// Return success
		w.Write([]byte("Callback successful: " + token))
		// exchange code for token
		// call /token
		// have token, store in state
		// redirect to home page
	}
}
