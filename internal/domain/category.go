package domain

import (
	"context"
	"time"

	"github.com/duyledat197/netfix-backend/internal/models"
	"github.com/duyledat197/netfix-backend/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CategoryDomain ...
type CategoryDomain interface {
	Create(ctx context.Context, category *models.Category) error
	Delete(ctx context.Context, id primitive.ObjectID) error
	GetList(ctx context.Context) ([]*models.Category, error)
}

type categoryDomain struct {
	categoryRepository repository.CategoryRepository
}

// NewCategoryDomain ...
func NewCategoryDomain(categoryRepository repository.CategoryRepository) CategoryDomain {
	return &categoryDomain{
		categoryRepository: categoryRepository,
	}
}

func (d *categoryDomain) Create(ctx context.Context, category *models.Category) error {
	category.CreatedAt = time.Now()
	return d.categoryRepository.Create(category)
}

func (d *categoryDomain) Delete(ctx context.Context, id primitive.ObjectID) error {
	return d.categoryRepository.Delete(id)
}

func (d *categoryDomain) GetList(ctx context.Context) ([]*models.Category, error) {
	return d.categoryRepository.FindAll()
}
