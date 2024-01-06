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

// @Summary Create RecipeGallery
// @Tags RecipeGallery
// @Security ApiKeyAuth
// @Description Create RecipeGallery
// @ID create-recipe-step
// @Accept  json
// @Produce  json
// @Param input formData model.CreateRecipeGallery true "RecipeGallery info"
// @Param photo formData file true "this is a test file"
// @Success 200 {object} model.RecipeGallery
// @Failure default {object} response.ErrorResponse
// @Router /v1/recipe-gallery [post]
func (h *Handler) createRecipeGalleryPhotoHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := &model.CreateRecipeGallery{}
		r.ParseMultipartForm(defaultMaxMemory)
		decoder := schema.NewDecoder()
		decoder.Decode(data, r.PostForm)

		file, fileHeader, err := r.FormFile("photo")
		if err != nil {
			response.Respond(w, r, http.StatusInternalServerError, err)
			return
		}
		defer file.Close()
		data.CreatedById = r.Context().Value(userCtx).(int)
		recipeGalleryPhoto, err := h.service.RecipeGalleryService.Create(
			data,
			file,
			fileHeader,
		)
		if err != nil {
			response.Respond(w, r, http.StatusInternalServerError, err)
			return
		}
		response.Respond(w, r, http.StatusOK, recipeGalleryPhoto)
	}
}

// @Summary Update RecipeGallery
// @Tags RecipeGallery
// @Security ApiKeyAuth
// @Description Update RecipeGallery
// @ID update-recipe-step
// @Accept  json
// @Produce  json
// @Param input body model.UpdateRecipeGallery true "Recipe gallery update data"
// @Param id path int true "Recipe id"
// @Success 200 {object} model.RecipeGallery
// @Failure default {object} response.ErrorResponse
// @Router /v1/recipe-gallery/{id} [patch]
func (h *Handler) updateRecipeGalleryPhotoHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		data := &model.UpdateRecipeGallery{}
		if err := json.NewDecoder(r.Body).Decode(data); err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		userId := r.Context().Value(userCtx).(int)
		data.UpdatedById = &userId
		recipeGalleryPhoto, err := h.service.RecipeGalleryService.Update(id, data)
		if err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		response.Respond(w, r, http.StatusOK, recipeGalleryPhoto)
	}
}

// @Summary Delete recipe gallery
// @Tags RecipeGallery
// @Security ApiKeyAuth
// @Description Delete recipe gallery
// @ID delete-recipe-gallery
// @Accept json
// @Produce  json
// @Param id path int true "Recipe gallery id"
// @Success 200 {object} model.RecipeGallery
// @Failure default {object} response.ErrorResponse
// @Router /v1/recipe-gallery/{id} [delete]
func (h *Handler) deleteRecipeGalleryPhotoHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		recipeGalleryPhoto, err := h.service.RecipeGalleryService.Delete(id)
		if err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		response.Respond(w, r, http.StatusOK, recipeGalleryPhoto)
	}
}
