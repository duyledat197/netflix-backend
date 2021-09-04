package grpc

import (
	"context"
	"time"

	"github.com/duyledat197/netfix-backend/internal/domain"
	"github.com/duyledat197/netfix-backend/internal/models"
	"github.com/duyledat197/netfix-backend/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type postDelivery struct {
	postDomain domain.PostDomain
	pb.UnimplementedPostServiceServer
}

// NewAuthDelivery ...
func NewPostDelivery(postDomain domain.PostDomain) pb.PostServiceServer {
	return &postDelivery{
		postDomain: postDomain,
	}
}

func (d *postDelivery) CreatePost(ctx context.Context, req *pb.CreatePostRequest) (*pb.CreatePostResponse, error) {
	categoryId, err := primitive.ObjectIDFromHex(req.CategoryId)
	if err != nil {
		return nil, err
	}
	post := &models.Post{
		Content:    req.Content,
		CategoryID: categoryId,
		Thumbnail:  req.Thumbnail,
		Describe:   req.Describe,
		Title:      req.Title,
	}

	if err := d.postDomain.Create(ctx, post); err != nil {
		return nil, err
	}
	return &pb.CreatePostResponse{
		Success: true,
	}, nil
}
func (d *postDelivery) GetPosts(ctx context.Context, req *pb.GetPostsRequest) (*pb.GetPostsResponse, error) {
	list, total, err := d.postDomain.GetList(ctx, req.Offset, req.Limit)
	if err != nil {
		return nil, err
	}
	var posts []*pb.GetPostDetailResponse
	for _, v := range list {
		posts = append(posts, &pb.GetPostDetailResponse{
			Content:    v.Content,
			Describe:   v.Describe,
			Thumbnail:  v.Thumbnail,
			CategoryId: v.CategoryID.Hex(),
			HeartCount: v.HeartCount,
			View:       v.View,
			CreatedAt:  v.CreatedAt.Format(time.RFC3339),
			Title:      v.Title,
			Id:         v.ID.Hex(),
			Slug:       v.Slug,
		})
	}
	return &pb.GetPostsResponse{
		Posts: posts,
		Total: total,
	}, nil
}
func (d *postDelivery) GetPostDetail(ctx context.Context, req *pb.GetPostDetailRequest) (*pb.GetPostDetailResponse, error) {

	p, err := d.postDomain.GetDetail(ctx, req.Slug)
	if err != nil {
		return nil, err
	}

	return &pb.GetPostDetailResponse{
		Content:    p.Content,
		Describe:   p.Describe,
		Thumbnail:  p.Thumbnail,
		CategoryId: p.CategoryID.Hex(),
		HeartCount: p.HeartCount,
		View:       p.View,
		CreatedAt:  p.CreatedAt.String(),
		Slug:       p.Slug,
	}, nil
}
func (d *postDelivery) UpdatePost(ctx context.Context, req *pb.UpdatePostRequest) (*pb.UpdatePostResponse, error) {
	postID, err := primitive.ObjectIDFromHex(req.PostId)
	if err != nil {
		return nil, err
	}
	categoryId, err := primitive.ObjectIDFromHex(req.CategoryId)
	if err != nil {
		return nil, err
	}
	if err := d.postDomain.Update(ctx, &models.Post{
		ID:         postID,
		CategoryID: categoryId,
		Content:    req.Content,
		Thumbnail:  req.Thumbnail,
		Describe:   req.Describe,
		Title:      req.Title,
	}); err != nil {
		return nil, err
	}
	return &pb.UpdatePostResponse{
		Success: true,
	}, nil
}
func (d *postDelivery) DeletePost(ctx context.Context, req *pb.DeletePostRequest) (*pb.DeletePostResponse, error) {
	postID, err := primitive.ObjectIDFromHex(req.PostId)
	if err != nil {
		return nil, err
	}
	if err := d.postDomain.Delete(ctx, postID); err != nil {
		return nil, err
	}
	return &pb.DeletePostResponse{
		Success: true,
	}, nil
}
