package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// server holds all server dependencies.
type server struct {
	db     string
	router *mux.Router
}

// ServeHTTP ensures that server satisfies the http.Handler interface.
func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

// newServer is a constructor for a server that sets up routes.
func newServer(db string, router *mux.Router) *server {
	s := &server{
		db:     db,
		router: router,
	}
	s.routes()
	return s
}
