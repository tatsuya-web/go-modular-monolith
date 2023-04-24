package router

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	"github.com/tatuya-web/go-modular-monolith/config"
	"github.com/tatuya-web/go-modular-monolith/handler"
	"github.com/tatuya-web/go-modular-monolith/modules/user_module/auth"
	"github.com/tatuya-web/go-modular-monolith/modules/user_module/user_handler"
	"github.com/tatuya-web/go-modular-monolith/modules/user_module/user_repository"
	"github.com/tatuya-web/go-modular-monolith/modules/user_module/user_service"
	"github.com/tatuya-web/go-modular-monolith/util"
)

func SetRouting(ctx context.Context, db *sqlx.DB, router *gin.Engine, cfg *config.Config) error {

	clocker := util.RealClocker{}
	urep := user_repository.UserRepo{Clocker: clocker}
	validate := validator.New()
	rcli, err := user_repository.NewKVS(ctx, cfg)
	if err != nil {
		return err
	}
	jwter, err := auth.NewJWTer(rcli, clocker)
	if err != nil {
		return err
	}

	//ルートグループ作成
	rg := router.Group("/api/v1")

	//ヘルスチェック
	healthHandler := handler.NewHealthhandler()
	router.GET("/health", healthHandler.ServeHTTP)

	//ユーザー登録
	registerUserHandler := user_handler.NewRegisterUserHandler(
		&user_service.RegisterUser{DB: db, Repo: &urep},
		validate,
	)
	rg.POST("register", registerUserHandler.ServeHTTP)

	//サインイン
	signinUserHandler := user_handler.NewSigninHandler(
		&user_service.Signin{DB: db, Repo: &urep, TokenGenerator: jwter},
		validate,
	)
	rg.POST("signin", signinUserHandler.ServeHTTP)

	return nil
}
