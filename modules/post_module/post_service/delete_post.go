package post_service

import (
	"context"
	"fmt"

	"github.com/tatuya-web/go-modular-monolith/modules/post_module/post_model"
	"github.com/tatuya-web/go-modular-monolith/repository"
)

type DeletePost struct {
	DBExec  repository.Execer
	DBQuery repository.Queryer
	Repo    PostDeleter
}

func (dp *DeletePost) DeletePost(ctx context.Context, id post_model.PostID) error {
	if !dp.Repo.IsOwnPost(ctx, dp.DBQuery, id) {
		return fmt.Errorf("failed to post: %w", repository.ErrUnauthorizedUser)
	}

	p := &post_model.Post{
		ID: id,
	}
	err := dp.Repo.DeletePost(ctx, dp.DBExec, p)
	if err != nil {
		return fmt.Errorf("failed to post: %w", err)
	}

	return nil
}
