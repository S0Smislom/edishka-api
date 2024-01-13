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

const (
	defaultMaxMemory = 32 << 20 // 32 MB
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

	// V1
	apiRouter := router.PathPrefix("/v1").Subrouter()
	// Auth
	apiRouter.HandleFunc("/login", h.logIn()).Methods(http.MethodPost)
	apiRouter.HandleFunc("/login/confirm", h.confirmCode()).Methods(http.MethodPost)
	apiRouter.HandleFunc("/login/refresh", h.refreshTokenHandler()).Methods(http.MethodPost)
	// Profile
	profileRouter := apiRouter.PathPrefix("/profile").Subrouter()
	profileRouter.Use(h.AuthenticateUser)
	profileRouter.HandleFunc("", h.getCurrentProfileHandler()).Methods(http.MethodGet)
	profileRouter.HandleFunc("", h.updateProfileHandler()).Methods(http.MethodPatch)
	profileRouter.HandleFunc("/photo", h.uploadProfilePhotoHandler()).Methods(http.MethodPost)
	profileRouter.HandleFunc("/photo", h.deleteProfilePhotoHandler()).Methods(http.MethodDelete)

	// Product
	productRouter := apiRouter.PathPrefix("/product").Subrouter()
	productRouter.HandleFunc("", h.getProductListHandler()).Methods(http.MethodGet)
	productRouter.HandleFunc("/{id:[0-9]+}", h.getProductByIdHandler()).Methods(http.MethodGet)
	protectedProductRouter := productRouter.PathPrefix("").Subrouter()
	protectedProductRouter.Use(h.AuthenticateUser)
	protectedProductRouter.HandleFunc("", h.createProductHandler()).Methods(http.MethodPost)
	protectedProductRouter.HandleFunc("/{id:[0-9]+}", h.deleteProductHandler()).Methods(http.MethodDelete)
	protectedProductRouter.HandleFunc("/{id:[0-9]+}", h.updateProductHandler()).Methods(http.MethodPatch)
	protectedProductRouter.HandleFunc("/{id:[0-9]+}/photo", h.uploadProductPhotoHandler()).Methods(http.MethodPost)
	protectedProductRouter.HandleFunc("/{id:[0-9]+}/photo", h.deleteProductPhotoHandler()).Methods(http.MethodDelete)

	// Recipe
	recipeRouter := apiRouter.PathPrefix("/recipe").Subrouter()
	protectedRecipeRouter := recipeRouter.PathPrefix("").Subrouter()
	protectedRecipeRouter.Use(h.AuthenticateUser)
	protectedRecipeRouter.HandleFunc("", h.createRecipeHandler()).Methods(http.MethodPost)
	protectedRecipeRouter.HandleFunc("/private", h.getRecipeListPrivateHandler()).Methods(http.MethodGet)
	protectedRecipeRouter.HandleFunc("/{id:[0-9]+}/private", h.getRecipeByIdPrivateHandler()).Methods(http.MethodGet)
	protectedRecipeRouter.HandleFunc("/{id:[0-9]+}", h.deleteRecipeHandler()).Methods(http.MethodDelete)
	protectedRecipeRouter.HandleFunc("/{id:[0-9]+}", h.updateRecipeHandler()).Methods(http.MethodPatch)
	recipeRouter.HandleFunc("", h.getRecipeListHandler()).Methods(http.MethodGet)
	recipeRouter.HandleFunc("/{id:[0-9]+}", h.getRecipeByIdHandler()).Methods(http.MethodGet)

	// RecipeStep
	recipeStepRouter := apiRouter.PathPrefix("/recipe-step").Subrouter()
	protectedRecipeStepRouter := recipeStepRouter.PathPrefix("").Subrouter()
	protectedRecipeStepRouter.Use(h.AuthenticateUser)
	protectedRecipeStepRouter.HandleFunc("", h.createRecipeStepHandler()).Methods(http.MethodPost)
	protectedRecipeStepRouter.HandleFunc("/{id:[0-9]+}", h.deleteRecipeStepHandler()).Methods(http.MethodDelete)
	protectedRecipeStepRouter.HandleFunc("/{id:[0-9]+}", h.updateRecipeStepHandler()).Methods(http.MethodPatch)
	protectedRecipeStepRouter.HandleFunc("/{id:[0-9]+}/photo", h.uploadRecipeStepPhotoHandler()).Methods(http.MethodPost)
	protectedRecipeStepRouter.HandleFunc("/{id:[0-9]+}/photo", h.deleteRecipeStepPhotoHandler()).Methods(http.MethodDelete)
	recipeStepRouter.HandleFunc("", h.getRecipeStepListHandler()).Methods(http.MethodGet)
	recipeStepRouter.HandleFunc("/{id:[0-9]+}", h.getRecipeStepByIdHandler()).Methods(http.MethodGet)

	// RecipeProduct
	recipeProductRouter := apiRouter.PathPrefix("/recipe-product").Subrouter()
	protectedRecipeProductRouter := recipeProductRouter.PathPrefix("").Subrouter()
	protectedRecipeProductRouter.Use(h.AuthenticateUser)
	protectedRecipeProductRouter.HandleFunc("", h.createRecipeProductHandler()).Methods(http.MethodPost)
	protectedRecipeProductRouter.HandleFunc("/{id:[0-9]+}", h.deleteRecipeProductHandler()).Methods(http.MethodDelete)
	protectedRecipeProductRouter.HandleFunc("/{id:[0-9]+}", h.updateRecipeProductHandler()).Methods(http.MethodPatch)
	recipeProductRouter.HandleFunc("", h.getRecipeProductListHandler()).Methods(http.MethodGet)
	recipeProductRouter.HandleFunc("/{id:[0-9]+}", h.getRecipeProductByIdHandler()).Methods(http.MethodGet)

	// RecipeGallery
	recipeGalleryRouter := apiRouter.PathPrefix("/recipe-gallery").Subrouter()
	protectedRecipeGalleryRouter := recipeGalleryRouter.PathPrefix("").Subrouter()
	protectedRecipeGalleryRouter.Use(h.AuthenticateUser)
	protectedRecipeGalleryRouter.HandleFunc("", h.createRecipeGalleryPhotoHandler()).Methods(http.MethodPost)
	protectedRecipeGalleryRouter.HandleFunc("/{id:[0-9]+}", h.deleteRecipeGalleryPhotoHandler()).Methods(http.MethodDelete)
	protectedRecipeGalleryRouter.HandleFunc("/{id:[0-9]+}", h.updateRecipeGalleryPhotoHandler()).Methods(http.MethodPatch)

	// ShoppingList
	shoppingListRouter := apiRouter.PathPrefix("/shopping-item").Subrouter()
	shoppingListRouter.Use(h.AuthenticateUser)
	shoppingListRouter.HandleFunc("", h.getShoppingListHandler()).Methods(http.MethodGet)
	shoppingListRouter.HandleFunc("", h.createShoppingItemHandler()).Methods(http.MethodPost)
	shoppingListRouter.HandleFunc("/{id:[0-9]+}", h.deleteShoppingItemHandler()).Methods(http.MethodDelete)
	shoppingListRouter.HandleFunc("/{id:[0-9]+}", h.updateShoppingItemHandler()).Methods(http.MethodPatch)

	// Diet
	dietRouter := apiRouter.PathPrefix("/diet").Subrouter()
	dietProtectedRouter := dietRouter.PathPrefix("").Subrouter()
	dietProtectedRouter.Use(h.AuthenticateUser)
	dietProtectedRouter.HandleFunc("", h.createDietHandler()).Methods(http.MethodPost)
	dietProtectedRouter.HandleFunc("/private", h.getDietListPrivateHandler()).Methods(http.MethodGet)
	dietProtectedRouter.HandleFunc("/{id:[0-9]+}/private", h.getDietByIdPrivateHandler()).Methods(http.MethodGet)
	dietProtectedRouter.HandleFunc("/{id:[0-9]+}", h.deleteDietHandler()).Methods(http.MethodDelete)
	dietProtectedRouter.HandleFunc("/{id:[0-9]+}", h.updateDietHandler()).Methods(http.MethodPatch)

	// DietItem
	dietItemRouter := apiRouter.PathPrefix("/diet-item").Subrouter()
	dietItemProtectedRouter := dietItemRouter.PathPrefix("").Subrouter()
	dietItemProtectedRouter.Use(h.AuthenticateUser)
	dietItemProtectedRouter.HandleFunc("", h.createDietItemHandler()).Methods(http.MethodPost)
	dietItemProtectedRouter.HandleFunc("/private", h.getDietItemListPrivateHandler()).Methods(http.MethodGet)
	dietItemProtectedRouter.HandleFunc("/{id:[0-9]+}/private", h.getDietItemByIdPrivateHandler()).Methods(http.MethodGet)
	dietItemProtectedRouter.HandleFunc("/{id:[0-9]+}", h.deleteDietItemHandler()).Methods(http.MethodDelete)
	dietItemProtectedRouter.HandleFunc("/{id:[0-9]+}", h.updateDietItemHandler()).Methods(http.MethodPatch)

	return router
}
