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

type MongoUserRepository struct {
	client *mongo.Client
}

func NewMongoUserRepository(client *mongo.Client) *MongoUserRepository {
	return &MongoUserRepository{client: client}
}

func (ur *MongoUserRepository) getUserCollection() *mongo.Collection {
	return ur.client.Database("Crud").Collection("user")
}

func (ur *MongoUserRepository) UserExists(email string) (bool, error) {
	userCollection := ur.getUserCollection()
	var user models.User
	err := userCollection.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		log.Print("Error finding user with email", email, "-", err.Error())
		return false, err
	}
	return true, nil
}

func (ur *MongoUserRepository) CreateUser(userReq models.User) error {
	userReq.ID = uuid.New().String()

	userCollection := ur.getUserCollection()
	_, err := userCollection.InsertOne(context.Background(), userReq)
	if err != nil {
		log.Print("Error creating user", err.Error())
		return err
	}

	return nil
}

func (ur *MongoUserRepository) ValidateUser(email string, password string) (string, error) {
	userCollection := ur.getUserCollection()
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

func (ur *MongoUserRepository) GetUserInfo(id string) (*models.GetUserInfoResponse, error) {
	userCollection := ur.getUserCollection()
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
