package repository

import (
	"github.com/duyledat197/netfix-backend/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PostRepository interface {
	Create(post *models.Post) error
	Update(id primitive.ObjectID, post *models.Post) error
	FindByID(id primitive.ObjectID) (*models.Post, error)
	FindAll(offset, limit int64, categoryID primitive.ObjectID) ([]*models.Post, error)
	Delete(id primitive.ObjectID) error
}
