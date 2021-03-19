package comment

import (
	"database/sql"
)

// Service methods
type Service interface {
	// Get methods.
	GetCommentAll()

	// Create methods.
	CreateComment(user, comment string) error
	CreateCommentLike(user string, commentID int64) error

	// Delete methods.
	DeleteComment(commentID int64) error
}

type service struct {
	name    string
	storage *sql.DB
}

func (s *service) GetCommentAll() {
	return
}

func (s *service) CreateComment(user, name string) error {
	return nil
}

func (s *service) CreateCommentLike(user string, commentID int64) error {
	return nil
}

func (s *service) DeleteComment(commentID int64) error {
	return nil
}

// New business data service
func New(st *sql.DB) Service {
	return &service{
		name:    "Data Layer Comment Service",
		storage: st,
	}
}
