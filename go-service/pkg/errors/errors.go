package errors

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/sevengit-wq/skill-test/pkg/models"
)

type ErrorHandler interface {
	Error() string
	Code() int
	Message() string
}

type CustomError struct {
	ErrorCode    int
	ErrorMessage string
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("code: %d, message: %s", e.ErrorCode, e.ErrorMessage)
}

func (e *CustomError) Code() int {
	return e.ErrorCode
}

func (e *CustomError) Message() string {
	return e.ErrorMessage
}

func ValidationError(body []byte, statusCode int) ErrorHandler {
	var validationErr models.ValidationErrorResponse
	if err := json.Unmarshal(body, &validationErr); err != nil {
		return &CustomError{
			ErrorCode:    http.StatusBadRequest,
			ErrorMessage: "failed to parse validation error: " + err.Error(),
		}
	}

	var detailMsgs []string
	for _, d := range validationErr.Detail {
		detailMsgs = append(detailMsgs, fmt.Sprintf("%s: %s", d.Path, d.Message))
	}

	return &CustomError{
		ErrorCode:    statusCode,
		ErrorMessage: fmt.Sprintf("%s: %s", validationErr.Error, strings.Join(detailMsgs, ", ")),
	}
}
