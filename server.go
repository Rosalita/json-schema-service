package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

// server holds all server dependencies.
type server struct {
	db     ClientIface
	router *mux.Router
}

// ServeHTTP ensures that server satisfies the http.Handler interface.
func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

// respond is a helper that make response functionality reusable for all requests.
func (s *server) respond(w http.ResponseWriter, r *http.Request, data interface{}, status int) error {
	w.WriteHeader(status)
	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			return err
		}

	}
	return nil
}

// decode is a helper that make decode functionality reusable for all requests
func (s *server) decode(w http.ResponseWriter, r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}

// newServer is a constructor for a server that sets up routes.
func newServer(db ClientIface, router *mux.Router) *server {
	s := &server{
		db:     db,
		router: router,
	}
	s.routes()
	return s
}
