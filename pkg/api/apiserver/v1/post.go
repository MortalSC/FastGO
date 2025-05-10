package v1

import "time"

type Post struct {
	PostID   string    `json:"post_id"`
	UserID   string    `json:"user_id"`
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}

type CreatePostRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type CreatePostResponse struct {
	PostID string `json:"post_id"`
}

type UpdatePostRequest struct {
	PostID  string  `json:"post_id"`
	Title   *string `json:"title"`
	Content *string `json:"content"`
}

type UpdatePostResponse struct{}

type DeletePostRequest struct {
	PostID []string `json:"post_id"`
}

type DeletePostResponse struct{}

type GetPostRequest struct {
	PostID string `json:"post_id" uri:"post_id"`
}

type GetPostResponse struct {
	Post *Post `json:"post"`
}

type ListPostRequest struct {
	Limit  int64   `json:"limit"`
	Offset int64   `json:"offset"`
	Title  *string `json:"title"`
}

type ListPostResponse struct {
	Total int64   `json:"total"`
	Posts []*Post `json:"posts"`
}
