package server

import (
	"net/http"

	"module/placeholder/internal/auth"
	"module/placeholder/internal/server/assets"
	"module/placeholder/internal/server/handlers"
	"module/placeholder/internal/server/handlers/page"
	"module/placeholder/internal/server/logging"
)

func (s *Server) Routes() {
	// filserver route for assets
	assetMux := http.NewServeMux()
	assetMux.Handle("GET /*", http.StripPrefix("/assets/", handlers.NewHashedAssets(assets.FS)))
	s.r.Handle("GET /assets/*", s.assetsMiddlewares(assetMux))

	// handlers for normal routes with all general middleware
	routesMux := http.NewServeMux()
	routesMux.Handle("GET /", page.Index())
	routesMux.Handle("GET /user", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := auth.UserFromContext(r.Context())
		if user != nil {
			logging.Write(w, []byte("Hello, "+user.ID+"!"))
			return
		}
		logging.Write(w, []byte("Hello, World!"))
	}))
	s.r.Handle("GET /", s.routeMiddlewares(routesMux))

	s.srv.Handler = s.r
}
