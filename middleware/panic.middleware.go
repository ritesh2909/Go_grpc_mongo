package middleware

import (
	"context"
	"fmt"
	"runtime/debug"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func PanicRecoveryInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	panicErrorPath := "unknown path"
	defer func() {
		if r := recover(); r != nil {
			stackTrace := string(debug.Stack())
			stackTraceParts := strings.Split(stackTrace, "\n")

			for i, v := range stackTraceParts {
				if strings.Contains(v, "panic({") {
					if i+3 < len(stackTraceParts) {
						panicErrorPath = stackTraceParts[i+3]
						break
					}
				}
			}

			fmt.Sprintf("Recovered from panic in method %s: %v path: %s", info.FullMethod, r, panicErrorPath)
			err = status.Errorf(codes.Internal, "%v", r)
		}
	}()
	return handler(ctx, req)
}
