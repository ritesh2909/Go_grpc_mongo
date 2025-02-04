package controllers

import (
	"context"
	"log"
	pb "user_crud/pb/user_crud/pb"
	"user_crud/services"
)

type UserController struct {
	us *services.UserService
	pb.UnimplementedUserServiceServer
}

func NewUserController(us *services.UserService) *UserController {
	uc := &UserController{
		us: us,
	}
	return uc
}

func (uc *UserController) RegisterUser(ctx context.Context, r *pb.RegisterRequest) (*pb.Empty, error) {
	log.Print("Call reached the controller layer")
	err := uc.us.RegisterUser(r.Name, r.Email, r.Password, r.Phone)
	if err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}

func (uc *UserController) LoginUser(ctx context.Context, r *pb.LoginRequest) (*pb.LoginResponse, error) {
	return &pb.LoginResponse{}, nil
}
