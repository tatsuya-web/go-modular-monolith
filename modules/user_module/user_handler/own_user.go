package user_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tatuya-web/go-modular-monolith/handler"
)

type OwnUser struct {
	Service OwnUserService
}

func NewOwnUserHandler(uo OwnUserService) *OwnUser {
	return &OwnUser{Service: uo}
}

func (uo *OwnUser) ServeHTTP(ctx *gin.Context) {
	user, err := uo.Service.OwnUser(ctx)
	if err != nil {
		handler.ErrResponse(ctx, http.StatusInternalServerError, "faild to user", err.Error())
		return
	}

	handler.APIResponse(ctx, http.StatusOK, "ユーザー情報", user)
}
