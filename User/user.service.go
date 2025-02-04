package user

import (
	"time"
	"user_crud/models"

	"github.com/golang-jwt/jwt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserServiceInterface interface {
	RegisterUser(userReq models.UserRegisterRequest) error
	LoginUser(userReq models.UserLoginRequest) (string, error)
}

type UserService struct {
	userRepo *UserRepository
}

func NewUserService(userRepo *UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (us *UserService) RegisterUser(r models.UserRegisterRequest) error {
	existingUser, err := us.userRepo.UserExists(r.Email)
	if err != nil {
		return err
	}

	if existingUser {
		return status.Error(codes.Aborted, "User already exists with this email")
	}

	user := models.User{
		Email:    r.Email,
		Password: r.Password,
		Name:     r.Name,
		Phone:    r.Phone,
	}

	err = us.userRepo.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (us *UserService) LoginUser(r models.UserLoginRequest) (string, error) {
	userId, err := us.userRepo.ValidateUser(r.Email, r.Password)
	if err != nil {
		return "", err
	}

	token, err := generateJWT(userId)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (us *UserService) GetUserInfo(r models.GetUserInfoRequest) (*models.GetUserInfoResponse, error) {
	userInfo, err := us.userRepo.GetUserInfo(r.ID)
	if err != nil {
		return nil, err
	}
	return userInfo, nil
}

func generateJWT(userId string) (string, error) {
	claims := jwt.MapClaims{
		"sub": userId,
		"iat": time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("PASSWORD"))
}
