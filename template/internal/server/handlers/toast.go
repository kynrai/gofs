package handlers

import (
	"net/http"

	"module/placeholder/internal/ui"

	"github.com/a-h/templ"
)

func ToastSuccess(w http.ResponseWriter, r *http.Request, msg string) {
	templ.Handler(ui.ToastSuccess(msg)).ServeHTTP(w, r)
}

func ToastInfo(w http.ResponseWriter, r *http.Request, msg string) {
	templ.Handler(ui.ToastInfo(msg)).ServeHTTP(w, r)
}

func ToastWarning(w http.ResponseWriter, r *http.Request, msg string) {
	templ.Handler(ui.ToastWarning(msg)).ServeHTTP(w, r)
}

func ToastError(w http.ResponseWriter, r *http.Request, status int, msg string) {
	w.Header().Add("HX-Reswap", "none")
	w.WriteHeader(status)
	templ.Handler(ui.ToastError(msg)).ServeHTTP(w, r)
}

func ToastSuccessDemo() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ToastSuccess(w, r, "Success!")
	})
}

func ToastInfoDemo() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ToastInfo(w, r, "Info!")
	})
}

func ToastWarningDemo() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ToastWarning(w, r, "Warning!")
	})
}

func ToastErrorDemo() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ToastError(w, r, http.StatusInternalServerError, "Error!")
	})
}
