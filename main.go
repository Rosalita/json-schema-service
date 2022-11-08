package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func setupDatabase() (string, error) {
	return "db", nil
}

func run() error {
	r := mux.NewRouter()

	db, err := setupDatabase()
	if err != nil {
		return fmt.Errorf("setup database error: %w", err)
	}

	server := newServer(db, r)

	http.ListenAndServe(":8080", server)
	return nil
}
