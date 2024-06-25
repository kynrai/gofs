package server

import (
	"io/fs"
	"module/placeholder/internal/server/assets"
	"module/placeholder/internal/server/handlers"
	"net/http"
)

func (s *Server) Routes() {
	s.r.Handle("GET /assets/*", http.StripPrefix("/assets/", http.FileServerFS(fs.FS(assets.FS))))

	s.r.Handle("GET /", handlers.PageIndex())

	s.srv.Handler = s.middlewares(s.r)
}
