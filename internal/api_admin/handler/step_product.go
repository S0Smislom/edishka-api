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

// @Summary Create StepProduct
// @Tags StepProduct
// @Security ApiKeyAuth
// @Description Create StepProduct
// @ID create-step-product
// @Accept  json
// @Produce  json
// @Param input body model.CreateStepProduct true "StepProduct info"
// @Success 200 {object} model.StepProduct
// @Failure default {object} response.ErrorResponse
// @Router /v1/step-product [post]
func (h *Handler) handlerCreateStepProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &model.CreateStepProduct{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			response.ErrorRespond(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		req.CreatedById = r.Context().Value(userCtx).(int)
		StepProduct, err := h.service.StepProductService.Create(req)
		if err != nil {
			response.ErrorRespond(w, r, http.StatusBadRequest, err)
			return
		}
		response.Respond(w, r, http.StatusOK, StepProduct)
	}
}

// @Summary Get StepProduct by id
// @Tags StepProduct
// @Security ApiKeyAuth
// @Description Get StepProduct by id
// @ID get-step-product
// @Accept json
// @Produce  json
// @Param id path int true "StepProduct id"
// @Success 200 {object} model.StepProduct
// @Failure default {object} response.ErrorResponse
// @Router /v1/step-product/{id} [get]
func (h *Handler) handlerGetStepProductById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			response.ErrorRespond(w, r, http.StatusOK, err)
			return
		}
		StepProduct, err := h.service.StepProductService.GetById(id)
		if err != nil {
			response.ErrorRespond(w, r, http.StatusNotFound, err)
			return
		}
		response.Respond(w, r, http.StatusOK, StepProduct)
	}
}

// @Summary Get StepProduct list
// @Tags StepProduct
// @Security ApiKeyAuth
// @Description Get StepProduct list
// @ID get-step-product-list
// @Accept json
// @Produce  json
// @Param limit query int false "limit" default(25)
// @Param offset query int false "offset" default(0)
// @Param filter query model.StepProductFilter true "StepProduct filters"
// @Success 200 {object} model.StepProductList
// @Failure default {object} response.ErrorResponse
// @Router /v1/step-product [get]
func (h *Handler) handlerGetStepProductList() http.HandlerFunc {
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
		StepProductList, err := h.service.StepProductService.GetList(limit, offset, req)
		if err != nil {
			response.ErrorRespond(w, r, http.StatusBadRequest, err)
			return
		}
		response.Respond(w, r, http.StatusOK, StepProductList)
	}
}

// @Summary Update StepProduct
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
func (h *Handler) handlerUpdateStepProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			response.ErrorRespond(w, r, http.StatusOK, err)
			return
		}
		data := &model.UpdateStepProduct{}
		if err := json.NewDecoder(r.Body).Decode(data); err != nil {
			response.ErrorRespond(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		userId := r.Context().Value(userCtx).(int)
		data.UpdatedById = &userId
		StepProduct, err := h.service.StepProductService.Update(id, data)
		if err != nil {
			response.ErrorRespond(w, r, http.StatusBadRequest, err)
			return
		}
		response.Respond(w, r, http.StatusOK, StepProduct)
	}
}

// @Summary Delete StepProduct
// @Tags StepProduct
// @Security ApiKeyAuth
// @Description Delete StepProduct
// @ID delete-step-product
// @Accept json
// @Produce  json
// @Param id path int true "StepProduct id"
// @Success 200 {object} model.StepProduct
// @Failure default {object} response.ErrorResponse
// @Router /v1/step-product/{id} [delete]
func (h *Handler) handlerDeleteStepProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			response.ErrorRespond(w, r, http.StatusOK, err)
			return
		}
		StepProduct, err := h.service.StepProductService.Delete(id)
		if err != nil {
			response.ErrorRespond(w, r, http.StatusBadRequest, err)
			return
		}
		response.Respond(w, r, http.StatusOK, StepProduct)
	}
}
