package app

import (
	"github.com/nkarpenko/playlog-test/api/data"
	"github.com/nkarpenko/playlog-test/api/domain/comment"
	"github.com/nkarpenko/playlog-test/api/domain/user"
)

const (
	// Version of the service
	Version string = "0.1.0"
)

// An App hold references to all subservices
type App struct {
	Data data.Data

	CommentService comment.Service
	UserService    user.Service
}
