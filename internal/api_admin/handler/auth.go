package handler

import (
	"encoding/json"
	"food/internal/api_admin/model"
	"food/pkg/response"
	"net/http"
)

// @Summary Login
// @Tags Auth
// @Description login
// @ID login
// @Accept  json
// @Produce  json
// @Param input body model.Login true "account info"
// @Success 200 {object} model.LoginResponse
// @Failure 400,404 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Failure default {object} response.ErrorResponse
// @Router /login [post]
func (h *Handler) login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := &model.Login{}
		if err := json.NewDecoder(r.Body).Decode(data); err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		responseData, err := h.service.AuthService.Login(data)
		if err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		response.Respond(w, r, http.StatusOK, responseData)
	}
}
