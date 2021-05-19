package constants

// MongoErrors - list of all the errors related to mongo db.
var AAuthErrors = map[string]interface{}{
	"ModuleName": "AUTHENTICATOR",
	"TokenNotFound": map[string]interface{}{
		"statusCode": 401,
		"message":    "Authentication token not found",
	},
}
