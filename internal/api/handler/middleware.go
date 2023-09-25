package handler

import (
	"context"
	"errors"
	"food/pkg/response"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (h *Handler) authenticateUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get(authorizationHeader)
		if header == "" {
			response.ErrorRespond(w, r, http.StatusInternalServerError, errors.New("empty auth header"))
			return
		}
		headerParts := strings.Split(header, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			response.ErrorRespond(w, r, http.StatusInternalServerError, errors.New("invalid auth header"))
			return
		}
		if len(headerParts[1]) == 0 {
			response.ErrorRespond(w, r, http.StatusInternalServerError, errors.New("token is empty"))
			return
		}
		userId, err := h.services.Auth.ParseToken(headerParts[1])
		if err != nil {
			response.ErrorRespond(w, r, http.StatusInternalServerError, errors.New("invalid auth header"))
			return
		}
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), userCtx, userId)))
	})
}

func (h *Handler) optionalAuthenticateUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get(authorizationHeader)
		if header == "" {
			next.ServeHTTP(w, r)
			return
		}
		headerParts := strings.Split(header, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			next.ServeHTTP(w, r)
			return
		}
		if len(headerParts[1]) == 0 {
			next.ServeHTTP(w, r)
			return
		}
		userId, err := h.services.Auth.ParseToken(headerParts[1])
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), userCtx, userId)))
	})
}
