package handler

import (
	"context"
	"food/internal/api/model"
	"food/pkg/exceptions"
	"food/pkg/response"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (h *Handler) AuthenticateUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get(authorizationHeader)
		claims, err := h.parseAuthHeader(header)
		if err != nil {
			response.ErrorRespond(w, r, &exceptions.UnauthorizedError{Msg: "invalid auth header"})
			return
		}
		if claims.TokenType != model.AccessTokenType {
			response.ErrorRespond(w, r, &exceptions.UnauthorizedError{Msg: "Wrong token type"})
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
		return nil, &exceptions.UnauthorizedError{Msg: "Empty auth header"}
	}
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return nil, &exceptions.UnauthorizedError{Msg: "Invalide auth header"}
	}
	if len(headerParts[1]) == 0 {
		return nil, &exceptions.UnauthorizedError{Msg: "Token in empty"}
	}
	claims, err := h.services.Auth.ParseToken(headerParts[1])
	if err != nil {
		return nil, err
	}
	return claims, nil
}
