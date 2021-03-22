package comment

import (
	"github.com/nkarpenko/playlog-test/api/data"
	"github.com/nkarpenko/playlog-test/api/data/model"
)

// Service domain name business methods
type Service interface {

	// Get methods
	GetCommentsAll() ([]model.Comment, error)

	// Create methods
	CreateComment(userID int64, comment string) (model.Comment, error)
	CreateCommentLike(userID, commentID int64) error

	// Delete methods.
	DeleteComment(userID, commentID int64) error
}

type service struct {
	name string
	data data.Data
}

func (s *service) GetCommentsAll() ([]model.Comment, error) {
	return s.data.Comment().GetCommentsAll()
}

func (s *service) CreateComment(userID int64, comment string) (model.Comment, error) {
	return s.data.Comment().CreateComment(userID, comment)
}

func (s *service) CreateCommentLike(userID, commentID int64) error {
	// Validate that the like already exists.
	ok, err := s.data.Comment().IsExistsCommentLike(userID, commentID)
	if err != nil {
		return err
	}

	// Create user if they do not exist.
	if !ok {
		return s.data.Comment().CreateCommentLike(userID, commentID)
	}

	// Else delete the like since that means they unliked the comment.
	return s.data.Comment().DeleteCommentLike(userID, commentID)
}

func (s *service) DeleteComment(userID, commentID int64) error {
	return s.data.Comment().DeleteComment(userID, commentID)
}

// New user service
func New(data data.Data) (Service, error) {
	return &service{
		name: "Comment Subservice",
		data: data,
	}, nil
}
