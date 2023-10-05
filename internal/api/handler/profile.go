package handler

import (
	"encoding/json"
	"food/internal/api/model"
	"food/pkg/response"
	"net/http"
)

// @Summary Get current user
// @Tags Profile
// @Security ApiKeyAuth
// @Description Get current user
// @ID get-current-user
// @Produce json
// @Success 200 {object} model.User
// @Failure default {object} response.ErrorResponse
// @Router /v1/profile [get]
func (h *Handler) getCurrentProfileHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		currentUserId := r.Context().Value(userCtx).(int)
		dbUser, err := h.services.User.GetById(currentUserId)
		if err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		response.Respond(w, r, http.StatusOK, dbUser)
	}
}

// @Summary Update current profile
// @Tags Profile
// @Security ApiKeyAuth
// @Description Update current user
// @ID update-current-user
// @Accept json
// @Produce json
// @Param input body model.UpdateUser true "User udpate data"
// @Success 200 {object} model.User
// @Failure default {object} response.ErrorResponse
// @Router /v1/profile [patch]
func (h *Handler) updateProfileHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		currentUserId := r.Context().Value(userCtx).(int)
		data := &model.UpdateUser{}
		if err := json.NewDecoder(r.Body).Decode(data); err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		dbUser, err := h.services.User.Update(currentUserId, data)
		if err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		response.Respond(w, r, http.StatusOK, dbUser)
	}
}

// @Summary Upload photo
// @Tags Profile
// @Security ApiKeyAuth
// @Description Upload profile photo
// @ID upload-profile-photo
// @Accept			multipart/form-data
// @Produce  json
// @Param	photo formData file	 true "this is a test file"
// @Success 200 {object} model.User
// @Failure default {object} response.ErrorResponse
// @Router /v1/profile/photo [post]
func (h *Handler) uploadProfilePhotoHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		currentUserId := r.Context().Value(userCtx).(int)
		file, fileHeader, err := r.FormFile("photo")
		if err != nil {
			response.Respond(w, r, http.StatusInternalServerError, err)
		}
		defer file.Close()
		dbUser, err := h.services.User.UploadPhoto(currentUserId, file, fileHeader)
		if err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		response.Respond(w, r, http.StatusOK, dbUser)
	}
}

// @Summary Delete photo
// @Tags Profile
// @Security ApiKeyAuth
// @Description Delete profile photo
// @ID delete-profile-photo
// @Produce  json
// @Success 200 {object} model.User
// @Failure default {object} response.ErrorResponse
// @Router /v1/profile/photo [delete]
func (h *Handler) deleteProfilePhotoHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		currentUserId := r.Context().Value(userCtx).(int)
		dbUser, err := h.services.User.DeletePhoto(currentUserId)
		if err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		response.Respond(w, r, http.StatusOK, dbUser)
	}
}
