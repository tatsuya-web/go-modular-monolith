package post_handler

import (
	"context"

	"github.com/tatuya-web/go-modular-monolith/modules/post_module/post_model"
)

//go:generate go run github.com/matryer/moq -out moq_test.go . AddPostService UpdatePostService DeletePostService ListPostService
type AddPostService interface {
	AddPost(ctx context.Context, title string, content string) (*post_model.Post, error)
}

type UpdatePostService interface {
	UpdatePost(ctx context.Context, id post_model.PostID, title string, content string) (*post_model.Post, error)
}

type DeletePostService interface {
	DeletePost(ctx context.Context, id post_model.PostID) error
}

type ListPostService interface {
	ListPosts(ctx context.Context) (post_model.Posts, error)
}
