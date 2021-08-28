package domain

import (
	"github.com/duyledat197/netfix-backend/internal/models"
	"github.com/duyledat197/netfix-backend/internal/repository"
)

// UserDomain ...
type UserDomain interface {
	CreateUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(userID string) (*models.User, error)
	GetListUser() ([]*models.User, error)
	UpdateUser(userID string, user *models.User) error
	DeleteUser(userID string) error
}

type userDomain struct {
	userRepository repository.UserRepository
}

// NewUserDomain ...
func NewUserDomain(userRepository repository.UserRepository) UserDomain {
	return &userDomain{
		userRepository,
	}
}

func (d *userDomain) CreateUser(user *models.User) error {
	return nil
}

func (d *userDomain) GetUserByEmail(email string) (*models.User, error) {
	return nil, nil
}

func (d *userDomain) GetUserByID(userID string) (*models.User, error) {
	return nil, nil
}

func (d *userDomain) GetListUser() ([]*models.User, error) {
	return []*models.User{}, nil
}

func (d *userDomain) UpdateUser(userID string, user *models.User) error {
	return nil
}

func (d *userDomain) DeleteUser(userID string) error {
	return nil
}
