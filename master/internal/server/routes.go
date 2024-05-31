package server

import "module/placeholder/internal/server/handlers"

func (s *Server) Routes() {
	s.r.Use(s.middlewares)
	s.r.Get("/", handlers.Index())
}
