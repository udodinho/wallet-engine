package helpers

import (
	"fmt"
	"github.com/go-playground/validator"
	"github.com/google/uuid"
	"strings"
	"time"
)

type FieldError struct {
	err validator.FieldError
}

func (f FieldError) String() string {
	var sb strings.Builder

	sb.WriteString("validation failed on field '" + f.err.Field() + "'")
	sb.WriteString(", condition: " + f.err.ActualTag())

	// Print condition parameters, e.g. one_of=red blue -> { red blue }
	if f.err.Param() != "" {
		sb.WriteString(" { " + f.err.Param() + " }")
	}

	if f.err.Value() != nil && f.err.Value() != "" {
		sb.WriteString(fmt.Sprintf(", actual: %v", f.err.Value()))
	}

	return sb.String()
}

func NewFieldError(err validator.FieldError) FieldError {
	return FieldError{err: err}
}

func (err ErrorResponse) Error() string {
	var errorBody ErrorBody
	return fmt.Sprintf("%v", errorBody)
}

func PrintErrorMessage(code, message string) error {
	var errorResponse ErrorResponse
	errorResponse.TimeStamp = time.Now().Format(time.RFC3339)
	errorResponse.ErrorReference = uuid.New()
	errorResponse.Errors = append(errorResponse.Errors, ErrorBody{
		Code:    code,
		Message: message,
		Source:  Config.AppName,
	})
	return errorResponse
}
