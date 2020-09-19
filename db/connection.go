package db

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoConnection : This variable has the connection object and can be used in any file inside the db folder
var MongoConnection = ConnectDB()
var clientOptions = options.Client().ApplyURI("")

// ConnectDB : Function that returns the connection to the database
func ConnectDB() *mongo.Client {
	setConnectionString()
	var clientOptions = options.Client().ApplyURI(os.Getenv("CONNECTION_STRING"))

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("Successful database connection")
	return client
}

// CheckConnection : Function that checks if the database is working at the moment
func CheckConnection() int {
	err := MongoConnection.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}

	return 1
}
