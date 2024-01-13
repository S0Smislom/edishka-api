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

// @Summary Create shopping item
// @Tags Shopping item
// @Security ApiKeyAuth
// @Description Create Shoppint item
// @ID create-shopping-item
// @Accept  json
// @Produce  json
// @Param input body model.CreateShoppingItem true "Shopping item info"
// @Success 200 {object} model.ShoppingItem
// @Failure default {object} response.ErrorResponse
// @Router /v1/shopping-item [post]
func (h *Handler) createShoppingItemHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &model.CreateShoppingItem{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		req.UserId = r.Context().Value(userCtx).(int)
		product, err := h.services.ShoppingList.Create(req)
		if err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		response.Respond(w, r, http.StatusOK, product)
	}
}

// // @Summary Get shopping item by id
// // @Tags Shopping item
// // @Security ApiKeyAuth
// // @Description Get shopping item by id
// // @ID get-shopping-item
// // @Accept json
// // @Produce  json
// // @Param id path int true "Shopping item id"
// // @Success 200 {object} model.ShoppingItem
// // @Failure default {object} response.ErrorResponse
// // @Router /v1/shopping-item/{id} [get]
// func (h *Handler) getShoppingItemByIdHandler() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		vars := mux.Vars(r)
// 		id, err := strconv.Atoi(vars["id"])
// 		if err != nil {
// 			response.ErrorRespond(w, r, err)
// 			return
// 		}
// 		product, err := h.services.ShoppingList.GetById(id)
// 		if err != nil {
// 			response.ErrorRespond(w, r, err)
// 			return
// 		}
// 		response.Respond(w, r, http.StatusOK, product)
// 	}
// }

// @Summary Get shopping list
// @Tags Shopping item
// @Security ApiKeyAuth
// @Description Get shopping list
// @ID get-shopping-list
// @Accept json
// @Produce  json
// @Param limit query int false "limit" default(25)
// @Param offset query int false "offset" default(0)
// @Param filter query model.ShoppingItemFilter true "Shopping item filters"
// @Success 200 {object} model.ShoppingList
// @Failure default {object} response.ErrorResponse
// @Router /v1/shopping-item [get]
func (h *Handler) getShoppingListHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &model.ShoppingItemFilter{}
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
		userId := r.Context().Value(userCtx).(int)
		shoppingList, err := h.services.ShoppingList.GetList(userId, limit, offset, req)
		if err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		response.Respond(w, r, http.StatusOK, shoppingList)
	}
}

// @Summary Update shopping item
// @Tags Shopping item
// @Security ApiKeyAuth
// @Description Update shopping item
// @ID update-shopping-item
// @Accept json
// @Produce  json
// @Param input body model.UpdateShoppingItem true "Shopping item update data"
// @Param id path int true "Shopping item id"
// @Success 200 {object} model.ShoppingItem
// @Failure default {object} response.ErrorResponse
// @Router /v1/shopping-item/{id} [patch]
func (h *Handler) updateShoppingItemHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		data := &model.UpdateShoppingItem{}
		if err := json.NewDecoder(r.Body).Decode(data); err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		userId := r.Context().Value(userCtx).(int)
		product, err := h.services.ShoppingList.Update(userId, id, data)
		if err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		response.Respond(w, r, http.StatusOK, product)
	}
}

// @Summary Delete shopping item
// @Tags Shopping item
// @Security ApiKeyAuth
// @Description Delete shopping item
// @ID delete-shopping-item
// @Accept json
// @Produce  json
// @Param id path int true "Shopping item id"
// @Success 200 {object} model.ShoppingItem
// @Failure default {object} response.ErrorResponse
// @Router /v1/shopping-item/{id} [delete]
func (h *Handler) deleteShoppingItemHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		currentUserId := r.Context().Value(userCtx).(int)
		product, err := h.services.ShoppingList.Delete(currentUserId, id)
		if err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		response.Respond(w, r, http.StatusOK, product)
	}
}
