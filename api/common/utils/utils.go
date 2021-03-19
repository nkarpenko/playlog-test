// Package utils is a collection of HTTP utilities
package utils

import (
	"net/http"

	"github.com/nkarpenko/playlog-test/api/common/response"
)

// HealthCheckURL used for all service
const HealthCheckURL = "/health"

type healthCheckStatus string

// Normalized health check response status
const (
	healthCheckStatusOK    healthCheckStatus = "ok"
	healthCheckStatusError healthCheckStatus = "error"
)

// Normalized health check response
type healthCheckResponse struct {
	Status  healthCheckStatus `json:"status" apitype:"string"`
	Message string            `json:"msg,omitempty"`
}

// HealthCheckResponse in a normalized format
// if ok, the response is {"status": "ok"}
// otherwise, the response is {"status": "error", "msg": "status message"}
// The error state is a "soft-error" state, the service is not down, it
// still responds but some parts might be down
func HealthCheckResponse(w http.ResponseWriter, ok bool, msg string) {
	if ok {
		response.JSON(w, 200, &healthCheckResponse{healthCheckStatusOK, ""})
	} else {
		response.JSON(w, 200, &healthCheckResponse{healthCheckStatusError, msg})
	}
}

// HealthCheckFunc for normalized heath check methods
type HealthCheckFunc func() (bool, string)
