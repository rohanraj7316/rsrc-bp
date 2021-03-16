package models

import (
	"context"
	"dwarf/api/helpers"
	"dwarf/api/schemas"
	"dwarf/api/services"
	"dwarf/configs"
	"math/rand"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// getHash - genrate random hash.
func getHash(url string, n int) (hashURL string, hash string, err error) {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

	var hashConfig = configs.HashConfig{}
	configs.Initialize(&hashConfig)

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	hash = string(b)
	hashURL = hashConfig.Domain + "/" + hash
	return hashURL, hash, nil
}

// CreateURL create shorter url
func CreateURL(c context.Context, rBody map[string]interface{}) (result interface{}, err error) {
	var d schemas.URLData

	d.URLOriginal = rBody["url"].(string)
	d.CreatedAt = time.Now()

	d.ExpireAt, err = helpers.DateFormatter(rBody["expireAt"].(string), time.RFC3339)
	if err != nil {
		return nil, err
	}

	d.URLHash, d.Hash, err = getHash(d.URLOriginal, 5)
	if err != nil {
		return nil, err
	}

	collection, err := services.GetCollection(c, schemas.URLDataCollectionName)
	if err != nil {
		return nil, err
	}

	result, err = collection.InsertOne(c, d)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// FindURLByID - find url data from the hash id
func FindURLByID(c context.Context, id string) (interface{}, error) {

	collection, err := services.GetCollection(c, schemas.URLDataCollectionName)
	if err != nil {
		return nil, err
	}

	oId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var result bson.M

	err = collection.FindOne(c, bson.M{"_id": oId}).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// UpdateURLByID - update url data from hash id
func UpdateURLByID(c context.Context, id string, rBody map[string]interface{}) (interface{}, error) {

	collection, err := services.GetCollection(c, schemas.URLDataCollectionName)
	if err != nil {
		return nil, err
	}

	oId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": oId}

	update := bson.M{
		"$set": rBody,
	}

	var result bson.M

	upsert := true
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}

	err = collection.FindOneAndUpdate(c, filter, update, &opt).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
