package post_model

import (
	"time"

	"github.com/tatuya-web/go-modular-monolith/modules/user_module/user_model"
)

type PostID int64

type Post struct {
	ID        PostID            `json:"id" db:"id"`
	Title     string            `json:"title" db:"title"`
	Content   string            `json:"content" db:"content"`
	UserID    user_model.UserID `json:"user_id" db:"user_id"`
	CreatedAt time.Time         `json:"created_at" db:"created_at"`
	UpdatedAt time.Time         `json:"updated_at" db:"updated_at"`
}

type Posts []*Post
