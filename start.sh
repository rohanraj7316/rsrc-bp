export Env=local
export Port=8080
export BodyLimit=5mb
export AllowCorsOrigin=*
export AllowCorsMethods=OPTIONS,GET,POST,PUT,PATCH,DELETE
export ProductName=URL_SHORTNER
export ModuleName=MONGO_DB

export Domain=https://a.com

export AAuthURL=https://httpbin.org/get

export MongoDbName=url_shortner
export MongoHost=127.0.0.1
export MongoPort=27017
export MongoUsername=
export MongoPassword=
export IsMongoCredentials=false

go run main.go
