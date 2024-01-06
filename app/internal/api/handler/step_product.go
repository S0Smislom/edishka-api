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

// @Summary Create Recipe Product
// @Tags RecipeProduct
// @Security ApiKeyAuth
// @Description Create Recipe Product
// @ID create-recipe-product
// @Accept  json
// @Produce  json
// @Param input body model.CreateRecipeProduct true "RecipeProduct info"
// @Success 200 {object} model.RecipeProduct
// @Failure default {object} response.ErrorResponse
// @Router /v1/recipe-product [post]
func (h *Handler) createRecipeProductHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &model.CreateRecipeProduct{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		recipeProduct, err := h.services.RecipeProduct.Create(r.Context().Value(userCtx).(int), req)
		if err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		response.Respond(w, r, http.StatusOK, recipeProduct)
	}
}

// @Summary Get Recipe Product by id
// @Tags RecipeProduct
// @Description Get Recipe Product by id
// @ID get-recipe-product
// @Accept json
// @Produce  json
// @Param id path int true "RecipeProduct id"
// @Success 200 {object} model.RecipeProduct
// @Failure default {object} response.ErrorResponse
// @Router /v1/recipe-product/{id} [get]
func (h *Handler) getRecipeProductByIdHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		recipeProduct, err := h.services.RecipeProduct.GetById(id)
		if err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		response.Respond(w, r, http.StatusOK, recipeProduct)
	}
}

// @Summary Get Recipe Product list
// @Tags RecipeProduct
// @Description Get Recipe Product list
// @ID get-recipe-product-list
// @Accept json
// @Produce  json
// @Param limit query int false "limit" default(25)
// @Param offset query int false "offset" default(0)
// @Param filter query model.RecipeProductFilter true "RecipeProduct filters"
// @Success 200 {object} model.RecipeProductList
// @Failure default {object} response.ErrorResponse
// @Router /v1/recipe-product [get]
func (h *Handler) getRecipeProductListHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &model.RecipeProductFilter{}
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
		recipeProductList, err := h.services.RecipeProduct.GetList(limit, offset, req)
		if err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		response.Respond(w, r, http.StatusOK, recipeProductList)
	}
}

// @Summary Update Recipe Product
// @Tags RecipeProduct
// @Security ApiKeyAuth
// @Description Update RecipeProduct
// @ID update-recipe-product
// @Accept json
// @Produce  json
// @Param input body model.UpdateRecipeProduct true "RecipeProduct update data"
// @Param id path int true "RecipeProduct id"
// @Success 200 {object} model.RecipeProduct
// @Failure default {object} response.ErrorResponse
// @Router /v1/recipe-product/{id} [patch]
func (h *Handler) updateRecipeProductHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		data := &model.UpdateRecipeProduct{}
		if err := json.NewDecoder(r.Body).Decode(data); err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		userId := r.Context().Value(userCtx).(int)
		data.UpdatedById = &userId
		recipeProduct, err := h.services.RecipeProduct.Update(id, userId, data)
		if err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		response.Respond(w, r, http.StatusOK, recipeProduct)
	}
}

// @Summary Delete Recipe Product
// @Tags RecipeProduct
// @Security ApiKeyAuth
// @Description Delete Recipe Product
// @ID delete-recipe-product
// @Accept json
// @Produce  json
// @Param id path int true "RecipeProduct id"
// @Success 200 {object} model.RecipeProduct
// @Failure default {object} response.ErrorResponse
// @Router /v1/recipe-product/{id} [delete]
func (h *Handler) deleteRecipeProductHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		currentUserId := r.Context().Value(userCtx).(int)
		recipeProduct, err := h.services.RecipeProduct.Delete(id, currentUserId)
		if err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		response.Respond(w, r, http.StatusOK, recipeProduct)
	}
}
