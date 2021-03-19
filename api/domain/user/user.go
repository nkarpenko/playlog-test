package user

import (
	"github.com/nkarpenko/playlog-test/api/data"
	"github.com/nkarpenko/playlog-test/api/data/model"
)

// Service domain name business methods
type Service interface {
	EnsureUserExists(name string) (model.User, error)
}

type service struct {
	name string
	data data.Data
}

func (s *service) EnsureUserExists(name string) (model.User, error) {

	// Validate that the user already exists.
	ok, err := s.data.User().IsExistsUser(name)
	if err != nil {
		return model.User{}, err
	}

	// Create user if they do not exist.
	if !ok {
		return s.data.User().CreateUser(name)
	}

	// Else return the user
	return s.data.User().GetUserByName(name)
}

// New user service
func New(data data.Data) (Service, error) {
	return &service{
		name: "User Subservice",
		data: data,
	}, nil
}
