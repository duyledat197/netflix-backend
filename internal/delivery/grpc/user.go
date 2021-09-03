package grpc

import (
	"context"

	"github.com/duyledat197/netfix-backend/common/enum"
	"github.com/duyledat197/netfix-backend/internal/domain"
	"github.com/duyledat197/netfix-backend/internal/models"
	"github.com/duyledat197/netfix-backend/pb"
)

type userDelivery struct {
	userDomain domain.UserDomain
	pb.UnimplementedUserServiceServer
}

// NewUserDelivery ...
func NewUserDelivery(userDomain domain.UserDomain) pb.UserServiceServer {
	return &userDelivery{
		userDomain: userDomain,
	}
}

func (r *userDelivery) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	user := &models.User{
		Email:    req.Email,
		Password: req.Password,
		Role:     enum.UserRole(req.Role),
	}
	if err := r.userDomain.CreateUser(user); err != nil {
		return nil, err
	}
	return &pb.CreateUserResponse{
		Success: true,
	}, nil
}
