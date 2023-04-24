package user_handler

import (
	"context"
	"net/http"

	"github.com/tatuya-web/go-modular-monolith/modules/user_module/user_model"
)

//go:generate go run github.com/matryer/moq -out moq_test.go . RegisterUserService SigninService SignoutService OwnUserService UpdateUserService DeleteUserService
type RegisterUserService interface {
	RegisterUser(ctx context.Context, email, password, role string) (*user_model.User, error)
}

type SigninService interface {
	Signin(ctx context.Context, email, pw string) (string, error)
}

type SignoutService interface {
	Signout(ctx context.Context, r *http.Request) error
}

type OwnUserService interface {
	OwnUser(ctx context.Context) (*user_model.User, error)
}

type UpdateUserService interface {
	UpdateUser(ctx context.Context, id user_model.UserID, email string) (*user_model.User, error)
}

type DeleteUserService interface {
	DeleteUser(ctx context.Context, r *http.Request, id user_model.UserID) error
}
