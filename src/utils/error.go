package utils

import (
	"net/http"

	"github.com/jhonnydsl/clinify-backend/src/dtos"
)

func NotFoundError(message string) *dtos.APIError {
	return &dtos.APIError{StatusCode: http.StatusNotFound, Message: message}
}

func BadRequestError(message string) *dtos.APIError {
	return &dtos.APIError{StatusCode: http.StatusBadRequest, Message: message}
}

func ConflictError(message string) *dtos.APIError {
	return &dtos.APIError{StatusCode: http.StatusConflict, Message: message}
}

func InternalServerError(message string) *dtos.APIError {
	return &dtos.APIError{StatusCode: http.StatusInternalServerError, Message: message}
}

func GetStatusCode(err error) int {
	apiErr, ok := err.(*dtos.APIError)
	if ok {
		return apiErr.StatusCode
	}
	return http.StatusInternalServerError
}