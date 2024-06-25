package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"module/placeholder/config"
)

type Server struct {
	r    *http.ServeMux
	srv  *http.Server
	conf config.Config
}

func New(conf config.Config) (*Server, error) {
	s := new(Server)
	s.conf = conf
	s.r = http.NewServeMux()
	s.srv = &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         fmt.Sprintf("%s:%s", conf.Host, conf.Port),
		Handler:      s.r,
	}
	return s, nil
}

func (s *Server) ListenAndServe() error {
	s.Routes()
	// address for use when testing cookies locally
	if s.conf.Host == "0.0.0.0" {
		log.Printf("server: listening on http://localhost:%s", s.conf.Port)
	} else {
		log.Printf("server: listening on http://%s", s.srv.Addr)
	}
	return s.srv.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}
