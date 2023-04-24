package post_handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tatuya-web/go-modular-monolith/handler"
	"github.com/tatuya-web/go-modular-monolith/modules/post_module/post_model"
)

type ListPost struct {
	Service ListPostService
}

type post struct {
	ID        post_model.PostID `json:"id"`
	Title     string            `json:"title"`
	Content   string            `json:"content"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
}

func NewListPostHandler(lps ListPostService) *ListPost {
	return &ListPost{Service: lps}
}

func (lp *ListPost) ServeHTTP(ctx *gin.Context) {
	posts, err := lp.Service.ListPosts(ctx)
	if err != nil {
		handler.ErrResponse(ctx, http.StatusInternalServerError, "faild to post", err.Error())
		return
	}

	rsp := []post{}
	for _, t := range posts {
		rsp = append(rsp, post{
			ID:        t.ID,
			Title:     t.Title,
			Content:   t.Content,
			CreatedAt: t.CreatedAt,
			UpdatedAt: t.UpdatedAt,
		})
	}

	handler.APIResponse(ctx, http.StatusOK, "post一覧", rsp)
}
