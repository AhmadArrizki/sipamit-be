package db

import (
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"sipamit-be/internal/config"
)

var Client *mongo.Database

func init() {
	if Client == nil {
		client, err := mongo.Connect(options.Client().ApplyURI(config.Mongo.Url))
		if err != nil {
			panic(err)
		}

		Client = client.Database(config.Mongo.Name)
	}
}
