package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"module/placeholder/config"
	"module/placeholder/internal/db/gormpg"

	"gorm.io/gorm"
)

type Server struct {
	r    *http.ServeMux
	srv  *http.Server
	conf config.Config
	db   *gorm.DB
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
	db, err := s.initdb()
	if err != nil {
		return nil, err
	}
	s.db = db
	return s, nil
}

func (s *Server) initdb() (*gorm.DB, error) {
	switch {
	case s.conf.Env.Local() && s.conf.DSN != "":
		return gormpg.PG(s.conf.DSN)
	case s.conf.Env.Dev() && s.conf.DSN != "":
		return gormpg.CloudSQL(s.conf.DSN)
	case s.conf.Env.Prod() && s.conf.DSN != "":
		return gormpg.CloudSQL(s.conf.DSN)
	default:
		log.Println("server: no database connection")
		return nil, nil
	}
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
