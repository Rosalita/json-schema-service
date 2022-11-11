package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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
		if err == mongo.ErrNoDocuments {
			http.Error(w, "schema not found", http.StatusNotFound)
			return
		}
		if err != nil {
			http.Error(w, "database error", http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, result.Schema)
	}
}
