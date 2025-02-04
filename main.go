package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"user_crud/controllers"
	"user_crud/db"
	"user_crud/repositories"
	"user_crud/services"

	pb "user_crud/pb/user_crud/pb"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		panic("Invalid mongo uri")
	}
	client, err := db.Connect(mongoURI)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	defer db.Disconnect()

	userRepo := repositories.NewMongoUserRepository(client)
	userService := services.NewUserService(userRepo)

	userController := controllers.NewUserController(userService)

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, userController)

	port := ":50051"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("Failed to listen on port %s: %v", port, err)
	}

	fmt.Println("Server is running on port", port)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
