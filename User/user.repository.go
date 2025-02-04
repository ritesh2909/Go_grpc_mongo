package user

import (
	"context"
	"log"
	"user_crud/models"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserRepository struct {
	client *mongo.Client
}

var userRepository *UserRepository

func NewMongoUserRepository(client *mongo.Client) *UserRepository {
	if userRepository == nil {
		userRepository = &UserRepository{
			client: client,
		}
	}

	return userRepository

}

func (ur *UserRepository) UserExists(email string) (bool, error) {
	if ur.client == nil {
		log.Fatal("Mongo client is nil")
	}
	var user models.User
	userCollection := ur.client.Database("Crud").Collection("user")
	err := userCollection.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		log.Print("Error finding user with email", email, "-", err.Error())
		return false, nil
	}

	return true, nil
}

func (ur *UserRepository) CreateUser(userReq models.User) error {
	userReq.ID = uuid.New().String()

	userCollection := ur.client.Database("Crud").Collection("user")
	_, err := userCollection.InsertOne(context.Background(), userReq)
	if err != nil {
		log.Print("Error creating user", err.Error())
	}

	return nil
}

func (ur *UserRepository) ValidateUser(email string, password string) (string, error) {
	userCollection := ur.client.Database("Crud").Collection("user")
	var user models.User
	err := userCollection.FindOne(context.Background(), bson.M{"email": email, "password": password}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return "", status.Error(codes.NotFound, "User not found")
		}
		log.Print("Error fetching user")
		return "", err
	}

	return user.ID, nil
}

func (ur *UserRepository) GetUserInfo(id string) (*models.GetUserInfoResponse, error) {
	userCollection := ur.client.Database("Crud").Collection("user")
	var user models.User
	err := userCollection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, status.Error(codes.NotFound, "User not found")
		}
		log.Print("Error fetching user info", err.Error())
		return nil, err
	}

	return &models.GetUserInfoResponse{
		Name:  user.Name,
		Email: user.Email,
		Phone: user.Phone,
	}, nil
}
