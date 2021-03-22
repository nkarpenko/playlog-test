package comment

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	sb "github.com/huandu/go-sqlbuilder"
	"github.com/nkarpenko/playlog-test/api/data/model"
)

// Service methods
type Service interface {
	// Get methods.
	GetCommentsAll() ([]model.Comment, error)
	GetCommentByID(commentID int64) (model.Comment, error)
	GetCommentLikesByID(commentID int64) ([]model.Like, error)

	// Create methods.
	CreateComment(userID int64, comment string) (model.Comment, error)
	CreateCommentLike(userID, commentID int64) error

	// Delete methods.
	DeleteComment(userID, commentID int64) error
	DeleteCommentLike(userID, commentID int64) error

	// Bool methods.
	IsExistsCommentLike(userID, commentID int64) (bool, error)
}

type service struct {
	name    string
	storage *sql.DB
}

func (s *service) GetCommentsAll() ([]model.Comment, error) {

	// Init comments.
	comments := []model.Comment{}

	// Build the query.
	query := sb.NewSelectBuilder()
	query.Select(
		"id",
		"comment",
		"user_id",
		"created",
		"updated",
	)
	query.From("comment")

	// Execute the query.
	sqlq, args := query.Build()
	rows, err := s.storage.Query(sqlq, args...)
	if err != nil {
		return comments, err
	}

	// Get the comments.
	defer rows.Close()
	for rows.Next() {

		// Get the single comment.
		var comment model.Comment
		err = rows.Scan(
			&comment.ID,
			&comment.Comment,
			&comment.UserID,
			&comment.Created,
			&comment.Updated)
		if err != nil {
			return comments, err
		}

		// Get the comment likes.
		likes, err := s.GetCommentLikesByID(comment.ID)
		if err != nil {
			return comments, err
		}

		// Append the likes
		comment.Likes = likes

		// Append the single comment to the slice.
		comments = append(comments, comment)
	}
	if err := rows.Err(); err != nil {
		return comments, err
	}

	// Successful request.
	return comments, nil
}

func (s *service) GetCommentByID(commentID int64) (model.Comment, error) {
	// Init comment.
	comment := model.Comment{}

	// Build and execute query.
	query := sb.NewSelectBuilder()
	query.Select(
		string("id"),
		string("comment"),
		string("user_id"),
		string("created"),
		string("updated"),
	)
	query.From("comment")
	query.Where(
		query.Equal("id", commentID),
	)
	sqlq, args := query.Build()
	rows, err := s.storage.Query(sqlq, args...)
	if err != nil {
		return comment, err
	}

	// Get the user.
	defer rows.Close()
	rows.Next()
	err = rows.Scan(
		&comment.ID,
		&comment.Comment,
		&comment.UserID,
		&comment.Created,
		&comment.Updated,
	)
	if err != nil {
		return comment, err
	}

	// Get the comment likes.
	likes, err := s.GetCommentLikesByID(comment.ID)
	if err != nil {
		return comment, err
	}

	// Append the likes
	comment.Likes = likes

	// Return the setting.
	return comment, nil
}

func (s *service) GetCommentLikesByID(commentID int64) ([]model.Like, error) {

	// Init likes.
	likes := []model.Like{}

	// Build the query.
	query := sb.NewSelectBuilder()
	query.Select(
		"user_id",
		"comment_id",
		"created",
	)
	query.From("comment_like")

	// Execute the query.
	sqlq, args := query.Build()
	rows, err := s.storage.Query(sqlq, args...)
	if err != nil {
		return likes, err
	}

	// Get the likes.
	defer rows.Close()
	for rows.Next() {

		// Get the single like.
		var like model.Like
		err = rows.Scan(
			&like.UserID,
			&like.CommentID,
			&like.Created)
		if err != nil {
			return likes, err
		}

		// Append the single like to the slice of likes.
		likes = append(likes, like)
	}
	if err := rows.Err(); err != nil {
		return likes, err
	}

	// Successful request.
	return likes, nil
}

func (s *service) CreateComment(userID int64, comment string) (model.Comment, error) {
	// Build query.
	query := sb.NewInsertBuilder()
	query.InsertInto("comment")
	query.Cols(
		string("comment"),
		string("user_id"),
	)
	query.Values(
		userID,
		comment,
	)

	// Execute the query.
	sqlq, args := query.Build()
	res, err := s.storage.Exec(sqlq, args...)
	if err != nil {
		return model.Comment{}, err
	}

	cid, err := res.LastInsertId()
	if err != nil {
		return model.Comment{}, err
	}

	// Successfully inserted.
	return s.GetCommentByID(cid)
}

func (s *service) CreateCommentLike(userID int64, commentID int64) error {
	// Build query.
	query := sb.NewInsertBuilder()
	query.InsertInto("comment_like")
	query.Cols(
		string("user_id"),
		string("comment_id"),
	)
	query.Values(
		userID,
		commentID,
	)

	// Execute the query.
	sqlq, args := query.Build()
	_, err := s.storage.Exec(sqlq, args...)
	if err != nil {
		return err
	}

	// Successfully inserted.
	return nil
}

func (s *service) DeleteComment(userID, commentID int64) error {
	// Build query to delete user setting.
	query := sb.NewDeleteBuilder()
	query.DeleteFrom("comment")
	query.Where(
		query.Equal("id", commentID),
		query.Equal("user_id", userID))

	// Execute the query.
	sqlq, args := query.Build()
	_, err := s.storage.Exec(sqlq, args...)
	if err != nil {
		return err
	}

	// Successfully deleted.
	return nil
}

func (s *service) DeleteCommentLike(userID, commentID int64) error {
	// Build query to delete user setting.
	query := sb.NewDeleteBuilder()
	query.DeleteFrom("comment_like")
	query.Where(
		query.Equal("user_id", userID),
		query.Equal("comment_id", commentID))

	// Execute the query.
	sqlq, args := query.Build()
	_, err := s.storage.Exec(sqlq, args...)
	if err != nil {
		return err
	}

	// Successfully deleted.
	return nil
}

func (s *service) IsExistsCommentLike(userID, commentID int64) (bool, error) {

	// Build query.
	query := sb.NewSelectBuilder()
	query.Select("*")
	query.From("comment_like")
	query.Where(
		query.Equal("comment_id", commentID),
		query.Equal("user_id", userID))

	sqlq, args := query.Build()
	rows, err := s.storage.Query(sqlq, args...)
	if err != nil {
		return false, err
	}

	// Comment like does not exist.
	defer rows.Close()
	if !rows.Next() {
		return false, nil
	}

	// Comment like exists.
	return true, nil
}

// New business data service
func New(st *sql.DB) Service {
	return &service{
		name:    "Data Layer Comment Service",
		storage: st,
	}
}
