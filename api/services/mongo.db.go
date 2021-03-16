package services

import (
	"context"
	"dwarf/configs"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DBConnection mongo connection
var DBConnection *mongo.Client

// DB mongo database
var DB *mongo.Database

// MongoConnect establish connection between application and mongo
func MongoConnect(c context.Context) error {

	var mongoConfig configs.DBConfig
	configs.Initialize(&mongoConfig)

	uri := "mongodb://" + mongoConfig.MongoHost + ":" + mongoConfig.MongoPort
	ctx, cancel := context.WithTimeout(c, 100*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}
	DBConnection = client
	DB = client.Database(mongoConfig.MongoDbName)
	return nil
}

// MongoConnectionHealth
func MongoConnectionHealth(c context.Context) error {
	ctx, cancel := context.WithTimeout(c, 2*time.Second)
	defer cancel()
	if err := DBConnection.Ping(ctx, nil); err != nil {
		return err
	}
	return nil
}

// GetCollection - return pts to the collection
func GetCollection(c context.Context, n string) (*mongo.Collection, error) {
	err := MongoConnectionHealth(c)
	if err != nil {
		return nil, err
	}
	return DB.Collection(n), nil
}
