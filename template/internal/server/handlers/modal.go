package handlers

import (
	"net/http"

	"module/placeholder/internal/ui/components"

	"github.com/a-h/templ"
)

func ModalDemo() http.Handler {
	return templ.Handler(components.ModalDemo())
}
