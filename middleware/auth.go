package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tatuya-web/go-modular-monolith/handler"
	"github.com/tatuya-web/go-modular-monolith/modules/user_module/auth"
)

func AuthMiddleware(j *auth.JWTer) gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		if err := j.FillContext(ctx); err != nil {
			handler.ErrResponse(ctx, http.StatusUnauthorized, "認証エラー", "アクセストークンが無効です。再ログインしてください。")
			return
		}
		ctx.Next()
	})
}
