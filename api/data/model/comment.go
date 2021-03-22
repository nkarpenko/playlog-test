package model

import "time"

type Comment struct {
	ID      int64     `json:"id"`
	Comment string    `json:"comment"`
	UserID  int64     `json:"user_id"`
	Created time.Time `json:"created,omitempty"`
	Updated time.Time `json:"updated,omitempty"`
	Likes   []Like    `json:"likes,omitempty"`
}

type Like struct {
	UserID    int64     `json:"user_id"`
	CommentID int64     `json:"comment_id,omitempty"`
	Created   time.Time `json:"created,omitempty"`
}
