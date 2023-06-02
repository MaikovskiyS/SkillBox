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
	notify          chan error
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
		notify:          make(chan error, 1),
		shutdownTimeout: _defaultShutdownTimeout,
	}
	return s
}

func (s *Server) Start() {
	log.Println("Server starting")
	s.server.ListenAndServe()
}

// // Notify -.
// func (s *Server) Notify() <-chan error {
// 	return s.notify
// }

// // Shutdown -.
// func (s *Server) Shutdown() error {
// 	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
// 	defer cancel()

// 	return s.server.Shutdown(ctx)
//}
