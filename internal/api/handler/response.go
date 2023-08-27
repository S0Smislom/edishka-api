package handler

import (
	"encoding/json"
	"net/http"
)

type responseWriter struct {
	http.ResponseWriter
	Code int
}

type errorResponse struct {
	Detail string `json:"detail"`
}

func (w *responseWriter) WriteHeader(statusCode int) {
	w.Code = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func (h *Handler) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

func (h *Handler) errorRespond(w http.ResponseWriter, r *http.Request, code int, err error) {
	h.respond(w, r, code, errorResponse{Detail: err.Error()})
}
