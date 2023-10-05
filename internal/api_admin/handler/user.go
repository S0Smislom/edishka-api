package handler

import (
	_ "food/internal/api_admin/model"
	"food/pkg/response"
	"net/http"
)

// @Summary Me
// @Tags User
// @Security ApiKeyAuth
// @Description Return current user
// @ID current-user
// @Produce  json
// @Success 200 {object} model.User
// @Failure 400,404 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Failure default {object} response.ErrorResponse
// @Router /v1/user/me [get]
func (h *Handler) me() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := h.service.UserService.GetById(r.Context().Value(userCtx).(int))
		if err != nil {
			response.ErrorRespond(w, r, err)
		}
		response.Respond(w, r, http.StatusOK, user)
	}
}
