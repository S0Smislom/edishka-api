package test

// import (
// 	"database/sql"
// 	"food/internal/api/handler"
// 	"food/internal/api/repository"
// 	"food/internal/api/repository/postgres"
// 	"food/internal/api/service"
// 	fileservice "food/internal/file_service"
// 	"food/internal/file_service/mockfileservice"
// 	"food/pkg/config"
// 	"food/pkg/database"
// 	"net/http"
// )

// func InitTestServer() (http.Handler, *service.Service, repository.Repository, fileservice.FileService, *config.Config, *sql.DB, error) {
// 	config, _ := config.InitTestConfig()
// 	db, err := database.InitTestDB(config.DatabaseURL)
// 	if err != nil {
// 		return nil, nil, nil, nil, nil, nil, err
// 		// t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}
// 	// defer database.TeardownTestDB(db, "product")
// 	fileService := mockfileservice.NewMockFileServicez()
// 	repo := postgres.NewRepository(db)
// 	service := service.NewService(repo, fileService, config)
// 	handler := handler.NewHandler(config, service)

// 	s := handler.InitRoutes()
// 	return s, service, repo, fileService, config, db, nil
// }
