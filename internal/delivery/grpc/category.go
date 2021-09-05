package grpc

import (
	"context"

	"github.com/duyledat197/netfix-backend/internal/domain"
	"github.com/duyledat197/netfix-backend/internal/models"
	"github.com/duyledat197/netfix-backend/pb"
	"github.com/gosimple/slug"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/emptypb"
)

type categoryDelivery struct {
	categoryDomain domain.CategoryDomain
	pb.UnimplementedCategoryServiceServer
}

// NewAuthDelivery ...
func NewCategoryDelivery(categoryDomain domain.CategoryDomain) pb.CategoryServiceServer {
	return &categoryDelivery{
		categoryDomain: categoryDomain,
	}
}

func (d *categoryDelivery) CreateCategory(ctx context.Context, req *pb.CreateCategoryRequest) (*pb.CreateCategoryResponse, error) {
	if err := d.categoryDomain.Create(ctx, &models.Category{
		Name: req.Name,
		Slug: slug.Make(req.Name),
	}); err != nil {
		return nil, err
	}
	return &pb.CreateCategoryResponse{
		Name:    req.Name,
		Success: true,
	}, nil
}

func (d *categoryDelivery) DeleteCategory(ctx context.Context, req *pb.DeleteCategoryRequest) (*pb.DeleteCategoryResponse, error) {
	c, err := primitive.ObjectIDFromHex(req.CategoryId)
	if err != nil {
		return nil, err
	}

	if err := d.categoryDomain.Delete(ctx, c); err != nil {
		return nil, err
	}
	return &pb.DeleteCategoryResponse{
		Success: true,
	}, nil
}

func (d *categoryDelivery) GetListCategory(ctx context.Context, req *emptypb.Empty) (*pb.GetListCategoryResponse, error) {
	cs, err := d.categoryDomain.GetList(ctx)
	if err != nil {
		return nil, err
	}
	var categories []*pb.Category
	for _, v := range cs {
		categories = append(categories, &pb.Category{
			Id:   v.ID.Hex(),
			Name: v.Name,
			Slug: v.Slug,
		})
	}
	return &pb.GetListCategoryResponse{
		Categories: categories,
		Success:    true,
	}, nil
}
