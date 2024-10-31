package db

import (
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"sipamit-be/internal/config"
)

func Connect() *mongo.Database {
	client, err := mongo.Connect(options.Client().ApplyURI(config.Mongo.Url))
	if err != nil {
		panic(err)
	}
	return client.Database(config.Mongo.Name)
}
