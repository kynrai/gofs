package handlers

import (
	"net/http"

	"module/placeholder/internal/ui/components"

	"github.com/a-h/templ"
)

func ToastSuccess(w http.ResponseWriter, r *http.Request, msg string) {
	templ.Handler(components.ToastSuccess(msg)).ServeHTTP(w, r)
}

func ToastInfo(w http.ResponseWriter, r *http.Request, msg string) {
	templ.Handler(components.ToastInfo(msg)).ServeHTTP(w, r)
}

func ToastWarning(w http.ResponseWriter, r *http.Request, msg string) {
	templ.Handler(components.ToastWarning(msg)).ServeHTTP(w, r)
}

func ToastError(w http.ResponseWriter, r *http.Request, status int, msg string) {
	w.Header().Add("HX-Reswap", "none")
	w.WriteHeader(status)
	templ.Handler(components.ToastError(msg)).ServeHTTP(w, r)
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
