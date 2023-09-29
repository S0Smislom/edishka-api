package handler

import (
	"context"
	"errors"
	"food/internal/api/model"
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
		claims, err := h.parseAuthHeader(header)
		if err != nil {
			response.ErrorRespond(w, r, http.StatusInternalServerError, errors.New("invalid auth header"))
			return
		}
		if claims.TokenType != model.AccessTokenType {
			response.ErrorRespond(w, r, http.StatusUnauthorized, errors.New("Wrong token type"))
			return
		}
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), userCtx, claims.UserId)))
	})
}

func (h *Handler) optionalAuthenticateUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get(authorizationHeader)

		claims, err := h.parseAuthHeader(header)
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}
		if claims.TokenType != model.AccessTokenType {
			next.ServeHTTP(w, r)
			return
		}
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), userCtx, claims.UserId)))
	})
}

func (h *Handler) parseAuthHeader(header string) (*model.TokenClaims, error) {
	if header == "" {
		return nil, errors.New("Empty auth header")
	}
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return nil, errors.New("Invalide auth header")
	}
	if len(headerParts[1]) == 0 {
		return nil, errors.New("Token in empty")
	}
	claims, err := h.services.Auth.ParseToken(headerParts[1])
	if err != nil {
		return nil, err
	}
	return claims, nil
}
