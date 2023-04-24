package post_handler

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/tatuya-web/go-modular-monolith/handler"
	"github.com/tatuya-web/go-modular-monolith/modules/post_module/post_model"
)

type DeletePost struct {
	Service   DeletePostService
	Validator *validator.Validate
}

func NewDeletePostHandler(dp DeletePostService, v *validator.Validate) *DeletePost {
	return &DeletePost{Service: dp, Validator: v}
}

func (dp *DeletePost) ServeHTTP(ctx *gin.Context) {
	var p struct {
		ID int64 `json:"id" validate:"required"`
	}

	if err := json.NewDecoder(ctx.Request.Body).Decode(&p); err != nil {
		handler.ErrResponse(ctx, http.StatusInternalServerError, "faild to post", err.Error())
		return
	}

	err := dp.Validator.Struct(p)
	if err != nil {
		handler.ErrResponse(ctx, http.StatusBadRequest, "faild to post", err.Error())
		return
	}

	if err := dp.Service.DeletePost(ctx, post_model.PostID(p.ID)); err != nil {
		handler.ErrResponse(ctx, http.StatusInternalServerError, "faild to post", err.Error())
		return
	}

	handler.APIResponse(ctx, http.StatusOK, "postを削除しました。", nil)
}
