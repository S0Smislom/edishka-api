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

// @Summary Create Recipe
// @Tags Recipe
// @Security ApiKeyAuth
// @Description Create Recipe
// @ID create-recipe
// @Accept  json
// @Produce  json
// @Param input body model.CreateRecipe true "Recipe info"
// @Success 200 {object} model.Recipe
// @Failure default {object} response.ErrorResponse
// @Router /v1/recipe [post]
func (h *Handler) createRecipeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &model.CreateRecipe{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		req.CreatedById = r.Context().Value(userCtx).(int)
		Recipe, err := h.services.Recipe.Create(req)
		if err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		response.Respond(w, r, http.StatusOK, Recipe)
	}
}

// @Summary Get Recipe by id
// @Tags Recipe
// @Description Get Recipe by id
// @ID get-recipe
// @Accept json
// @Produce  json
// @Param id path int true "Recipe id"
// @Success 200 {object} model.Recipe
// @Failure default {object} response.ErrorResponse
// @Router /v1/recipe/{id} [get]
func (h *Handler) getRecipeByIdHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		Recipe, err := h.services.Recipe.GetById(id)
		if err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		response.Respond(w, r, http.StatusOK, Recipe)
	}
}

// @Summary Get Recipe by id private
// @Tags Recipe
// @Security ApiKeyAuth
// @Description Get Recipe by id private
// @ID get-recipe-private
// @Accept json
// @Produce  json
// @Param id path int true "Recipe id"
// @Success 200 {object} model.Recipe
// @Failure default {object} response.ErrorResponse
// @Router /v1/recipe/{id}/private [get]
func (h *Handler) getRecipeByIdPrivateHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		Recipe, err := h.services.Recipe.GetByIdPrivate(id, r.Context().Value(userCtx).(int))
		if err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		response.Respond(w, r, http.StatusOK, Recipe)
	}
}

// @Summary Get Recipe list
// @Tags Recipe
// @Description Get Recipe list
// @ID get-recipe-list
// @Accept json
// @Produce  json
// @Param limit query int false "limit" default(25)
// @Param offset query int false "offset" default(0)
// @Param filter query model.RecipeFilter true "Recipe filters"
// @Success 200 {object} model.RecipeList
// @Failure default {object} response.ErrorResponse
// @Router /v1/recipe [get]
func (h *Handler) getRecipeListHandler() http.HandlerFunc {
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
		RecipeList, err := h.services.Recipe.GetList(limit, offset, req)
		if err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		response.Respond(w, r, http.StatusOK, RecipeList)
	}
}

// @Summary Get Recipe list private
// @Tags Recipe
// @Security ApiKeyAuth
// @Description Get Recipe list private
// @ID get-recipe-list-private
// @Accept json
// @Produce  json
// @Param limit query int false "limit" default(25)
// @Param offset query int false "offset" default(0)
// @Param filter query model.RecipeFilter true "Recipe filters"
// @Success 200 {object} model.RecipeList
// @Failure default {object} response.ErrorResponse
// @Router /v1/recipe/private [get]
func (h *Handler) getRecipeListPrivateHandler() http.HandlerFunc {
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
		currentUserId := r.Context().Value(userCtx).(int)
		RecipeList, err := h.services.Recipe.GetListPrivate(limit, offset, currentUserId, req)
		if err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		response.Respond(w, r, http.StatusOK, RecipeList)
	}
}

// @Summary Update Recipe
// @Tags Recipe
// @Security ApiKeyAuth
// @Description Update Recipe
// @ID update-recipe
// @Accept json
// @Produce  json
// @Param input body model.UpdateRecipe true "Recipe update data"
// @Param id path int true "Recipe id"
// @Success 200 {object} model.Recipe
// @Failure default {object} response.ErrorResponse
// @Router /v1/recipe/{id} [patch]
func (h *Handler) updateRecipeHandler() http.HandlerFunc {
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
		Recipe, err := h.services.Recipe.Update(id, userId, data)
		if err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		response.Respond(w, r, http.StatusOK, Recipe)
	}
}

// @Summary Delete Recipe
// @Tags Recipe
// @Security ApiKeyAuth
// @Description Delete Recipe
// @ID delete-recipe
// @Accept json
// @Produce  json
// @Param id path int true "Recipe id"
// @Success 200 {object} model.Recipe
// @Failure default {object} response.ErrorResponse
// @Router /v1/recipe/{id} [delete]
func (h *Handler) deleteRecipeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		currentUserId := r.Context().Value(userCtx).(int)
		Recipe, err := h.services.Recipe.Delete(id, currentUserId)
		if err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		response.Respond(w, r, http.StatusOK, Recipe)
	}
}
