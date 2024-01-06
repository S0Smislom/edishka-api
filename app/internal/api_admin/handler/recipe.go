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

// @Summary Create recipe
// @Tags Recipe
// @Security ApiKeyAuth
// @Description Create recipe
// @ID create-recipe
// @Accept  json
// @Produce  json
// @Param input body model.CreateRecipe true "Recipe info"
// @Success 200 {object} model.Recipe
// @Failure default {object} response.ErrorResponse
// @Router /v1/recipe [post]
func (h *Handler) handlerCreateRecipe() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &model.CreateRecipe{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		req.CreatedById = r.Context().Value(userCtx).(int)
		recipe, err := h.service.RecipeService.Create(req)
		if err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		response.Respond(w, r, http.StatusOK, recipe)
	}
}

// @Summary Get recipe by id
// @Tags Recipe
// @Security ApiKeyAuth
// @Description Get recipe by id
// @ID get-recipe
// @Accept json
// @Produce  json
// @Param id path int true "Recipe id"
// @Success 200 {object} model.Recipe
// @Failure default {object} response.ErrorResponse
// @Router /v1/recipe/{id} [get]
func (h *Handler) handlerGetRecipeById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		recipe, err := h.service.RecipeService.GetById(id)
		if err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		response.Respond(w, r, http.StatusOK, recipe)
	}
}

// @Summary Get recipe list
// @Tags Recipe
// @Security ApiKeyAuth
// @Description Get recipe list
// @ID get-recipe-list
// @Accept json
// @Produce  json
// @Param limit query int false "limit" default(25)
// @Param offset query int false "offset" default(0)
// @Param filter query model.RecipeFilter true "Recipe filters"
// @Success 200 {object} model.RecipeList
// @Failure default {object} response.ErrorResponse
// @Router /v1/recipe [get]
func (h *Handler) handlerGetRecipeList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &model.RecipeFilter{}
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
		recipeList, err := h.service.RecipeService.GetList(limit, offset, req)
		if err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		response.Respond(w, r, http.StatusOK, recipeList)
	}
}

// @Summary Update recipe
// @Tags Recipe
// @Security ApiKeyAuth
// @Description Update recipe
// @ID update-recipe
// @Accept json
// @Produce  json
// @Param input body model.UpdateRecipe true "Recipe update data"
// @Param id path int true "Recipe id"
// @Success 200 {object} model.Recipe
// @Failure default {object} response.ErrorResponse
// @Router /v1/recipe/{id} [patch]
func (h *Handler) handlerUpdateRecipe() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		data := &model.UpdateRecipe{}
		if err := json.NewDecoder(r.Body).Decode(data); err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		userId := r.Context().Value(userCtx).(int)
		data.UpdatedById = &userId
		recipe, err := h.service.RecipeService.Update(id, data)
		if err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		response.Respond(w, r, http.StatusOK, recipe)
	}
}

// @Summary Delete recipe
// @Tags Recipe
// @Security ApiKeyAuth
// @Description Delete recipe
// @ID delete-recipe
// @Accept json
// @Produce  json
// @Param id path int true "Recipe id"
// @Success 200 {object} model.Recipe
// @Failure default {object} response.ErrorResponse
// @Router /v1/recipe/{id} [delete]
func (h *Handler) handlerDeleteRecipe() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		recipe, err := h.service.RecipeService.Delete(id)
		if err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		response.Respond(w, r, http.StatusOK, recipe)
	}
}
