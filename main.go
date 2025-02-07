package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"user_crud/db"
	"user_crud/middleware"

	pb "user_crud/pb/user_crud/pb"
	user "user_crud/user"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	_, err = db.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize DB: %v", err)
	}

	fmt.Println("DB connected")

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(middleware.PanicRecoveryInterceptor, middleware.AuthMiddleware),
	)

	userController := user.NewUserController(user.NewUserService(user.NewUserRepository()))

	pb.RegisterUserServiceServer(grpcServer, userController)

	port := ":50051"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", port, err)
	}

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		fmt.Println("Starting gRPC server on", port)
		if err := grpcServer.Serve(lis); err != nil {
			fmt.Printf("gRPC server error: %v\n", err)
		}
	}()

	<-signalChan
	fmt.Println("Received shutdown signal. Gracefully shutting down...")

	grpcServer.GracefulStop()
	fmt.Println("gRPC server successfully shut down.")
}
