package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBService struct{}

var MongoClient *mongo.Client

func (m *MongoDBService) Connect(ctx context.Context, uri string) (interface{}, error) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	MongoClient = client
	return client, nil
}

func (m *MongoDBService) Disconnect(ctx context.Context) error {
	return MongoClient.Disconnect(ctx)
}
