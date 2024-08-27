package server

import (
	"net/http"

	"module/placeholder/internal/server/assets"
	"module/placeholder/internal/server/handlers"
	"module/placeholder/internal/server/handlers/page"
	"module/placeholder/internal/server/logging"
)

func (s *Server) Routes() {
	// filserver route for assets
	assetMux := http.NewServeMux()
	assetMux.Handle("GET /{path...}", http.StripPrefix("/assets/", handlers.NewHashedAssets(assets.FS)))
	s.r.Handle("GET /assets/{path...}", s.assetsMiddlewares(assetMux))

	// handlers for normal routes with all general middleware
	routesMux := http.NewServeMux()
	routesMux.Handle("GET /", page.Index())
	routesMux.Handle("GET /modal", handlers.ModalDemo())

	routesMux.Handle("GET /hello", http.HandlerFunc(hello))

	routesMux.Handle("GET /toast-success", handlers.ToastSuccessDemo())
	routesMux.Handle("GET /toast-info", handlers.ToastInfoDemo())
	routesMux.Handle("GET /toast-warning", handlers.ToastWarningDemo())
	routesMux.Handle("GET /toast-error", handlers.ToastErrorDemo())

	s.r.Handle("GET /", s.routeMiddlewares(routesMux))

	s.srv.Handler = s.r
}

func hello(w http.ResponseWriter, r *http.Request) {
	defer logging.Trace(r.Context(), "hello").End()
	w.Write([]byte("Hello, World!"))
}
