package main

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/santhosh-tekuri/jsonschema"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *server) handleSchemaValidate() http.HandlerFunc {

	type response struct {
		Action  string `json:"action"`
		ID      string `json:"id"`
		Status  string `json:"status"`
		Message string `json:"message,omitempty"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		schemaID := vars["id"]

		var requestData map[string]interface{}
		s.decode(w, r, &requestData)

		removeNullValues(requestData)

		payload, err := json.Marshal(requestData)
		if err != nil {
			http.Error(w, "json marshal error", http.StatusInternalServerError)
			return
		}

		database := s.db.Database("validation_service")
		collection := database.Collection("schemas")

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var result schemaData

		filter := bson.D{{Key: "schema_id", Value: schemaID}}

		err = collection.FindOne(ctx, filter).Decode(&result)
		if err == mongo.ErrNoDocuments {
			http.Error(w, "schema not found", http.StatusNotFound)
			return
		}
		if err != nil {
			http.Error(w, "database error", http.StatusInternalServerError)
			return
		}

		compiler := jsonschema.NewCompiler()
		if err := compiler.AddResource(schemaID, strings.NewReader(result.Schema)); err != nil {
			http.Error(w, "internal error", http.StatusInternalServerError)
			return
		}

		schema, err := compiler.Compile(schemaID)
		if err != nil {
			http.Error(w, "internal error", http.StatusInternalServerError)
			return
		}

		reader := bytes.NewReader(payload)
		if err = schema.Validate(reader); err != nil {

			resp := response{
				Action:  "validateDocument",
				ID:      schemaID,
				Status:  "error",
				Message: err.Error(),
			}

			if err := s.respond(w, r, resp, http.StatusOK); err != nil {
				http.Error(w, "internal error", http.StatusInternalServerError)
				return
			}
			return
		}

		resp := response{
			Action: "validateDocument",
			ID:     schemaID,
			Status: "success",
		}

		if err := s.respond(w, r, resp, http.StatusOK); err != nil {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}
	}
}

// removeNullValues is a helper function that removes null values from a map
func removeNullValues(m map[string]interface{}) {
	val := reflect.ValueOf(m)
	for _, e := range val.MapKeys() {
		v := val.MapIndex(e)
		if v.IsNil() {
			delete(m, e.String())
			continue
		}
		switch t := v.Interface().(type) {
		case map[string]interface{}:
			removeNullValues(t)
		}
	}
}
