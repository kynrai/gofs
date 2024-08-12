package page

import (
	"module/placeholder/internal/ui"
	"net/http"

	"github.com/a-h/templ"
)

func Index() http.Handler {
	return templ.Handler(ui.Index())
}
