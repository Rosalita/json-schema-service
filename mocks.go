package main

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func newMockDbClient() mockDbClient {
	return mockDbClient{}
}

type mockDbClient struct{}

func (m mockDbClient) Connect(ctx context.Context) error {
	return nil
}

func (m mockDbClient) Disconnect(ctx context.Context) error {
	return nil
}

func (m mockDbClient) Database(name string, opts ...*options.DatabaseOptions) *mongo.Database {
	return nil
}

// dbClient is an interface that both a real mongoDb client and
// a mock client can pass through. This interface allows
// dependencies to be mocked in tests.
type dbClient interface {
	Connect(ctx context.Context) error
	Disconnect(ctx context.Context) error
	Database(name string, opts ...*options.DatabaseOptions) *mongo.Database
}
