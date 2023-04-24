package router

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	"github.com/tatuya-web/go-modular-monolith/config"
	"github.com/tatuya-web/go-modular-monolith/middleware"
	"github.com/tatuya-web/go-modular-monolith/modules/post_module/post_handler"
	"github.com/tatuya-web/go-modular-monolith/modules/post_module/post_repository"
	"github.com/tatuya-web/go-modular-monolith/modules/post_module/post_service"
	"github.com/tatuya-web/go-modular-monolith/modules/user_module/auth"
	"github.com/tatuya-web/go-modular-monolith/modules/user_module/user_handler"
	"github.com/tatuya-web/go-modular-monolith/modules/user_module/user_repository"
	"github.com/tatuya-web/go-modular-monolith/modules/user_module/user_service"
	"github.com/tatuya-web/go-modular-monolith/util"
)

func SetAuthRouting(ctx context.Context, db *sqlx.DB, router *gin.Engine, cfg *config.Config) error {
	clocker := util.RealClocker{}
	prep := post_repository.PostRepo{Clocker: clocker}
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
	rg := router.Group("/api/v1").
		Use(middleware.AuthMiddleware(jwter))

	//POST登録
	addPostHandler := post_handler.NewAddPosthandler(
		&post_service.AddPost{DB: db, Repo: &prep},
		validate,
	)
	rg.POST("posts", addPostHandler.ServeHTTP)

	//POST更新
	updatePostHandler := post_handler.NewUpdatePosthandler(
		&post_service.UpdatePost{DBExec: db, DBQuery: db, Repo: &prep},
		validate,
	)
	rg.PATCH("posts", updatePostHandler.ServeHTTP)

	//POST削除
	deletePostHandler := post_handler.NewDeletePostHandler(
		&post_service.DeletePost{DBExec: db, DBQuery: db, Repo: &prep},
		validate,
	)
	rg.DELETE("posts", deletePostHandler.ServeHTTP)

	//POST一覧 (roleがユーザーの場合は自身にひもずくPOSTだけがし取得される)
	listPostHandler := post_handler.NewListPostHandler(
		&post_service.ListPost{DB: db, Repo: &prep},
	)
	rg.GET("posts", listPostHandler.ServeHTTP)

	//プロフィール
	ownUserHandler := user_handler.NewOwnUserHandler(
		&user_service.OwnUser{DB: db, Repo: &urep},
	)
	rg.GET("user", ownUserHandler.ServeHTTP)

	//USER更新
	updateUserHandler := user_handler.NewUpdateUserHandler(
		&user_service.UpdateUser{DB: db, Repo: &urep},
		validate,
	)
	rg.PATCH("user", updateUserHandler.ServeHTTP)

	//USER削除
	deleteUserHandler := user_handler.NewDeleteUserHandler(
		&user_service.DeleteUser{DB: db, Repo: &urep, TokenDeleter: jwter},
		validate,
	)
	rg.DELETE("user", deleteUserHandler.ServeHTTP)

	//サインアウト
	signoutHandler := user_handler.NewSignoutHandler(
		&user_service.Signout{TokenDeleter: jwter},
	)
	rg.POST("signout", signoutHandler.ServeHTTP)

	return nil
}
