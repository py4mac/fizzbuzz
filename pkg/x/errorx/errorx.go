package errorx

import (
	"fmt"

	"github.com/pkg/errors"
)

// Error ...
type Error struct {
	// Cause original error
	Cause error `json:"-"`

	// Message from this error
	Message string `json:"message"`
}

// GetMessage returns the message from Error
func (e *Error) GetMessage() string {
	return e.Message
}

// GetCause returns the cause from Error
func (e *Error) GetCause() error {
	return errors.Cause(e.Cause)
}

// Wrap the error with Error type
func Wrap(err error, message string) error {
	return &Error{Cause: errors.Wrap(err, message), Message: message}
}

// Error return a custom message from Error
func (e *Error) Error() string {
	if e.Cause == nil {
		return e.Message
	}

	if e.Message == "" {
		return fmt.Sprintf("%s", e.GetCause())
	}

	return fmt.Sprintf("%s: %s", e.GetMessage(), e.GetCause())
}
