package repository

import "github.com/duyledat197/netfix-backend/internal/models"

// UserRepository ...
type UserRepository interface {
	Create(user *models.User) error
	UpdateInfo(userID string, user *models.User) error
	UpdatePassword(userID string, password string) error
	FindByID(userID string) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	FindAll(offset int64, limit int64) ([]*models.User, error)
}
