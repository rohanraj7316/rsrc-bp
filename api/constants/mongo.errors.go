package constants

var ModuleName = "MONGO"

// MongoErrors - list of all the errors related to mongo db.
var MongoErrors = map[string]interface{}{
	"DocumentNotFound": map[string]interface{}{
		"statusCode": 404,
		"message":    "unable to find document",
		// "code":       ProductName + "_" + ModuleName,
	},
	"MongoConnection": map[string]interface{}{
		"statusCode": 500,
		"message":    "error in mongo connection",
	},
	"RunTime": map[string]interface{}{
		"statusCode": 500,
		"message":    "error occur in run time",
	},
}
