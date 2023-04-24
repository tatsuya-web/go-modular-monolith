package post_repository

import (
	"context"

	"github.com/tatuya-web/go-modular-monolith/modules/post_module/post_model"
	"github.com/tatuya-web/go-modular-monolith/modules/user_module/auth"
	"github.com/tatuya-web/go-modular-monolith/modules/user_module/user_model"
	"github.com/tatuya-web/go-modular-monolith/repository"
)

type PostRepo repository.Repository

func (r *PostRepo) AddPost(ctx context.Context, db repository.Execer, p *post_model.Post) error {
	p.CreatedAt = r.Clocker.Now()
	p.UpdatedAt = r.Clocker.Now()

	sql := `INSERT INTO posts (
		title, content, user_id, created_at, updated_at
	) VALUES (?, ?, ?, ?, ?)`
	result, err := db.ExecContext(ctx, sql, p.Title, p.Content, p.UserID, p.CreatedAt, p.UpdatedAt)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	p.ID = post_model.PostID(id)
	return nil
}

func (r *PostRepo) UpdatePost(ctx context.Context, db repository.Execer, p *post_model.Post) error {
	p.UpdatedAt = r.Clocker.Now()

	sql := `UPDATE posts
								SET title = ?,
								content = ?,
								updated_at = ?
								WHERE id = ?`

	_, err := db.ExecContext(
		ctx, sql, p.Title, p.Content, p.UpdatedAt, p.ID,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *PostRepo) DeletePost(ctx context.Context, db repository.Execer, p *post_model.Post) error {
	sql := `DELETE FROM posts WHERE id = ?`
	_, err := db.ExecContext(
		ctx, sql, p.ID,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *PostRepo) ListPosts(ctx context.Context, db repository.Queryer, id user_model.UserID) (post_model.Posts, error) {
	posts := post_model.Posts{}

	sql := `SELECT
								id, title, content, user_id, created_at, updated_at
								FROM posts
								WHERE user_id = ?`
	if err := db.SelectContext(ctx, &posts, sql, id); err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *PostRepo) IsOwnPost(ctx context.Context, db repository.Queryer, id post_model.PostID) bool {
	ownID, ok := auth.GetUserID(ctx)
	if !ok {
		return false
	}

	var userID int64

	sql := `SELECT
								user_id
								FROM posts
								WHERE id = ?`
	db.QueryRowxContext(ctx, sql, id).Scan(&userID)

	return ownID == user_model.UserID(userID)
}
