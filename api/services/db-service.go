package services

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DBConnection mongo connection
var DBConnection *mongo.Client

// DB mongo database
var DB *mongo.Database

// MongoConnect creates the new client
func MongoConnect(c context.Context, db string) error {
	uri := "mongodb://localhost:27017"
	ctx, cancel := context.WithTimeout(c, 100*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}
	DBConnection = client
	DB = client.Database(db)
	return nil
}

// MongoConnectionHealth connection health check
func MongoConnectionHealth(c context.Context) error {
	ctx, cancel := context.WithTimeout(c, 2*time.Second)
	defer cancel()
	if err := DBConnection.Ping(ctx, nil); err != nil {
		return err
	}
	return nil
}

// Collection defines the name of the constructor
type Collection struct {
	collection *mongo.Collection
}

// FindOne - return a single document
func (c Collection) FindOne(ctx context.Context, filter interface{}, opt *options.FindOneOptions) (*mongo.SingleResult, error) {
	result := c.collection.FindOne(ctx, filter, opt)
	if err := result.Err(); err != nil {
		return nil, err
	}
	return result, nil
}

// Find - returns an array of document
func (c Collection) Find(ctx context.Context, filter interface{}, opt *options.FindOptions) (*mongo.Cursor, error) {
	result, err := c.collection.Find(ctx, filter, opt)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateByID - update and returns the updated document
func (c Collection) UpdateByID(ctx context.Context, id interface{}, opt *options.UpdateOptions) (*mongo.UpdateResult, error) {
	result, err := c.collection.UpdateByID(ctx, id, opt)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Update - update and returns the updated documents
func (c Collection) Update(ctx context.Context, filter interface{}, update interface{}, opt *options.UpdateOptions) (*mongo.UpdateResult, error) {
	result, err := c.collection.UpdateMany(ctx, filter, update, opt)
	if err != nil {
		return nil, err
	}
	return result, nil
}
