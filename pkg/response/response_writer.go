package response

import (
	"encoding/json"
	"errors"
	"food/pkg/exceptions"
	"net/http"
)

type ResponseWriter struct {
	http.ResponseWriter
	Code int
}

type ErrorResponse struct {
	Type   string `json:"type"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

func (w *ResponseWriter) WriteHeader(statusCode int) {
	w.Code = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func Respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

func ErrorRespond(w http.ResponseWriter, r *http.Request, err error) {
	code, response := handleError(err)
	Respond(w, r, code, response)
}

func handleError(err error) (int, ErrorResponse) {
	var objectNotFoundError *exceptions.ObjectNotFoundError
	var userPermissionError *exceptions.UserPermissionError
	var duplicateError *exceptions.DuplicateError
	var logicError *exceptions.LogicError
	var unauthorizedError *exceptions.UnauthorizedError
	var wrongPasswordError *exceptions.WrongPasswordError
	var validationError *exceptions.ValidationError

	switch {
	case errors.As(err, &objectNotFoundError):
		return http.StatusNotFound, ErrorResponse{"Value Error", "Not Found", err.Error()}
	case errors.As(err, &userPermissionError):
		return http.StatusForbidden, ErrorResponse{"Permission Error", "Forbidden", err.Error()}
	case errors.As(err, &duplicateError):
		return http.StatusBadRequest, ErrorResponse{"Value Error", "Duplicate", err.Error()}
	case errors.As(err, &logicError):
		return http.StatusForbidden, ErrorResponse{"Logic Error", "Logic Error", err.Error()}
	case errors.As(err, &unauthorizedError):
		return http.StatusUnauthorized, ErrorResponse{"Permission Error", "Unauthorized", err.Error()}
	case errors.As(err, &wrongPasswordError):
		return http.StatusUnauthorized, ErrorResponse{"Permission Error", "Wrong Password", err.Error()}
	case errors.As(err, &validationError):
		return http.StatusUnprocessableEntity, ErrorResponse{"Validation Error", "Unprocessable Entity", err.Error()}
	default:
		return http.StatusInternalServerError, ErrorResponse{"Server Error", "Internal Server Error", err.Error()}
	}
}
