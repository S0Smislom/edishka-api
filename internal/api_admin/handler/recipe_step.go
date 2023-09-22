package handler

import (
	"encoding/json"
	"food/internal/api_admin/model"
	"food/pkg/response"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

// @Summary Create RecipeStep
// @Tags RecipeStep
// @Security ApiKeyAuth
// @Description Create RecipeStep
// @ID create-recipe-step
// @Accept  json
// @Produce  json
// @Param input body model.CreateRecipeStep true "RecipeStep info"
// @Success 200 {object} model.RecipeStep
// @Failure default {object} response.ErrorResponse
// @Router /v1/recipe-step [post]
func (h *Handler) handlerCreateRecipeStep() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &model.CreateRecipeStep{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			response.ErrorRespond(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		RecipeStep, err := h.service.RecipeStepService.Create(req)
		if err != nil {
			response.ErrorRespond(w, r, http.StatusBadRequest, err)
			return
		}
		response.Respond(w, r, http.StatusOK, RecipeStep)
	}
}

// @Summary Get RecipeStep by id
// @Tags RecipeStep
// @Security ApiKeyAuth
// @Description Get RecipeStep by id
// @ID get-recipe-step
// @Accept json
// @Produce  json
// @Param id path int true "RecipeStep id"
// @Success 200 {object} model.RecipeStep
// @Failure default {object} response.ErrorResponse
// @Router /v1/recipe-step/{id} [get]
func (h *Handler) handlerGetRecipeStepById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			response.ErrorRespond(w, r, http.StatusOK, err)
			return
		}
		RecipeStep, err := h.service.RecipeStepService.GetById(id)
		if err != nil {
			response.ErrorRespond(w, r, http.StatusNotFound, err)
			return
		}
		response.Respond(w, r, http.StatusOK, RecipeStep)
	}
}

// @Summary Get RecipeStep list
// @Tags RecipeStep
// @Security ApiKeyAuth
// @Description Get RecipeStep list
// @ID get-recipe-step-list
// @Accept json
// @Produce  json
// @Param limit query int false "limit" default(25)
// @Param offset query int false "offset" default(0)
// @Param filter query model.RecipeStepFilter true "RecipeStep filters"
// @Success 200 {object} model.RecipeStepList
// @Failure default {object} response.ErrorResponse
// @Router /v1/recipe-step [get]
func (h *Handler) handlerGetRecipeStepList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &model.RecipeStepFilter{}
		decoder := schema.NewDecoder()
		decoder.Decode(req, r.URL.Query())

		limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
		if err != nil {
			limit = 25
		}
		offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
		if err != nil {
			offset = 0
		}
		RecipeStepList, err := h.service.RecipeStepService.GetList(limit, offset, req)
		if err != nil {
			response.ErrorRespond(w, r, http.StatusBadRequest, err)
			return
		}
		response.Respond(w, r, http.StatusOK, RecipeStepList)
	}
}

// @Summary Update RecipeStep
// @Tags RecipeStep
// @Security ApiKeyAuth
// @Description Update RecipeStep
// @ID update-recipe-step
// @Accept json
// @Produce  json
// @Param input body model.UpdateRecipeStep true "RecipeStep update data"
// @Param id path int true "RecipeStep id"
// @Success 200 {object} model.RecipeStep
// @Failure default {object} response.ErrorResponse
// @Router /v1/recipe-step/{id} [patch]
func (h *Handler) handlerUpdateRecipeStep() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			response.ErrorRespond(w, r, http.StatusOK, err)
			return
		}
		data := &model.UpdateRecipeStep{}
		if err := json.NewDecoder(r.Body).Decode(data); err != nil {
			response.ErrorRespond(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		RecipeStep, err := h.service.RecipeStepService.Update(id, data)
		if err != nil {
			response.ErrorRespond(w, r, http.StatusBadRequest, err)
			return
		}
		response.Respond(w, r, http.StatusOK, RecipeStep)
	}
}

// @Summary Delete RecipeStep
// @Tags RecipeStep
// @Security ApiKeyAuth
// @Description Delete RecipeStep
// @ID delete-recipe-step
// @Accept json
// @Produce  json
// @Param id path int true "RecipeStep id"
// @Success 200 {object} model.RecipeStep
// @Failure default {object} response.ErrorResponse
// @Router /v1/recipe-step/{id} [delete]
func (h *Handler) handlerDeleteRecipeStep() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			response.ErrorRespond(w, r, http.StatusOK, err)
			return
		}
		RecipeStep, err := h.service.RecipeStepService.Delete(id)
		if err != nil {
			response.ErrorRespond(w, r, http.StatusBadRequest, err)
			return
		}
		response.Respond(w, r, http.StatusOK, RecipeStep)
	}
}

// @Summary Upload photo
// @Tags RecipeStep
// @Security ApiKeyAuth
// @Description Upload recipe step photo
// @ID upload-recipe-step-photo
// @Accept			multipart/form-data
// @Produce  json
// @Param id path int true "recipeStep id"
// @Param	photo formData file	 true "this is a test file"
// @Success 200 {object} model.RecipeStep
// @Failure default {object} response.ErrorResponse
// @Router /v1/recipe-step/{id}/photo [post]
func (h *Handler) uploadRecipeStepPhotoHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			response.ErrorRespond(w, r, http.StatusOK, err)
			return
		}
		file, fileHeader, err := r.FormFile("photo")
		if err != nil {
			response.Respond(w, r, http.StatusInternalServerError, err)
		}
		defer file.Close()
		product, err := h.service.RecipeStepService.UploadPhoto(id, file, fileHeader)
		if err != nil {
			response.ErrorRespond(w, r, http.StatusBadRequest, err)
			return
		}
		response.Respond(w, r, http.StatusOK, product)
	}
}

// @Summary Delete photo
// @Tags RecipeStep
// @Security ApiKeyAuth
// @Description Delete recipe step photo
// @ID delete-recipe-step-photo
// @Produce  json
// @Param id path int true "recipe step id"
// @Success 200 {object} model.RecipeStep
// @Router /v1/recipe-step/{id}/photo [delete]
func (h *Handler) deleteRecipeStepPhotoHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			response.ErrorRespond(w, r, http.StatusOK, err)
			return
		}
		product, err := h.service.RecipeStepService.DeletePhoto(id)
		if err != nil {
			response.ErrorRespond(w, r, http.StatusBadRequest, err)
			return
		}
		response.Respond(w, r, http.StatusOK, product)
	}
}
