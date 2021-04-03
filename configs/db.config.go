package configs

// DBConfig config.
type DBConfigStruct struct {
	MongoDbName              string
	MongoHost                string
	MongoPort                string
	MongoUsername            string
	MongoPassword            string
	MongoReplicate           string
	MongoReadPreference      string
	MongoPemPath             string
	MongoServerIdentityCheck string
	IsMongoCredentials       bool
}

var DBConfig DBConfigStruct
