package grpc

import (
	"context"

	"github.com/status-mok/server/internal/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewErrorUnaryInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		resp, err := handler(ctx, req)
		if err == nil {
			return resp, nil
		}

		causeErr := errors.Cause(errors.Unwrap(err))
		if code := status.Code(causeErr); code != codes.OK && code != codes.Unknown {
			return nil, status.Error(code, err.Error())
		}

		return nil, err
	}
}
