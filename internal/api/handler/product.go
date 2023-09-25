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

// @Summary Create product
// @Tags Product
// @Security ApiKeyAuth
// @Description Create product
// @ID create-product
// @Accept  json
// @Produce  json
// @Param input body model.CreateProduct true "product info"
// @Success 200 {object} model.Product
// @Failure default {object} response.ErrorResponse
// @Router /v1/product [post]
func (h *Handler) createProductHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &model.CreateProduct{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			response.ErrorRespond(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		req.CreatedById = r.Context().Value(userCtx).(int)
		product, err := h.services.Product.Create(req)
		if err != nil {
			response.ErrorRespond(w, r, http.StatusBadRequest, err)
			return
		}
		response.Respond(w, r, http.StatusOK, product)
	}
}

// @Summary Get product by id
// @Tags Product
// @Description Get product by id
// @ID get-product
// @Accept json
// @Produce  json
// @Param id path int true "product id"
// @Success 200 {object} model.Product
// @Failure default {object} response.ErrorResponse
// @Router /v1/product/{id} [get]
func (h *Handler) getProductByIdHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			response.ErrorRespond(w, r, http.StatusOK, err)
			return
		}
		product, err := h.services.Product.GetById(id)
		if err != nil {
			response.ErrorRespond(w, r, http.StatusNotFound, err)
			return
		}
		response.Respond(w, r, http.StatusOK, product)
	}
}

// @Summary Get product list
// @Tags Product
// @Description Get product list
// @ID get-product-list
// @Accept json
// @Produce  json
// @Param limit query int false "limit" default(25)
// @Param offset query int false "offset" default(0)
// @Param filter query model.ProductFilter true "product filters"
// @Success 200 {object} model.ProductList
// @Failure default {object} response.ErrorResponse
// @Router /v1/product [get]
func (h *Handler) getProductListHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &model.ProductFilter{}
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
		productList, err := h.services.Product.GetList(limit, offset, req)
		if err != nil {
			response.ErrorRespond(w, r, http.StatusBadRequest, err)
			return
		}
		response.Respond(w, r, http.StatusOK, productList)
	}
}

// @Summary Update product
// @Tags Product
// @Security ApiKeyAuth
// @Description Update product
// @ID update-product
// @Accept json
// @Produce  json
// @Param input body model.UpdateProduct true "Product update data"
// @Param id path int true "Product id"
// @Success 200 {object} model.Product
// @Failure default {object} response.ErrorResponse
// @Router /v1/product/{id} [patch]
func (h *Handler) updateProductHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			response.ErrorRespond(w, r, http.StatusOK, err)
			return
		}
		data := &model.UpdateProduct{}
		if err := json.NewDecoder(r.Body).Decode(data); err != nil {
			response.ErrorRespond(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		userId := r.Context().Value(userCtx).(int)
		data.UpdatedById = &userId
		product, err := h.services.Product.Update(id, userId, data)
		if err != nil {
			response.ErrorRespond(w, r, http.StatusBadRequest, err)
			return
		}
		response.Respond(w, r, http.StatusOK, product)
	}
}

// @Summary Delete product
// @Tags Product
// @Security ApiKeyAuth
// @Description Delete product
// @ID delete-product
// @Accept json
// @Produce  json
// @Param id path int true "Product id"
// @Success 200 {object} model.Product
// @Failure default {object} response.ErrorResponse
// @Router /v1/product/{id} [delete]
func (h *Handler) deleteProductHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			response.ErrorRespond(w, r, http.StatusOK, err)
			return
		}
		currentUserId := r.Context().Value(userCtx).(int)
		product, err := h.services.Product.Delete(id, currentUserId)
		if err != nil {
			response.ErrorRespond(w, r, http.StatusBadRequest, err)
			return
		}
		response.Respond(w, r, http.StatusOK, product)
	}
}

// @Summary Upload photo
// @Tags Product
// @Security ApiKeyAuth
// @Description Upload product photo
// @ID upload-product-photo
// @Accept			multipart/form-data
// @Produce  json
// @Param id path int true "Product id"
// @Param	photo formData file	 true "this is a test file"
// @Success 200 {object} model.Product
// @Failure default {object} response.ErrorResponse
// @Router /v1/product/{id}/photo [post]
func (h *Handler) uploadProductPhotoHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			response.ErrorRespond(w, r, http.StatusOK, err)
			return
		}
		file, fileHeader, err := r.FormFile("photo")
		if err != nil {
			response.Respond(w, r, http.StatusInternalServerError, err)
		}
		defer file.Close()
		currentUserId := r.Context().Value(userCtx).(int)
		product, err := h.services.Product.UploadPhoto(id, currentUserId, file, fileHeader)
		if err != nil {
			response.ErrorRespond(w, r, http.StatusBadRequest, err)
			return
		}
		response.Respond(w, r, http.StatusOK, product)
	}
}

// @Summary Delete photo
// @Tags Product
// @Security ApiKeyAuth
// @Description Delete product photo
// @ID delete-product-photo
// @Produce  json
// @Param id path int true "Product id"
// @Success 200 {object} model.Product
// @Router /v1/product/{id}/photo [delete]
func (h *Handler) deleteProductPhotoHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			response.ErrorRespond(w, r, http.StatusOK, err)
			return
		}
		currentUserId := r.Context().Value(userCtx).(int)
		product, err := h.services.Product.DeletePhoto(id, currentUserId)
		if err != nil {
			response.ErrorRespond(w, r, http.StatusBadRequest, err)
			return
		}
		response.Respond(w, r, http.StatusOK, product)
	}
}
