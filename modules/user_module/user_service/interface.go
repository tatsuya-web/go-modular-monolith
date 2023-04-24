package user_service

import (
	"context"
	"net/http"

	"github.com/tatuya-web/go-modular-monolith/modules/user_module/user_model"
	"github.com/tatuya-web/go-modular-monolith/repository"
)

//go:generate go run github.com/matryer/moq -out moq_test.go . PostLister UserRegister UserGetter TokenGenerator TokenDeleter OwnGetter UserDeleter
type UserRegister interface {
	RegisterUser(ctx context.Context, db repository.Execer, u *user_model.User) error
}

type TokenGenerator interface {
	GenerateToken(ctx context.Context, u user_model.User) ([]byte, error)
}

type TokenDeleter interface {
	DeleteToken(ctx context.Context, r *http.Request, id user_model.UserID) error
}

type UserGetter interface {
	GetUser(ctx context.Context, db repository.Queryer, email string) (*user_model.User, error)
}

type UserUpdater interface {
	UpdateUser(ctx context.Context, db repository.Execer, u *user_model.User) error
}

type UserDeleter interface {
	DeleteUser(ctx context.Context, db repository.Execer, p *user_model.User) error
}

type OwnGetter interface {
	GetOwn(ctx context.Context, db repository.Queryer, id user_model.UserID) (*user_model.User, error)
}
