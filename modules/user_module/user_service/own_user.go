package user_service

import (
	"context"
	"fmt"

	"github.com/tatuya-web/go-modular-monolith/modules/user_module/auth"
	"github.com/tatuya-web/go-modular-monolith/modules/user_module/user_model"
	"github.com/tatuya-web/go-modular-monolith/repository"
)

type OwnUser struct {
	DB   repository.Queryer
	Repo OwnGetter
}

func (uo *OwnUser) OwnUser(ctx context.Context) (*user_model.User, error) {
	id, ok := auth.GetUserID(ctx)
	if !ok {
		return nil, fmt.Errorf("ユーザーが見つかりません。")
	}

	user, err := uo.Repo.GetOwn(ctx, uo.DB, id)
	if err != nil {
		return nil, fmt.Errorf("failed to list: %w", err)
	}
	return user, nil
}
