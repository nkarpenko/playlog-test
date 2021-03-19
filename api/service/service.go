// Package service is the main goroutine of the service
package service

import (
	"context"
	"fmt"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
	"github.com/nkarpenko/playlog-test/api/app"
	"github.com/nkarpenko/playlog-test/api/common/middleware"
	"github.com/nkarpenko/playlog-test/api/common/utils"
	"github.com/nkarpenko/playlog-test/api/conf"
	"github.com/nkarpenko/playlog-test/api/data"
	"github.com/nkarpenko/playlog-test/api/domain/comment"
	"github.com/nkarpenko/playlog-test/api/domain/user"
	"github.com/nkarpenko/playlog-test/api/handler"
	handlerComment "github.com/nkarpenko/playlog-test/api/handler/comment"
	handlerUser "github.com/nkarpenko/playlog-test/api/handler/user"
)

// Service instance
type Service interface {
	Start()
	Stop(ctx context.Context) error
	Config() *conf.Configuration
}

type service struct {
	app         *app.App
	config      *conf.Configuration
	server      *http.Server
	router      *mux.Router
	healthCheck utils.HealthCheckFunc
}

// Start the service
func (s *service) Start() {
	s.registerRoutes()
	go func() {
		addr := fmt.Sprintf("%v:%v", s.config.Host, s.config.Port)
		log.Infof("Listening on %s", addr)
		s.server.ListenAndServe()
	}()
}

// Stop the app
func (s *service) Stop(ctx context.Context) error {
	s.server.Shutdown(ctx)
	s.app.Data.Close()
	return nil
}

// The Config of the service.
func (s *service) Config() *conf.Configuration {
	return s.config
}

// HealthCheck of the whole service
func (s *service) HealthCheck() bool {
	return true
}

// New Service instace.
func New(c *conf.Configuration) (Service, error) {
	s := service{
		app:    &app.App{},
		config: c,
	}

	// Data interface.
	data, err := data.New(&c.Database)
	if err != nil {
		return &s, err
	}

	// Set data interface.
	s.app.Data = data

	// Set domain subservices.
	s.app.CommentService, err = comment.New(s.app.Data)
	if err != nil {
		return &s, fmt.Errorf("unable to create comment domain service: %+v", err)
	}

	s.app.UserService, err = user.New(s.app.Data)
	if err != nil {
		return &s, fmt.Errorf("unable to create user domain service: %+v", err)
	}

	// Router
	s.router = mux.NewRouter()

	// HTTP server
	addr := fmt.Sprintf("%v:%v", s.config.Host, s.config.Port)
	s.server = &http.Server{
		Addr:         addr,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		// Gorilla MUX with applied HTTP Middlewares
		Handler: middleware.Apply(s.router),
	}

	return &s, nil
}

// Register the core/base routes here
func (s *service) registerRoutes() {
	handler.Handlers(s.router, s.config, s.healthCheck)
	handlerUser.Handlers(s.router, s.app)
	handlerComment.Handlers(s.router, s.app)
}
