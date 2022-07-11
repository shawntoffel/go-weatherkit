package weatherkit

import (
	"fmt"
	"net/http"
	"time"
)

// ErrorResponse is returned in response to an API error.
type ErrorResponse struct {
	Timestamp *time.Time `json:"timestamp,omitempty"`
	Status    int        `json:"status,omitempty"`
	Error     string     `json:"error,omitempty"`
	Message   string     `json:"message,omitempty"`
	Path      string     `json:"path,omitempty"`
}

type RestError struct {
	Response      *http.Response
	ErrorResponse *ErrorResponse
}

func (e *RestError) Error() string {
	return fmt.Sprintf("http: status code: %d %s %s", e.Response.StatusCode, http.StatusText(e.Response.StatusCode), e.ErrorResponse.Message)
}
