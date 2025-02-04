package user

import (
	"time"
	"user_crud/models"

	"github.com/golang-jwt/jwt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserService struct {
	userRepo *UserRepository
}

func NewUserService(userRepo *UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (us *UserService) RegisterUser(name, email, password, phone string) error {
	existingUser, err := us.userRepo.UserExists(email)
	if err != nil {
		return err
	}

	if existingUser {
		return status.Error(codes.Aborted, "User already exists with this email")
	}

	user := models.User{
		Email:    email,
		Password: password,
		Name:     name,
		Phone:    phone,
	}

	err = us.userRepo.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (us *UserService) LoginUser(email, password string) (string, error) {
	userId, err := us.userRepo.ValidateUser(email, password)
	if err != nil {
		return "", err
	}

	token, err := generateJWT(userId)
	if err != nil {
		return "", err
	}

	return token, nil
}

func generateJWT(userId string) (string, error) {
	claims := jwt.MapClaims{
		"sub": userId,
		"iat": time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("PASSWORD"))
}
