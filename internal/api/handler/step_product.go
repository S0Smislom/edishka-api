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

// @Summary Create Step Product
// @Tags StepProduct
// @Security ApiKeyAuth
// @Description Create Step Product
// @ID create-step-product
// @Accept  json
// @Produce  json
// @Param input body model.CreateStepProduct true "StepProduct info"
// @Success 200 {object} model.StepProduct
// @Failure default {object} response.ErrorResponse
// @Router /v1/step-product [post]
func (h *Handler) createStepProductHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &model.CreateStepProduct{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		StepProduct, err := h.services.StepProduct.Create(r.Context().Value(userCtx).(int), req)
		if err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		response.Respond(w, r, http.StatusOK, StepProduct)
	}
}

// @Summary Get Step Product by id
// @Tags StepProduct
// @Description Get Step Product by id
// @ID get-step-product
// @Accept json
// @Produce  json
// @Param id path int true "StepProduct id"
// @Success 200 {object} model.StepProduct
// @Failure default {object} response.ErrorResponse
// @Router /v1/step-product/{id} [get]
func (h *Handler) getStepProductByIdHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		StepProduct, err := h.services.StepProduct.GetById(id)
		if err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		response.Respond(w, r, http.StatusOK, StepProduct)
	}
}

// @Summary Get Step Product list
// @Tags StepProduct
// @Description Get Step Product list
// @ID get-step-product-list
// @Accept json
// @Produce  json
// @Param limit query int false "limit" default(25)
// @Param offset query int false "offset" default(0)
// @Param filter query model.StepProductFilter true "StepProduct filters"
// @Success 200 {object} model.StepProductList
// @Failure default {object} response.ErrorResponse
// @Router /v1/step-product [get]
func (h *Handler) getStepProductListHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &model.StepProductFilter{}
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
		StepProductList, err := h.services.StepProduct.GetList(limit, offset, req)
		if err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		response.Respond(w, r, http.StatusOK, StepProductList)
	}
}

// @Summary Update Step Product
// @Tags StepProduct
// @Security ApiKeyAuth
// @Description Update StepProduct
// @ID update-step-product
// @Accept json
// @Produce  json
// @Param input body model.UpdateStepProduct true "StepProduct update data"
// @Param id path int true "StepProduct id"
// @Success 200 {object} model.StepProduct
// @Failure default {object} response.ErrorResponse
// @Router /v1/step-product/{id} [patch]
func (h *Handler) updateStepProductHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		data := &model.UpdateStepProduct{}
		if err := json.NewDecoder(r.Body).Decode(data); err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		userId := r.Context().Value(userCtx).(int)
		data.UpdatedById = &userId
		StepProduct, err := h.services.StepProduct.Update(id, userId, data)
		if err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		response.Respond(w, r, http.StatusOK, StepProduct)
	}
}

// @Summary Delete Step Product
// @Tags StepProduct
// @Security ApiKeyAuth
// @Description Delete Step Product
// @ID delete-step-product
// @Accept json
// @Produce  json
// @Param id path int true "StepProduct id"
// @Success 200 {object} model.StepProduct
// @Failure default {object} response.ErrorResponse
// @Router /v1/step-product/{id} [delete]
func (h *Handler) deleteStepProductHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		currentUserId := r.Context().Value(userCtx).(int)
		StepProduct, err := h.services.StepProduct.Delete(id, currentUserId)
		if err != nil {
			response.ErrorRespond(w, r, err)
			return
		}
		response.Respond(w, r, http.StatusOK, StepProduct)
	}
}
