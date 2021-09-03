package errors

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	// ErrPasswordIsNotCorrect ...
	ErrPasswordIsNotCorrect = status.Errorf(codes.InvalidArgument, "password is not correct")

	// ErrEmailNotFound ...
	ErrEmailNotFound = status.Errorf(codes.InvalidArgument, "email not found")

	// ErrEmailAlreadyExists ...
	ErrEmailAlreadyExists = status.Errorf(codes.AlreadyExists, "email already exists")
)
