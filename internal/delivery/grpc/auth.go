package grpc

import (
	"context"

	"github.com/duyledat197/netfix-backend/internal/domain"
	"github.com/duyledat197/netfix-backend/pb"
)

type authDelivery struct {
	authDomain domain.AuthDomain
	pb.UnimplementedAuthServiceServer
}

// NewAuthDelivery ...
func NewAuthDelivery(authDomain domain.AuthDomain) pb.AuthServiceServer {
	return &authDelivery{
		authDomain: authDomain,
	}
}

func (d *authDelivery) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}
	u, tkn, err := d.authDomain.Login(ctx, req.GetEmail(), req.GetPassword())
	if err != nil {
		return nil, err
	}
	res := &pb.LoginResponse{
		UserID:       u.ID.Hex(),
		CommentCount: u.CommentCount,
		PostCount:    u.PostCount,
		NiceName:     u.NiceName,
		Role:         pb.UserRole(u.Role),
		Token:        tkn,
	}
	return res, nil
}
