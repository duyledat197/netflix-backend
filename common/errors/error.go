package errors

import "google.golang.org/grpc"

var (
	// ErrPasswordIsNotCorrect ...
	ErrPasswordIsNotCorrect = grpc.Errorf(10001, "password is not correct")

	// ErrEmailNotFound ...
	ErrEmailNotFound = grpc.Errorf(10002, "email not found")

	// ErrEmailAlreadyExists ...
	ErrEmailAlreadyExists = grpc.Errorf(10003, "email already exists")
)
