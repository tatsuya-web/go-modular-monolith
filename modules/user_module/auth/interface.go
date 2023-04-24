package auth

import (
	"context"

	"github.com/tatuya-web/go-modular-monolith/modules/user_module/user_model"
)

//go:generate go run github.com/matryer/moq -out moq_test.go . Store
type Store interface {
	Save(ctx context.Context, key string, userID user_model.UserID) error
	Load(ctx context.Context, key string) (user_model.UserID, error)
	Delete(ctx context.Context, key string) error
}
