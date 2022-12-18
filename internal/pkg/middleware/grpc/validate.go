package grpc

import (
	"context"
	"strings"

	"github.com/status-mok/server/internal/pkg/errors"
	"google.golang.org/grpc"
)

type validator interface {
	ValidateAll() error
}

func NewMessageValidatorInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		v, ok := req.(validator)
		if !ok {
			return handler(ctx, req)
		}

		err := v.ValidateAll()
		if err != nil {
			msg := strings.ReplaceAll(err.Error(), "runes", "symbols")
			return nil, errors.New(msg)
		}

		return handler(ctx, req)
	}
}
