package handlers

import (
	"module/placeholder/internal/ui"
	"net/http"

	"github.com/a-h/templ"
)

func PageIndex() http.Handler {
	return templ.Handler(ui.Index())
}
