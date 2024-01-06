package handler

import (
	"encoding/json"
	"food/internal/api/model"
	"food/pkg/response"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

// @Summary Create Recipe Step
// @Tags RecipeStep
// @Security ApiKeyAuth
// @Description Create Recipe Step
// @ID create-recipe-step
// @Accept  json
// @Produce  json
// @Param input body model.CreateRecipeStep true "RecipeStep info"
// @Success 200 {object} model.RecipeStep
// @Failure default {object} response.ErrorResponse
// @Router /v1/recipe-step [post]
func (h *Handler) createRecipeStepHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &model.CreateRecipeStep{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		recipeStep, err := h.services.RecipeStep.Create(r.Context().Value(userCtx).(int), req)
		if err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		response.Respond(w, r, http.StatusOK, recipeStep)
	}
}

// @Summary Get RecipeStep by id
// @Tags RecipeStep
// @Description Get RecipeStep by id
// @ID get-recipe-step
// @Accept json
// @Produce  json
// @Param id path int true "RecipeStep id"
// @Success 200 {object} model.RecipeStep
// @Failure default {object} response.ErrorResponse
// @Router /v1/recipe-step/{id} [get]
func (h *Handler) getRecipeStepByIdHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		recipeStep, err := h.services.RecipeStep.GetById(id)
		if err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		response.Respond(w, r, http.StatusOK, recipeStep)
	}
}

// @Summary Get Recipe Step list
// @Tags RecipeStep
// @Description Get RecipeStep list
// @ID get-recipe-step-list
// @Accept json
// @Produce  json
// @Param limit query int false "limit" default(25)
// @Param offset query int false "offset" default(0)
// @Param filter query model.RecipeStepFilter true "Recipe step filters"
// @Success 200 {object} model.RecipeStepList
// @Failure default {object} response.ErrorResponse
// @Router /v1/recipe-step [get]
func (h *Handler) getRecipeStepListHandler() http.HandlerFunc {
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
		recipeStepList, err := h.services.RecipeStep.GetList(limit, offset, req)
		if err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		response.Respond(w, r, http.StatusOK, recipeStepList)
	}
}

// @Summary Update Recipe Step
// @Tags RecipeStep
// @Security ApiKeyAuth
// @Description Update Recipe Step
// @ID update-recipe-step
// @Accept json
// @Produce  json
// @Param input body model.UpdateRecipeStep true "Recipe step update data"
// @Param id path int true "Recipe Step id"
// @Success 200 {object} model.RecipeStep
// @Failure default {object} response.ErrorResponse
// @Router /v1/recipe-step/{id} [patch]
func (h *Handler) updateRecipeStepHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		data := &model.UpdateRecipeStep{}
		if err := json.NewDecoder(r.Body).Decode(data); err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		userId := r.Context().Value(userCtx).(int)
		data.UpdatedById = &userId
		recipeStep, err := h.services.RecipeStep.Update(id, userId, data)
		if err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		response.Respond(w, r, http.StatusOK, recipeStep)
	}
}

// @Summary Delete Recipe Step
// @Tags RecipeStep
// @Security ApiKeyAuth
// @Description Delete Recipe Step
// @ID delete-recipe-step
// @Accept json
// @Produce  json
// @Param id path int true "Recipe Step id"
// @Success 200 {object} model.RecipeStep
// @Failure default {object} response.ErrorResponse
// @Router /v1/recipe-step/{id} [delete]
func (h *Handler) deleteRecipeStepHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		currentUserId := r.Context().Value(userCtx).(int)
		recipeStep, err := h.services.RecipeStep.Delete(id, currentUserId)
		if err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		response.Respond(w, r, http.StatusOK, recipeStep)
	}
}

// @Summary Upload photo
// @Tags RecipeStep
// @Security ApiKeyAuth
// @Description Upload RecipeStep photo
// @ID upload-recipe-step-photo
// @Accept			multipart/form-data
// @Produce  json
// @Param id path int true "RecipeStep id"
// @Param	photo formData file	 true "this is a test file"
// @Success 200 {object} model.RecipeStep
// @Failure default {object} response.ErrorResponse
// @Router /v1/recipe-step/{id}/photo [post]
func (h *Handler) uploadRecipeStepPhotoHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		file, fileHeader, err := r.FormFile("photo")
		if err != nil {
			response.Respond(w, r, http.StatusInternalServerError, err)
		}
		defer file.Close()
		currentUserId := r.Context().Value(userCtx).(int)
		RecipeStep, err := h.services.RecipeStep.UploadPhoto(id, currentUserId, file, fileHeader)
		if err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		response.Respond(w, r, http.StatusOK, RecipeStep)
	}
}

// @Summary Delete photo
// @Tags RecipeStep
// @Security ApiKeyAuth
// @Description Delete RecipeStep photo
// @ID delete-recipe-step-photo
// @Produce  json
// @Param id path int true "RecipeStep id"
// @Success 200 {object} model.RecipeStep
// @Router /v1/recipe-step/{id}/photo [delete]
func (h *Handler) deleteRecipeStepPhotoHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		currentUserId := r.Context().Value(userCtx).(int)
		RecipeStep, err := h.services.RecipeStep.DeletePhoto(id, currentUserId)
		if err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		response.Respond(w, r, http.StatusOK, RecipeStep)
	}
}
