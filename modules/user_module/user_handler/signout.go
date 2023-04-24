package user_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tatuya-web/go-modular-monolith/handler"
)

type Signout struct {
	Service SignoutService
}

func NewSignoutHandler(s SignoutService) *Signout {
	return &Signout{Service: s}
}

func (s *Signout) ServeHTTP(ctx *gin.Context) {
	if err := s.Service.Signout(ctx, ctx.Request); err != nil {
		handler.ErrResponse(ctx, http.StatusInternalServerError, "サインアウトエラー", err.Error())
		return
	}

	handler.APIResponse(ctx, http.StatusOK, "サインアウトしました。", nil)
}
