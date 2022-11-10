package main

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"reflect"
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

func (mc mockCollection) FindOne(ctx context.Context, v interface{}) SingleResultIface {
	if val, ok := v.(bson.D); ok {
		m := val.Map()
		id := m["schema_id"]
		if id == "config-schema" {
			return mockSingleResult{}
		}
	}
	return mockSingleResultNotFound{}
}

type mockSingleResult struct{}

func (ms mockSingleResult) Decode(v interface{}) error {
	mockResult := schemaData{Schema: `{"mock":"schema"}`}
	mockVal := reflect.ValueOf(mockResult)
	reflect.ValueOf(v).Elem().Set(mockVal)
	return nil
}

type mockSingleResultNotFound struct{}

func (ms mockSingleResultNotFound) Decode(v interface{}) error {
	mockResult := schemaData{Schema: ""}
	mockVal := reflect.ValueOf(mockResult)
	reflect.ValueOf(v).Elem().Set(mockVal)
	return nil
}

// newMockDbClient is a helper function that returns a new mock db client
func newMockDbClient() mockClient {
	return mockClient{}
}
