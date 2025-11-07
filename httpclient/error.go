package httpclient

import (
	goerrors "errors"

	"github.com/lynx-go/x/errors"
)

// Error represents an HTTP client error
type Error struct {
	*errors.InternalError
	StatusCode int
	Response   []byte
}

func (e *Error) Unwrap() error {
	return e.InternalError.Unwrap()
}

func (e *Error) Error() string {
	return e.InternalError.Error()
}

// NewError creates a new HTTP client error
func NewError(statusCode int, response []byte) *Error {
	return &Error{
		InternalError: &errors.InternalError{
			Code:    errors.ErrCodeHTTPClient,
			Message: "http client error",
			Err:     errors.NewError("http client error").Mark(errors.ErrHTTPClient),
		},
		StatusCode: statusCode,
		Response:   response,
	}
}

// IsHTTPError checks if an error is an HTTP client error
func IsHTTPError(err error) (*Error, bool) {
	var httpErr *Error
	if goerrors.As(err, &httpErr) {
		return httpErr, true
	}
	return nil, false
}
