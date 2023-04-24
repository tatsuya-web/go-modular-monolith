package fixture

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/tatuya-web/go-modular-monolith/modules/user_module/user_model"
)

func User(u *user_model.User) *user_model.User {
	result := &user_model.User{
		ID:        user_model.UserID(rand.Int()),
		Email:     "tatuya" + strconv.Itoa(rand.Int())[:5] + "@example.com",
		Password:  "password",
		Role:      "admin",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if u == nil {
		return result
	}
	if u.ID != 0 {
		result.ID = u.ID
	}
	if u.Email != "" {
		result.Email = u.Email
	}
	if u.Password != "" {
		result.Password = u.Password
	}
	if u.Role != "" {
		result.Role = u.Role
	}
	if !u.CreatedAt.IsZero() {
		result.CreatedAt = u.CreatedAt
	}
	if !u.UpdatedAt.IsZero() {
		result.UpdatedAt = u.UpdatedAt
	}
	return result
}
