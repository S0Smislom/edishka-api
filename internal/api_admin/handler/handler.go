package handler

import (
	"fmt"
	"food/internal/api_admin/service"
	"food/pkg/config"
	"food/pkg/middleware"
	"net/http"

	_ "food/docs/api_admin"

	httpSwagger "github.com/swaggo/http-swagger/v2"

	"github.com/gorilla/mux"
)

type Handler struct {
	config  *config.Config
	service *service.Service
}

func NewHandler(config *config.Config, service *service.Service) *Handler {
	return &Handler{config: config, service: service}
}

func (h *Handler) InitRoutes() http.Handler {
	router := mux.NewRouter()
	router.Use(middleware.LogRequest)
	router.PathPrefix("/docs/").Handler(httpSwagger.Handler(
		httpSwagger.URL(fmt.Sprintf("%s/docs/api/doc.json", h.config.BaseAdminAPIURL)), //The url pointing to API definition
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	)).Methods(http.MethodGet)

	// Auth
	router.HandleFunc("/login", h.login()).Methods(http.MethodPost)
	api := router.PathPrefix("/v1").Subrouter()
	api.Use(h.authenticateUser)
	// User
	user := api.PathPrefix("/user").Subrouter()
	user.HandleFunc("/me", h.me()).Methods(http.MethodGet)

	// Product
	product := api.PathPrefix("/product").Subrouter()
	product.HandleFunc("", h.handlerCreateProduct()).Methods(http.MethodPost)
	product.HandleFunc("", h.handlerGetProductList()).Methods(http.MethodGet)
	product.HandleFunc("/{id:[0-9]+}", h.handlerGetProductById()).Methods(http.MethodGet)
	// TODO Подумать про метод put, мб будет проще с ним
	product.HandleFunc("/{id:[0-9]+}", h.handlerUpdateProduct()).Methods(http.MethodPatch)
	product.HandleFunc("/{id:[0-9]+}", h.handlerDeleteProduct()).Methods(http.MethodDelete)
	return router
}
