package user_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/tatuya-web/go-modular-monolith/handler"
	"github.com/tatuya-web/go-modular-monolith/modules/user_module/user_model"
)

type RegisterUser struct {
	Service   RegisterUserService
	Validator *validator.Validate
}

func NewRegisterUserHandler(ru RegisterUserService, v *validator.Validate) *RegisterUser {
	return &RegisterUser{Service: ru, Validator: v}
}

func (ru *RegisterUser) ServeHTTP(ctx *gin.Context) {
	var input struct {
		Email    string `json:"email" validate:"required,email,max=255"`
		Password string `json:"password" validate:"required"`
		Role     string `json:"role" validate:"required"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		handler.ErrResponse(ctx, http.StatusBadRequest, err.Error(), err.Error())
		return
	}

	if err := ru.Validator.Struct(input); err != nil {
		handler.ErrResponse(ctx, http.StatusBadRequest, err.Error(), err.Error())
		return
	}

	u, err := ru.Service.RegisterUser(ctx, input.Email, input.Password, input.Role)

	if err != nil {
		handler.ErrResponse(ctx, http.StatusBadRequest, err.Error(), err.Error())
		return
	}
	rsp := struct {
		ID user_model.UserID `json:"id"`
	}{ID: u.ID}

	handler.APIResponse(ctx, http.StatusCreated, "本登録が完了しました。", rsp)
}
