package repository

import (
	"github.com/duyledat197/netfix-backend/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CategoryRepository interface {
	Create(category *models.Category) error
	FindAll() ([]*models.Category, error)
	FindByID(id primitive.ObjectID) (*models.Category, error)
	Delete(id primitive.ObjectID) error
}
