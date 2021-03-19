package response

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// JSON response
func JSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// InternalServerError response with and error logged into the console
func InternalServerError(w http.ResponseWriter, err error) {
	log.WithError(err).Error("Internal error")
	w.WriteHeader(http.StatusInternalServerError)
}

type apiStatus string

// Normalized API response status
// ok status has data prop in response
// error status has error object in response
const (
	APIStatusOK    apiStatus = "ok"
	APIStatusError apiStatus = "error"
)

// APIResp is a normalized API response
// data and error is auto omitted when empty
type APIResp struct {
	Status apiStatus   `json:"status"`
	Data   interface{} `json:"data,omitempty"`
	Error  *APIError   `json:"error,omitempty"`
}

// Data is a normalized data wrapper
// This is just a helper struct for API documentation
// It represents the resolved APIResp struct for
// API success response.
type Data struct {
	Status apiStatus   `json:"status" apitype:"string"`
	Data   interface{} `json:"data"`
}

// OK is a normalized data wrapper
// This is just a helper struct for API documentation
// It represents the resolved APIResp struct for
// API OK response.
type OK struct {
	Status apiStatus `json:"status" apitype:"string"`
}

// Error is a normalized data wrapper
// This is just a helper struct for API documentation
// It represents the resolved APIResp struct for
// API error response.
type Error struct {
	Status apiStatus `json:"status" apitype:"string"`
	Error  *APIError `json:"error"`
}

// APIResponse in a normalized form
//
//  Response format:
//  {
// 	 "status": "ok"
// 	 "data": [],{}
//  }
func APIResponse(w http.ResponseWriter, data interface{}) {
	JSON(w, http.StatusOK, APIResp{
		Status: APIStatusOK,
		Data:   data,
	})
}

// APIResponseError in a normalized form
//
// err param is an optional internal error
// which will be logged before the actual API Error
//
//  Response format:
//  {
// 	 "status": "error"
// 	 "error": [],{}
//  }
func APIResponseError(w http.ResponseWriter, apiErr APIError, err error) {
	if err != nil {
		log.Error(err)
	}
	entry := log.WithFields(log.Fields{
		"code": apiErr.Code,
	})
	entry.Error(apiErr.Message)
	JSON(w, apiErr.Status, APIResp{
		Status: APIStatusError,
		Error:  &apiErr,
	})
}
