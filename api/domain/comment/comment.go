package comment

import (
	"github.com/nkarpenko/playlog-test/api/data"
)

// Service domain name business methods
type Service interface {
}

type service struct {
	name string
	data data.Data
}

// New user service
func New(data data.Data) (Service, error) {
	return &service{
		name: "Comment Subservice",
		data: data,
	}, nil
}
