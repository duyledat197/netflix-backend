package domain

import (
	"context"
	"fmt"
	"time"

	"github.com/duyledat197/netfix-backend/internal/models"
	"github.com/duyledat197/netfix-backend/internal/repository"
	"github.com/gosimple/slug"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// PostDomain ...
type PostDomain interface {
	Create(ctx context.Context, post *models.Post) error
	Update(ctx context.Context, post *models.Post) error
	GetDetail(ctx context.Context, slug string) (*models.Post, error)
	GetList(ctx context.Context, offset, limit int64) ([]*models.Post, int64, error)
	Delete(ctx context.Context, id primitive.ObjectID) error
}

type postDomain struct {
	postRepository     repository.PostRepository
	categoryRepository repository.CategoryRepository
}

// NewPostDomain ...
func NewPostDomain(postRepository repository.PostRepository, categoryRepository repository.CategoryRepository) PostDomain {
	return &postDomain{
		postRepository:     postRepository,
		categoryRepository: categoryRepository,
	}
}

func (d *postDomain) Create(ctx context.Context, post *models.Post) error {
	c, err := d.categoryRepository.FindByID(post.CategoryID)
	if err != nil {
		return fmt.Errorf("category not exists")
	}
	post.CreatedAt = time.Now()
	post.HeartCount = 0
	post.View = 0
	post.Slug = slug.Make(post.Title)
	post.CategorySlug = c.Slug
	post.CategoryName = c.Name
	return d.postRepository.Create(post)
}
func (d *postDomain) Update(ctx context.Context, post *models.Post) error {
	post.UpdatedAt = time.Now()
	return d.postRepository.Update(post.ID, post)
}
func (d *postDomain) Delete(ctx context.Context, id primitive.ObjectID) error {
	return d.postRepository.Delete(id)
}

func (d *postDomain) GetDetail(ctx context.Context, slug string) (*models.Post, error) {
	return d.postRepository.FindBySlug(slug)
}
func (d *postDomain) GetList(ctx context.Context, offset, limit int64) ([]*models.Post, int64, error) {
	posts, err := d.postRepository.FindAll(offset, limit, primitive.NilObjectID)
	if err != nil {
		return nil, 0, err
	}
	total, err := d.postRepository.Count()
	if err != nil {
		return nil, 0, err
	}
	return posts, total, nil
}
