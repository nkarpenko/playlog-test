package user

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	sb "github.com/huandu/go-sqlbuilder"
	"github.com/nkarpenko/playlog-test/api/data/model"
)

// Service methods
type Service interface {
	// Create methods.
	CreateUser(name string) (model.User, error)

	// Get methods.
	GetUserByName(name string) (model.User, error)

	// Bool methods.
	IsExistsUser(name string) (bool, error)
}

type service struct {
	name    string
	storage *sql.DB
}

func (s *service) CreateUser(name string) (model.User, error) {
	// Build and execute query to create new user.
	query := sb.NewInsertBuilder()
	query.InsertInto("user")
	query.Cols(
		string("name"),
	)
	query.Values(
		name,
	)

	// Execute the query.
	sqlq, args := query.Build()
	_, err := s.storage.Exec(sqlq, args...)
	if err != nil {
		return model.User{}, err
	}

	// Successfully inserted.
	return s.GetUserByName(name)
}

func (s *service) GetUserByName(name string) (model.User, error) {
	// Init USER.
	var user model.User

	// Build and execute query to get `setting` from `user_setting` table.
	query := sb.NewSelectBuilder()
	query.Select(
		string("id"),
		string("name"),
	)
	query.From("user")
	query.Where(
		query.Equal("name", name),
	)
	sqlq, args := query.Build()
	rows, err := s.storage.Query(sqlq, args...)
	if err != nil {
		return user, err
	}

	// Get the setting.
	defer rows.Close()
	rows.Next()
	err = rows.Scan(
		&user.ID,
		&user.Name,
	)
	if err != nil {
		return user, err
	}

	// Return the setting.
	return user, nil
}

func (s *service) IsExistsUser(name string) (bool, error) {

	query := sb.NewSelectBuilder()
	query.Select("*")
	query.From("user")
	query.Where(
		query.Equal("name", name))
	//query.Limit(1)

	sqlq, args := query.Build()
	rows, err := s.storage.Query(sqlq, args...)
	if err != nil {
		return false, err
	}

	// User does exist.
	defer rows.Close()
	if !rows.Next() {
		return false, nil
	}

	// User has settings.
	return true, nil
}

// New business data service
func New(st *sql.DB) Service {
	return &service{
		name:    "Data Layer User Service",
		storage: st,
	}
}
