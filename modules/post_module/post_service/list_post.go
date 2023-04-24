package post_service

import (
	"context"
	"fmt"

	"github.com/tatuya-web/go-modular-monolith/modules/post_module/post_model"
	"github.com/tatuya-web/go-modular-monolith/modules/user_module/auth"
	"github.com/tatuya-web/go-modular-monolith/repository"
)

type ListPost struct {
	DB   repository.Queryer
	Repo PostLister
}

func (lp *ListPost) ListPosts(ctx context.Context) (post_model.Posts, error) {
	id, ok := auth.GetUserID(ctx)
	if !ok {
		return nil, fmt.Errorf("ユーザーが見つかりません。")
	}

	posts, err := lp.Repo.ListPosts(ctx, lp.DB, id)
	if err != nil {
		return nil, fmt.Errorf("failed to list: %w", err)
	}
	return posts, nil
}
