package main

import (
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func (s *server) handleSchemaUpload() http.HandlerFunc {

	type response struct {
		Action string `json:"action"`
		ID     string `json:"id"`
		Status string `json:"status"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		schemaID := vars["id"]

		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "error reading body", http.StatusBadRequest)
			return
		}

		schema := string(body)
		log.Println(schema)

		resp := response{
			Action: "uploadSchema",
			ID:     schemaID,
			Status: "success",
		}

		if err := s.respond(w, r, resp, http.StatusCreated); err != nil {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}
	}
}
