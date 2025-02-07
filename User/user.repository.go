package user

import (
	"log"
	"os"
	"user_crud/db"
	"user_crud/models"
)

type UserRepository interface {
	UserExists(email string) (bool, error)
	CreateUser(userReq models.User) error
	ValidateUser(email string, password string) (string, error)
	GetUserInfo(id string) (*models.GetUserInfoResponse, error)
}

func NewUserRepository() UserRepository {
	dbType := os.Getenv("DB_TYPE")

	switch dbType {
	case "mongo":
		return NewMongoUserRepository(db.MongoClient) // @ritesh handle more dbs
	default:
		log.Fatalf("Unsupported DBTYPE: %s", dbType)
		return nil
	}
}
