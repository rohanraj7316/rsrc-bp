package models

import (
	"dwarf/api/helpers"
	"dwarf/api/schemas"
	"dwarf/configs"
	"math/rand"
	"time"
)

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
func CreateURL(rBody map[string]interface{}) (result interface{}, err error) {
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
	return d, nil
}

// FindURLByID - find url data from the hash id
func FindURLByID() {}

// UpdateURLByID - update url data from hash id
func UpdateURLByID() {}
