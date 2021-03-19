package response

import "fmt"

// APIError Normalized struct for an API error
type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	// HTTP Status Code
	Status int `json:"-"`
}

// Error message
func (e APIError) Error() string {
	return fmt.Sprintf("[%d]: %s", e.Code, e.Message)
}

// SetMessage returns a new API error based on the
// current one, with updated error message
func (e APIError) SetMessage(msg string) APIError {
	err := APIError(e)
	err.Message = msg
	return err
}

// NewAPIErr init
func NewAPIErr(code int, status int, message string) error {
	return &APIError{Code: code, Message: message, Status: status}
}
