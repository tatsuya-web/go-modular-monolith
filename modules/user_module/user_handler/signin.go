package user_handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/tatuya-web/go-modular-monolith/handler"
	"github.com/tatuya-web/go-modular-monolith/repository"
)

type Signin struct {
	Service   SigninService
	Validator *validator.Validate
}

func NewSigninHandler(s SigninService, v *validator.Validate) *Signin {
	return &Signin{Service: s, Validator: v}
}

func (s *Signin) ServeHTTP(ctx *gin.Context) {
	var input struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}
	const errTitle = "サインインエラー"

	if err := ctx.ShouldBindJSON(&input); err != nil {
		handler.ErrResponse(ctx, http.StatusBadRequest, errTitle, err.Error())
		return
	}

	if err := s.Validator.Struct(input); err != nil {
		handler.ErrResponse(ctx, http.StatusBadRequest, errTitle, err.Error())
		return
	}
	jwt, err := s.Service.Signin(ctx, input.Email, input.Password)
	if err != nil {
		if errors.Is(err, repository.ErrNotFoundSession) {
			handler.ErrResponse(ctx, http.StatusUnauthorized, errTitle, repository.ErrNotFoundSession.Error())
			return
		}
		if errors.Is(err, repository.ErrAlreadyEntry) {
			handler.ErrResponse(ctx, http.StatusConflict, errTitle, repository.ErrAlreadyEntry.Error())
			return
		}
		handler.ErrResponse(ctx, http.StatusInternalServerError, errTitle, err.Error())
		return
	}

	rsp := struct {
		AccessToken string `json:"access_token"`
	}{
		AccessToken: jwt,
	}

	handler.APIResponse(ctx, http.StatusCreated, "サインイン成功しました。", rsp)
}
