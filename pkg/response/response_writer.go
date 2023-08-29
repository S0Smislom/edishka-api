package response

import (
	"encoding/json"
	"net/http"
)

type ResponseWriter struct {
	http.ResponseWriter
	Code int
}

type ErrorResponse struct {
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

func ErrorRespond(w http.ResponseWriter, r *http.Request, code int, err error) {
	Respond(w, r, code, ErrorResponse{Detail: err.Error()})
}
