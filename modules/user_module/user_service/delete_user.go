package user_service

import (
	"context"
	"fmt"
	"net/http"

	"github.com/tatuya-web/go-modular-monolith/modules/user_module/auth"
	"github.com/tatuya-web/go-modular-monolith/modules/user_module/user_model"
	"github.com/tatuya-web/go-modular-monolith/repository"
)

type DeleteUser struct {
	DB           repository.Execer
	Repo         UserDeleter
	TokenDeleter TokenDeleter
}

func (du *DeleteUser) DeleteUser(ctx context.Context, r *http.Request, id user_model.UserID) error {
	if !auth.CheckOwn(ctx, id) {
		return fmt.Errorf("権限のないユーザーです。")
	}

	if err := du.TokenDeleter.DeleteToken(ctx, r, id); err != nil {
		return fmt.Errorf("failed to token: %w", err)
	}

	u := &user_model.User{
		ID: id,
	}
	err := du.Repo.DeleteUser(ctx, du.DB, u)
	if err != nil {
		return fmt.Errorf("failed to post: %w", err)
	}

	return nil
}
