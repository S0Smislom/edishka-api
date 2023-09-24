package handler

import (
	"fmt"
	"food/internal/api/service"
	"food/pkg/config"
	"food/pkg/middleware"
	"net/http"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger/v2"

	_ "food/docs/api"
)

type Handler struct {
	services *service.Service
	config   *config.Config
}

func NewHandler(config *config.Config, services *service.Service) *Handler {
	return &Handler{config: config, services: services}
}

func (h *Handler) InitRoutes() http.Handler {
	router := mux.NewRouter()
	router.Use(middleware.LogRequest)
	router.PathPrefix("/docs/").Handler(httpSwagger.Handler(
		httpSwagger.URL(fmt.Sprintf("%s/docs/api/doc.json", h.config.BaseAPIURL)), //The url pointing to API definition
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	)).Methods(http.MethodGet)

	// Auth
	router.HandleFunc("/login", h.logIn()).Methods(http.MethodPost)
	router.HandleFunc("/login/confirm", h.confirmCode()).Methods(http.MethodPost)
	// router.HandleFunc("/refresh", h.refreshToken()).Methods(http.MethodGet)

	// V1
	apiRouter := router.PathPrefix("/v1").Subrouter()
	// Profile
	profileRouter := apiRouter.PathPrefix("/profile").Subrouter()
	profileRouter.Use(h.authenticateUser)
	profileRouter.HandleFunc("", h.getCurrentProfileHandler()).Methods(http.MethodGet)
	profileRouter.HandleFunc("", h.updateProfileHandler()).Methods(http.MethodPatch)
	profileRouter.HandleFunc("/photo", h.uploadProfilePhotoHandler()).Methods(http.MethodPost)
	profileRouter.HandleFunc("/photo", h.deleteProfilePhotoHandler()).Methods(http.MethodDelete)
	return router
}
