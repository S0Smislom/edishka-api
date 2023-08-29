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
	return router
}
