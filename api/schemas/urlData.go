package schemas

import "time"

// URLDataCollectionName collection name
const URLDataCollectionName = "urlData"

// URLData mongo schema
type URLData struct {
	CreatedAt   time.Time `bson:"createdAt" json:"createdAt" validate:"required"`
	ExpireAt    time.Time `bson:"expireAt" json:"expireAt" validate:"required" default:""`
	Hash        string    `bson:"hash" json:"hash" validate:"required"`
	URLHash     string    `bson:"urlHash" json:"urlHash" validate:"required"`
	URLOriginal string    `bson:"urlOriginal" json:"urlOriginal" validate:"required"`
	IsWebToWeb  bool      `bson:"isWebToWeb" json:"isWebToWeb" validate:"required" default:"false"`
	IsWebToMob  bool      `bson:"isWebToMob" json:"isWebToMob" validate:"required" default:"false"`
	IsMobToApp  bool      `bson:"isMobToApp" json:"isMobToApp" validate:"required" default:"false"`
	IsAppToApp  bool      `bson:"isAppToApp" json:"isAppToApp" validate:"required" default:"false"`
}
