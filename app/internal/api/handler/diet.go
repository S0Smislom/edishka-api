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
// @Tags Diet
// @Security ApiKeyAuth
// @Description Create diet
// @ID create-diet
// @Accept  json
// @Produce  json
// @Param input body model.CreateDiet true "diet info"
// @Success 200 {object} model.Diet
// @Failure default {object} response.ErrorResponse
// @Router /v1/diet [post]
func (h *Handler) createDietHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &model.CreateDiet{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		currentUserId := r.Context().Value(userCtx).(int)
		result, err := h.services.Diet.Create(currentUserId, req)
		if err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		response.Respond(w, r, http.StatusOK, result)
	}
}

// @Summary Get diet by id
// @Tags Diet
// @Security ApiKeyAuth
// @Description Get diet by id
// @ID get-diet
// @Accept json
// @Produce  json
// @Param id path int true "Diet id"
// @Success 200 {object} model.Diet
// @Failure default {object} response.ErrorResponse
// @Router /v1/diet/{id}/private [get]
func (h *Handler) getDietByIdPrivateHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		currentUserId := r.Context().Value(userCtx).(int)
		result, err := h.services.Diet.GetByIdPrivate(currentUserId, id)
		if err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		response.Respond(w, r, http.StatusOK, result)
	}
}

// @Summary Get diet list
// @Tags Diet
// @Security ApiKeyAuth
// @Description Get diet list
// @ID get-diet-list
// @Accept json
// @Produce  json
// @Param limit query int false "limit" default(25)
// @Param offset query int false "offset" default(0)
// @Param filter query model.DietFilter true "Diet item filters"
// @Success 200 {object} model.DietList
// @Failure default {object} response.ErrorResponse
// @Router /v1/diet/private [get]
func (h *Handler) getDietListPrivateHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &model.DietFilter{}
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
		result, err := h.services.Diet.GetListPrivate(userId, limit, offset, req)
		if err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		response.Respond(w, r, http.StatusOK, result)
	}
}

// @Summary Update diet
// @Tags Diet
// @Security ApiKeyAuth
// @Description Update diet
// @ID update-diet
// @Accept json
// @Produce  json
// @Param input body model.UpdateDiet true "Diet update data"
// @Param id path int true "Diet item id"
// @Success 200 {object} model.Diet
// @Failure default {object} response.ErrorResponse
// @Router /v1/diet/{id} [patch]
func (h *Handler) updateDietHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		data := &model.UpdateDiet{}
		if err := json.NewDecoder(r.Body).Decode(data); err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		userId := r.Context().Value(userCtx).(int)
		result, err := h.services.Diet.Update(userId, id, data)
		if err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		response.Respond(w, r, http.StatusOK, result)
	}
}

// @Summary Delete diet
// @Tags Diet
// @Security ApiKeyAuth
// @Description Delete diet item
// @ID delete-diet
// @Accept json
// @Produce  json
// @Param id path int true "Diet item id"
// @Success 200 {object} model.Diet
// @Failure default {object} response.ErrorResponse
// @Router /v1/diet/{id} [delete]
func (h *Handler) deleteDietHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		currentUserId := r.Context().Value(userCtx).(int)
		product, err := h.services.Diet.Delete(currentUserId, id)
		if err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		response.Respond(w, r, http.StatusOK, product)
	}
}
