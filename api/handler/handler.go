// Package handler is a collection of API base endpoint handlers
package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nkarpenko/playlog-test/api/app"
	"github.com/nkarpenko/playlog-test/api/common/request"
	"github.com/nkarpenko/playlog-test/api/common/response"
	"github.com/nkarpenko/playlog-test/api/common/utils"
	"github.com/nkarpenko/playlog-test/api/conf"
)

// Handlers all of request
func Handlers(r *mux.Router, c *conf.Configuration, hcf utils.HealthCheckFunc) {
	r.HandleFunc("/", Info(c))
	r.HandleFunc(utils.HealthCheckURL, HealthCheck(hcf))
}

// ServiceInfo response
type ServiceInfo struct {
	Name    string `json:"name"`
	Desc    string `json:"desc"`
	Version string `json:"version"`
}

// Info request
func Info(c *conf.Configuration) request.HandlerFunc {
	return request.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response.APIResponse(w, &ServiceInfo{
			c.Name, c.Desc, app.Version,
		})
	})
}

// HealthCheck request
func HealthCheck(hcf utils.HealthCheckFunc) request.HandlerFunc {
	return request.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ok, msg := hcf()
		utils.HealthCheckResponse(w, ok, msg)
		return
	})
}
