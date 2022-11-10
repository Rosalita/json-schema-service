package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"time"

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

		database := s.db.Database("validation_service")
		collection := database.Collection("schemas")

		document := schemaData{
			ID:     schemaID,
			Schema: schema,
		}
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		_, err = collection.InsertOne(ctx, document)
		if err != nil {
			http.Error(w, "database error", http.StatusInternalServerError)
			return
		}

		log.Println("Inserted a single document")

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
