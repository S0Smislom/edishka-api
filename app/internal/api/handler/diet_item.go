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

// @Summary Create diet item
// @Tags DietItem
// @Security ApiKeyAuth
// @Description Create diet item
// @ID create-diet-item
// @Accept  json
// @Produce  json
// @Param input body model.CreateDietItem true "diet info"
// @Success 200 {object} model.DietItem
// @Failure default {object} response.ErrorResponse
// @Router /v1/diet-item [post]
func (h *Handler) createDietItemHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &model.CreateDietItem{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		currentUserId := r.Context().Value(userCtx).(int)
		result, err := h.services.DietItem.Create(currentUserId, req)
		if err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		response.Respond(w, r, http.StatusOK, result)
	}
}

// @Summary Get diet item by id
// @Tags DietItem
// @Security ApiKeyAuth
// @Description Get diet item by id
// @ID get-diet-item
// @Accept json
// @Produce  json
// @Param id path int true "Diet item id"
// @Success 200 {object} model.DietItem
// @Failure default {object} response.ErrorResponse
// @Router /v1/diet-item/{id}/private [get]
func (h *Handler) getDietItemByIdPrivateHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		currentUserId := r.Context().Value(userCtx).(int)
		result, err := h.services.DietItem.GetByIdPrivate(currentUserId, id)
		if err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		response.Respond(w, r, http.StatusOK, result)
	}
}

// @Summary Get diet item list
// @Tags DietItem
// @Security ApiKeyAuth
// @Description Get diet item list
// @ID get-diet-item-list
// @Accept json
// @Produce  json
// @Param limit query int false "limit" default(25)
// @Param offset query int false "offset" default(0)
// @Param filter query model.DietItemFilter true "Diet item item filters"
// @Success 200 {object} model.DietItemList
// @Failure default {object} response.ErrorResponse
// @Router /v1/diet-item/private [get]
func (h *Handler) getDietItemListPrivateHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &model.DietItemFilter{}
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
		result, err := h.services.DietItem.GetListPrivate(userId, limit, offset, req)
		if err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		response.Respond(w, r, http.StatusOK, result)
	}
}

// @Summary Update diet item
// @Tags DietItem
// @Security ApiKeyAuth
// @Description Update diet item
// @ID update-diet-item
// @Accept json
// @Produce  json
// @Param input body model.UpdateDietItem true "Diet item update data"
// @Param id path int true "Diet item id"
// @Success 200 {object} model.DietItem
// @Failure default {object} response.ErrorResponse
// @Router /v1/diet-item/{id} [patch]
func (h *Handler) updateDietItemHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		data := &model.UpdateDietItem{}
		if err := json.NewDecoder(r.Body).Decode(data); err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		userId := r.Context().Value(userCtx).(int)
		result, err := h.services.DietItem.Update(userId, id, data)
		if err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		response.Respond(w, r, http.StatusOK, result)
	}
}

// @Summary Delete diet item
// @Tags DietItem
// @Security ApiKeyAuth
// @Description Delete diet item
// @ID delete-diet-item
// @Accept json
// @Produce  json
// @Param id path int true "Diet item item id"
// @Success 200 {object} model.DietItem
// @Failure default {object} response.ErrorResponse
// @Router /v1/diet-item/{id} [delete]
func (h *Handler) deleteDietItemHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		currentUserId := r.Context().Value(userCtx).(int)
		product, err := h.services.DietItem.Delete(currentUserId, id)
		if err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		response.Respond(w, r, http.StatusOK, product)
	}
}
