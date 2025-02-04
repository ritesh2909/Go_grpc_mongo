package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func Connect() (*mongo.Client, error) {

	clientOptions := options.Client().ApplyURI("mongodb+srv://ritesh:ritesh@cluster0.atsdgyg.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0")

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	MongoClient = client
	return client, nil
}

func Disconnect() error {
	return MongoClient.Disconnect(context.TODO())
}

func GetCollection(collectionName string) *mongo.Collection {
	return MongoClient.Database("Crud").Collection(collectionName)
}
