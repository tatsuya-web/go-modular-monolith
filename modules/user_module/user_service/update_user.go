package user_service

import (
	"context"
	"fmt"

	"github.com/tatuya-web/go-modular-monolith/modules/user_module/auth"
	"github.com/tatuya-web/go-modular-monolith/modules/user_module/user_model"
	"github.com/tatuya-web/go-modular-monolith/repository"
)

type UpdateUser struct {
	DB   repository.Execer
	Repo UserUpdater
}

func (uu *UpdateUser) UpdateUser(ctx context.Context, id user_model.UserID, email string) (*user_model.User, error) {
	if !auth.CheckOwn(ctx, id) {
		return nil, fmt.Errorf("権限のないユーザーです。")
	}

	u := &user_model.User{
		ID:    id,
		Email: email,
	}
	err := uu.Repo.UpdateUser(ctx, uu.DB, u)
	if err != nil {
		return nil, fmt.Errorf("failed to post: %w", err)
	}

	return u, nil
}
