package main

import (
	"context"
	"flag"
	"food/internal/api/handler"
	"food/internal/api/repository/gormrepo"
	"food/internal/api/server"
	"food/internal/api/service"
	"food/internal/file_service/minio"
	"food/pkg/config"
	"food/pkg/database"
	miniofileprovider "food/pkg/file_provider/minio_file_provider"
	objectstorage "food/pkg/object_storage"
	"log"
	"os"
	"os/signal"
	"syscall"

	docs "food/docs/api"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", ".env", "path to config file")
}

// @title Food API
// @version 1.0
// @description API Server for FOOD Application

// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	flag.Parse()
	logrus.SetFormatter(new(logrus.JSONFormatter))
	config, err := config.InitEnvConfig(configPath)
	if err != nil {
		log.Fatal(err)
	}
	db, err := database.InitDB(config.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	minioClient, err := objectstorage.InitMinio(
		config.MinioEndpoint,
		config.MinioAccessKey,
		config.MinioSecretKey,
		config.MinioUseSSL,
	)
	if err != nil {
		log.Fatal(err)
	}
	fileProvider := miniofileprovider.NewMinioFileProvider(minioClient)
	fileService := minio.NewFileServcie(fileProvider)

	docs.SwaggerInfo.Host = config.BaseHost

	// repo := postgres.NewRepository(db)
	repo := gormrepo.NewRepository(gormDB)
	service := service.NewService(repo, fileService, config)
	handler := handler.NewHandler(config, service)
	srv := new(server.Server)
	go func() {
		if err := srv.Run(config.BindAddr, handler.InitRoutes()); err != nil {
			log.Fatal(err)
		}
	}()
	logrus.Print("API Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("API Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}
}
