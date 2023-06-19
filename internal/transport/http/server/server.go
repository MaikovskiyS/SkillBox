package server

import (
	"log"
	"net/http"
	"time"
)

const (
	_defaultReadTimeout     = 5 * time.Second
	_defaultWriteTimeout    = 5 * time.Second
	_defaultAddr            = ":8080"
	_defaultShutdownTimeout = 3 * time.Second
)

// Server -.
type Server struct {
	server          *http.Server
	shutdownTimeout time.Duration
}

// New -.
func New(handler http.Handler) *Server {
	httpServer := &http.Server{
		Handler:      handler,
		ReadTimeout:  _defaultReadTimeout,
		WriteTimeout: _defaultWriteTimeout,
		Addr:         _defaultAddr,
	}

	s := &Server{
		server:          httpServer,
		shutdownTimeout: _defaultShutdownTimeout,
	}
	return s
}

func (s *Server) Start() {
	log.Println("Server starting")
	s.server.ListenAndServe()
}
