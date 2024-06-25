package server

import "module/placeholder/internal/server/handlers"

func (s *Server) Routes() {
	s.r.Handle("GET /", handlers.Index())

	s.srv.Handler = s.middlewares(s.r)
}
