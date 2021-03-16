export Env=local
export Port=8080
export BodyLimit=5mb
export AllowCorsOrigin=*
export AllowCorsMethods=OPTIONS,GET,POST,PUT,PATCH,DELETE

export Domain=https://a.com

export MongoDbName=url_shortner
export MongoHost=127.0.0.1
export MongoPort=27017
export MongoUsername=
export MongoPassword=
export IsMongoCredentials=true

go run main.go
