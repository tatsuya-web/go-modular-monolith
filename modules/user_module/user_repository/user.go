package user_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-sql-driver/mysql"
	"github.com/tatuya-web/go-modular-monolith/modules/user_module/user_model"
	"github.com/tatuya-web/go-modular-monolith/repository"
)

type UserRepo repository.Repository

func (r *UserRepo) RegisterUser(ctx context.Context, db repository.Execer, u *user_model.User) error {
	u.CreatedAt = r.Clocker.Now()
	u.UpdatedAt = r.Clocker.Now()

	sql := `INSERT INTO users (
			email, password, role, created_at, updated_at
			) VALUES (?, ?, ?, ?, ?)`
	result, err := db.ExecContext(ctx, sql, u.Email, u.Password, u.Role, u.CreatedAt, u.UpdatedAt)
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == repository.ErrCodeMySQLDuplicateEntry {
			return fmt.Errorf("cannot create same email user: %w", repository.ErrAlreadyEntry)
		}
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	u.ID = user_model.UserID(id)
	return nil
}

func (r *UserRepo) GetUser(ctx context.Context, db repository.Queryer, email string) (*user_model.User, error) {
	u := &user_model.User{}
	sql := `SELECT
			id, email, password, role, created_at, updated_at 
			FROM users WHERE email = ?`
	if err := db.GetContext(ctx, u, sql, email); err != nil {
		return nil, err
	}
	return u, nil
}

func (r *UserRepo) GetOwn(ctx context.Context, db repository.Queryer, id user_model.UserID) (*user_model.User, error) {
	u := &user_model.User{}

	sql := `SELECT
			id, email, role, created_at, updated_at
			FROM users
			WHERE id = ?`
	if err := db.GetContext(ctx, u, sql, id); err != nil {
		return nil, err
	}
	return u, nil
}

func (r *UserRepo) UpdateUser(ctx context.Context, db repository.Execer, u *user_model.User) error {
	u.UpdatedAt = r.Clocker.Now()

	sql := `UPDATE users
			SET email = ?,
			updated_at = ?
			WHERE id = ?`
	_, err := db.ExecContext(
		ctx, sql, u.Email, u.UpdatedAt, u.ID,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepo) DeleteUser(ctx context.Context, db repository.Execer, u *user_model.User) error {
	sql := `DELETE FROM users WHERE id = ?`
	_, err := db.ExecContext(
		ctx, sql, u.ID,
	)
	if err != nil {
		return err
	}

	return nil
}
