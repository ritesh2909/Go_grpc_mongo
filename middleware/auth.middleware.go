package middleware

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/golang-jwt/jwt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var JWT_SECRET_KEY []byte = []byte("PASSWORD")

func AuthMiddleware(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

	publicMethods := []string{
		"/pb.UserService/RegisterUser",
		"/pb.UserService/LoginUser",
	}

	for _, method := range publicMethods {
		if info.FullMethod == method {
			return handler(ctx, req)
		}
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("missing metadata")
	}

	authHeader := md["authorization"]
	if len(authHeader) == 0 {
		return nil, errors.New("missing Authorization header")
	}

	tokenString := strings.TrimPrefix(authHeader[0], "Bearer ")
	if tokenString == "" {
		return nil, errors.New("missing token")
	}

	claims, err := parseJWT(tokenString)
	if err != nil {
		return nil, fmt.Errorf("invalid token: %v", err)
	}

	userId, ok := claims["sub"].(string)
	if !ok {
		return nil, errors.New("invalid token: user ID missing")
	}

	newCtx := context.WithValue(ctx, "userId", userId)

	return handler(newCtx, req)
}

func parseJWT(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(*jwt.Token) (interface{}, error) {
		return JWT_SECRET_KEY, nil
	})
	if err != nil || !token.Valid {
		log.Print("Error parsing token", err.Error())
		return nil, status.Errorf(codes.Unauthenticated, "invalid auth token: %v", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token claims")
}
