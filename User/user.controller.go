package user

import (
	"context"
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
	err := uc.us.RegisterUser(r.Name, r.Email, r.Password, r.Phone)
	if err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}

func (uc *UserController) LoginUser(ctx context.Context, r *pb.LoginRequest) (*pb.LoginResponse, error) {
	return &pb.LoginResponse{}, nil
}
