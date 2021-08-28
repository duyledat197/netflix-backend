package grpc

import (
	"context"

	"github.com/duyledat197/netfix-backend/internal/domain"
	"google.golang.org/grpc"
)

func NewGRPCServer(ctx context.Context, authDomain domain.AuthDomain) *grpc.Server {
	return nil
}
