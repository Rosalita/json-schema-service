package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func setupDatabase() (*mongo.Client, func(), error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return nil, nil, err
	}

	cleanup := func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}
	return client, cleanup, nil
}

func run() error {
	r := mux.NewRouter()

	db, dbCleanup, err := setupDatabase()
	if err != nil {
		return fmt.Errorf("setup database error: %w", err)
	}
	defer dbCleanup()

	server := newServer(db, r)

	http.ListenAndServe(":8080", server)
	return nil
}
