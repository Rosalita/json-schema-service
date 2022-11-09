package main

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

// mongoClient is a wrapper for a mongo client.
// this wrapper returns interfaces instead of concrete types.
// a real mongo client can exist inside this wrapper.
// the wrapper allows creation of a mock client that
// can be used in tests instead of the real wrapped client.
type mongoClient struct {
	cl *mongo.Client
}

func (mc *mongoClient) Database(dbName string) DatabaseIface {
	db := mc.cl.Database(dbName)
	return &mongoDatabase{db: db}
}

func (mc *mongoClient) Connect(ctx context.Context) error {
	err := mc.cl.Connect(ctx)
	return err
}

// mongoDatabase is a wrapper for a *mongo.Database
type mongoDatabase struct {
	db *mongo.Database
}

func (md *mongoDatabase) Collection(colName string) CollectionIface {
	collection := md.db.Collection(colName)
	return &mongoCollection{coll: collection}
}

// mongoCollection is a wrapper for a *mongo.Collection
type mongoCollection struct {
	coll *mongo.Collection
}

func (mc *mongoCollection) InsertOne(ctx context.Context, document interface{}) (interface{}, error) {
	id, err := mc.coll.InsertOne(ctx, document)
	return id.InsertedID, err
}

// ClientIface describes methods on a mongo client
type ClientIface interface {
	Database(string) DatabaseIface
	Connect(ctx context.Context) error
}

// DatabaseIface describes methods on a mongo database
type DatabaseIface interface {
	Collection(name string) CollectionIface
}

// CollectionIface describes methods on a mongo collection
type CollectionIface interface {
	//	FindOne(context.Context, interface{}) SingleResultIface
	InsertOne(context.Context, interface{}) (interface{}, error)
}

// SingleResultIface describes methods on a mongo single result
type SingleResultIface interface {
	Decode(v interface{}) error
}
