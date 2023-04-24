package user_handler

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/tatuya-web/go-modular-monolith/handler"
	"github.com/tatuya-web/go-modular-monolith/modules/user_module/user_model"
)

type DeleteUser struct {
	Service   DeleteUserService
	Validator *validator.Validate
}

func NewDeleteUserHandler(du DeleteUserService, v *validator.Validate) *DeleteUser {
	return &DeleteUser{Service: du, Validator: v}
}

func (du *DeleteUser) ServeHTTP(ctx *gin.Context) {
	var u struct {
		UserID int64 `json:"user_id" validate:"required"`
	}

	if err := json.NewDecoder(ctx.Request.Body).Decode(&u); err != nil {
		handler.ErrResponse(ctx, http.StatusInternalServerError, "faild to user", err.Error())
		return
	}

	err := du.Validator.Struct(u)
	if err != nil {
		handler.ErrResponse(ctx, http.StatusBadRequest, "faild to user", err.Error())
		return
	}

	if err := du.Service.DeleteUser(ctx, ctx.Request, user_model.UserID(u.UserID)); err != nil {
		handler.ErrResponse(ctx, http.StatusInternalServerError, "faild to user", err.Error())
		return
	}

	handler.APIResponse(ctx, http.StatusOK, "ユーザーを削除しました。", nil)
}
