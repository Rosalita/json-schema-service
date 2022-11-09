package main

import (
	"context"
)

// mockClient is a mock implementation of a mongo client
// which satisfies the same interface as the mongo client wrapper.
// mockClient is used in unit tests to avoid dependency on a
// real mongo database.
type mockClient struct{}

func (mc mockClient) Connect(ctx context.Context) error {
	return nil
}
func (mc mockClient) Database(dbName string) DatabaseIface {
	return mockDatabase{}
}

type mockDatabase struct{}

func (md mockDatabase) Collection(colName string) CollectionIface {
	return mockCollection{}
}

type mockCollection struct{}

func (mc mockCollection) InsertOne(ctx context.Context, document interface{}) (interface{}, error) {
	return nil, nil
}

// newMockDbClient is a helper function that returns a new mock db client
func newMockDbClient() mockClient {
	return mockClient{}
}
