package db

import (
	"context"
	"fmt"
	"os"
)

func InitDB() (DBService, error) {
	dbType := os.Getenv("DB_TYPE")

	var dbService DBService

	switch dbType {
	case "mongo":
		dbService = &MongoDBService{}
	case "postgres":
		return nil, fmt.Errorf("Postgres is unimplemented")
	case "redis":
		return nil, fmt.Errorf("Redis is unimplemented")
	default:
		return nil, fmt.Errorf("unsupported DB type: %s", dbType)
	}

	ctx := context.Background()

	if _, err := dbService.Connect(ctx, os.Getenv("MONGO_URI")); err != nil {
		return nil, fmt.Errorf("could not initialize DB: %v", err)
	}

	return dbService, nil
}
