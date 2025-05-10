package v1

import "time"

type User struct {
	UserID    string    `json:"user_id"`
	Username  string    `json:"username"`
	Nickname  string    `json:"nickname"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	PostCount int64     `json:"post_count"`
	CreateAt  time.Time `json:"create_at"`
	UpdateAt  time.Time `json:"update_at"`
}

type CreateUserRequest struct {
	Username string  `json:"username"`
	Password string  `json:"password"`
	Nickname *string `json:"nickname"`
	Email    string  `json:"email"`
	Phone    string  `json:"phone"`
}

type CreateUserResponse struct {
	UserID string `json:"user_id"`
}

type UpdateUserRequest struct {
	Username *string `json:"username"`
	Nickname *string `json:"nickname"`
	Email    *string `json:"email"`
	Phone    *string `json:"phone"`
}

type UpdateUserResponse struct{}

type DeleteUserRequest struct{}

type DeleteUserResponse struct{}

type GetUserRequest struct {
	UserID string `json:"user_id"`
}

type GetUserResponse struct {
	User *User `json:"user"`
}

type ListUserRequest struct {
	Limit  int64 `json:"limit"`
	Offset int64 `json:"offset"`
}

type ListUserResponse struct {
	Total int64   `json:"total"`
	Users []*User `json:"users"`
}

// ====== login with token ========

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token    string    `json:"token"`
	ExpireAt time.Time `json:"expire_at"`
}

type RefreshTokenRequest struct{}

type RefreshTokenResponse struct {
	Token    string    `json:"token"`
	ExpireAt time.Time `json:"expire_at"`
}

type ChangePasswordRequest struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

type ChangePasswordResponse struct{}
