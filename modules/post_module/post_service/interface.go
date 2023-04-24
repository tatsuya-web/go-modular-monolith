package post_service

import (
	"context"

	"github.com/tatuya-web/go-modular-monolith/modules/post_module/post_model"
	"github.com/tatuya-web/go-modular-monolith/modules/user_module/user_model"
	"github.com/tatuya-web/go-modular-monolith/repository"
)

//go:generate go run github.com/matryer/moq -out moq_test.go . PostAdder PostUpdater PostLister
type PostAdder interface {
	AddPost(ctx context.Context, db repository.Execer, p *post_model.Post) error
}

type PostUpdater interface {
	IsOwnPost(ctx context.Context, db repository.Queryer, id post_model.PostID) bool
	UpdatePost(ctx context.Context, db repository.Execer, p *post_model.Post) error
}

type PostDeleter interface {
	IsOwnPost(ctx context.Context, db repository.Queryer, id post_model.PostID) bool
	DeletePost(ctx context.Context, db repository.Execer, p *post_model.Post) error
}

type PostLister interface {
	ListPosts(ctx context.Context, db repository.Queryer, id user_model.UserID) (post_model.Posts, error)
}
