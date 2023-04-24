package post_service

import (
	"context"
	"fmt"

	"github.com/tatuya-web/go-modular-monolith/modules/post_module/post_model"
	"github.com/tatuya-web/go-modular-monolith/repository"
)

type UpdatePost struct {
	DBExec  repository.Execer
	DBQuery repository.Queryer
	Repo    PostUpdater
}

func (up *UpdatePost) UpdatePost(ctx context.Context, id post_model.PostID, title string, content string) (*post_model.Post, error) {
	if !up.Repo.IsOwnPost(ctx, up.DBQuery, id) {
		return nil, fmt.Errorf("failed to post: %w", repository.ErrUnauthorizedUser)
	}

	p := &post_model.Post{
		ID:      id,
		Title:   title,
		Content: content,
	}
	err := up.Repo.UpdatePost(ctx, up.DBExec, p)
	if err != nil {
		return nil, fmt.Errorf("failed to post: %w", err)
	}

	return p, nil
}
