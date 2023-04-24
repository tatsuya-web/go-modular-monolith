package user_service

import (
	"context"
	"fmt"

	"github.com/tatuya-web/go-modular-monolith/modules/user_module/user_model"
	"github.com/tatuya-web/go-modular-monolith/repository"
	"golang.org/x/crypto/bcrypt"
)

type RegisterUser struct {
	DB   repository.Execer
	Repo UserRegister
}

func (r *RegisterUser) RegisterUser(
	ctx context.Context,
	email,
	password,
	role string) (*user_model.User, error) {
	pw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("cannot hash password: %w", err)
	}
	u := &user_model.User{
		Email:    email,
		Password: string(pw),
		Role:     role,
	}

	if err := r.Repo.RegisterUser(ctx, r.DB, u); err != nil {
		return nil, fmt.Errorf("faild to register: %w", err)
	}
	return u, nil
}
