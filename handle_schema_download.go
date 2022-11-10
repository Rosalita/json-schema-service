package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *server) handleSchemaDownload() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		schemaID := vars["id"]

		database := s.db.Database("validation_service")
		collection := database.Collection("schemas")

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var result schemaData

		filter := bson.D{{Key: "schema_id", Value: schemaID}}

		err := collection.FindOne(ctx, filter).Decode(&result)
		if err != nil {
			log.Println(err)
		}

		if result.Schema == "" {
			http.Error(w, "not found", http.StatusNotFound)
		}

		fmt.Fprintf(w, result.Schema)
	}
}
