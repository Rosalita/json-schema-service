package main

import (
	"fmt"
	"net/http"
)

func (s *server) handleSchemaDownload() http.HandlerFunc {
	type request struct {
		// TODO string
	}

	type response struct {
		//TODO string `json:"todo`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "download handler")
	}
}