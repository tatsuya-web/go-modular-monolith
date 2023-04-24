package post_service

import (
	"context"
	"fmt"

	"github.com/tatuya-web/go-modular-monolith/modules/post_module/post_model"
	"github.com/tatuya-web/go-modular-monolith/modules/user_module/auth"
	"github.com/tatuya-web/go-modular-monolith/repository"
)

type AddPost struct {
	DB   repository.Execer
	Repo PostAdder
}

func (ap *AddPost) AddPost(ctx context.Context, title string, content string) (*post_model.Post, error) {
	id, ok := auth.GetUserID(ctx)
	if !ok {
		return nil, fmt.Errorf("ユーザーが見つかりません。")
	}

	p := &post_model.Post{
		Title:   title,
		Content: content,
		UserID:  id,
	}
	err := ap.Repo.AddPost(ctx, ap.DB, p)
	if err != nil {
		return nil, fmt.Errorf("failed to post: %w", err)
	}

	return p, nil
}
