package errors

import (
	"fmt"
	"net/http"
)

// RestErr defines REST error
type RestErr struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Err     string `json:"error"`
}

// Error implements error interface
func (r RestErr) Error() string {
	return fmt.Sprintf("message: %s; status: %d; error: %s",
		r.Message, r.Status, r.Err)
}

// BadRequest defines bad request error
func BadRequest(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusBadRequest,
		Err:     "bad_request",
	}
}

// NotFoundError defines not found error
func NotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusNotFound,
		Err:     "not_found",
	}
}
