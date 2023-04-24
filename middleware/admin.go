package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tatuya-web/go-modular-monolith/handler"
	"github.com/tatuya-web/go-modular-monolith/modules/user_module/auth"
)

func AdminMiddleware() gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		if auth.IsAdmin(ctx.Request.Context()) {
			handler.ErrResponse(ctx, http.StatusUnauthorized, "認証エラー", "アクセス権限がありません。")
			return
		}
		ctx.Next()
	})
}
