package handlers

import (
	"net/http"

	"module/placeholder/internal/ui"

	"github.com/a-h/templ"
)

func ModalDemo() http.Handler {
	return templ.Handler(ui.ModalDemo())
}
