package handler

import (
	"encoding/json"
	"food/internal/api/model"
	"food/pkg/exceptions"
	"food/pkg/response"
	"net/http"
)

// @Summary SignUp
// @Tags Auth
// @Description create account
// @ID create-account
// @Accept  json
// @Produce  json
// @Param input body model.Login true "account info"
// @Success 200 {object} model.LoginResponse
// @Failure 400,404 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Failure default {object} response.ErrorResponse
// @Router /login [post]
func (h *Handler) logIn() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := &model.Login{}
		if err := json.NewDecoder(r.Body).Decode(data); err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		user, err := h.services.Auth.CreateUser(data)
		if err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		response.Respond(w, r, http.StatusOK, user)
	}
}

// @Summary Confirm Code
// @Tags Auth
// @Description Generate access/refresh tokens
// @ID generate-tokens
// @Accept  json
// @Produce  json
// @Param input body model.LoginConfirm true "account info"
// @Success 200 {object} model.LoginConfirmResponse
// @Failure 400,404 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Failure default {object} response.ErrorResponse
// @Router /login/confirm [post]
func (h *Handler) confirmCode() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := &model.LoginConfirm{}
		if err := json.NewDecoder(r.Body).Decode(data); err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		response_data, err := h.services.Auth.Login(data)
		if err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		response.Respond(w, r, http.StatusOK, response_data)
	}
}

// @Summary Refresh token
// @Tags Auth
// @Security ApiKeyAuth
// @Description Generate new access/refresh tokens
// @ID refresh-tokens
// @Produce json
// @Success 200 {object} model.LoginConfirmResponse
// @Failure default {object} response.ErrorResponse
// @Router /login/refresh [post]
func (h *Handler) refreshTokenHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get(authorizationHeader)
		claims, err := h.parseAuthHeader(header)
		if err != nil {
			response.ErrorRespond(w, r, &exceptions.UserPermissionError{Msg: "invalid auth header"})
			return
		}
		if claims.TokenType != model.RefreshTokenType {
			response.ErrorRespond(w, r, &exceptions.UserPermissionError{Msg: "Wrong token type"})
			return
		}
		response_data, err := h.services.Auth.Refresh(claims.UserId)
		if err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		response.Respond(w, r, http.StatusOK, response_data)

	}
}
