package user

import (
	"context"
	"user_crud/models"
	pb "user_crud/pb/user_crud/pb"
)

type UserController struct {
	us *UserService
	pb.UnimplementedUserServiceServer
}

func NewUserController(us *UserService) *UserController {
	uc := &UserController{
		us: us,
	}
	return uc
}

func (uc *UserController) RegisterUser(ctx context.Context, r *pb.RegisterRequest) (*pb.Empty, error) {
	err := uc.us.RegisterUser(models.UserRegisterRequest{
		Name:     r.Name,
		Email:    r.Email,
		Password: r.Password,
		Phone:    r.Phone,
	})
	if err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}

func (uc *UserController) LoginUser(ctx context.Context, r *pb.LoginRequest) (*pb.LoginResponse, error) {
	token, err := uc.us.LoginUser(models.UserLoginRequest{
		Email:    r.Email,
		Password: r.Password,
	})
	if err != nil {
		return nil, err
	}

	return &pb.LoginResponse{
		Token: token,
	}, nil
}

func (uc *UserController) GetUserInfo(ctx context.Context, r *pb.Empty) (*pb.GetUserInfoResponse, error) {
	userId := ctx.Value("userId").(string)

	userInfo, err := uc.us.GetUserInfo(models.GetUserInfoRequest{ID: userId})
	if err != nil {
		return nil, err
	}

	return &pb.GetUserInfoResponse{
		Email: userInfo.Email,
		Name:  userInfo.Email,
		Phone: userInfo.Phone,
	}, nil
}
