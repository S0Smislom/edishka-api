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
	product.HandleFunc("/{id:[0-9]+}/photo", h.uploadPhotoHandler()).Methods(http.MethodPost)
	product.HandleFunc("/{id:[0-9]+}/photo", h.deletePhotoHandler()).Methods(http.MethodDelete)

	// Recipe
	recipe := api.PathPrefix("/recipe").Subrouter()
	recipe.HandleFunc("", h.handlerCreateRecipe()).Methods(http.MethodPost)
	recipe.HandleFunc("", h.handlerGetRecipeList()).Methods(http.MethodGet)
	recipe.HandleFunc("/{id:[0-9]+}", h.handlerGetRecipeById()).Methods(http.MethodGet)
	recipe.HandleFunc("/{id:[0-9]+}", h.handlerUpdateRecipe()).Methods(http.MethodPatch)
	recipe.HandleFunc("/{id:[0-9]+}", h.handlerDeleteRecipe()).Methods(http.MethodDelete)

	// RecipeStep
	recipeStep := api.PathPrefix("/recipe-step").Subrouter()
	recipeStep.HandleFunc("", h.handlerCreateRecipeStep()).Methods(http.MethodPost)
	recipeStep.HandleFunc("", h.handlerGetRecipeStepList()).Methods(http.MethodGet)
	recipeStep.HandleFunc("/{id:[0-9]+}", h.handlerGetRecipeStepById()).Methods(http.MethodGet)
	recipeStep.HandleFunc("/{id:[0-9]+}", h.handlerUpdateRecipeStep()).Methods(http.MethodPatch)
	recipeStep.HandleFunc("/{id:[0-9]+}", h.handlerDeleteRecipeStep()).Methods(http.MethodDelete)

	// StepProduct
	stepProduct := api.PathPrefix("/step-product").Subrouter()
	stepProduct.HandleFunc("", h.handlerCreateStepProduct()).Methods(http.MethodPost)
	stepProduct.HandleFunc("", h.handlerGetStepProductList()).Methods(http.MethodGet)
	stepProduct.HandleFunc("/{id:[0-9]+}", h.handlerGetStepProductById()).Methods(http.MethodGet)
	stepProduct.HandleFunc("/{id:[0-9]+}", h.handlerUpdateStepProduct()).Methods(http.MethodPatch)
	stepProduct.HandleFunc("/{id:[0-9]+}", h.handlerDeleteStepProduct()).Methods(http.MethodDelete)
	return router
}
