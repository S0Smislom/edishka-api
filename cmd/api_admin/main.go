package main

import (
	"context"
	"flag"
	"food/internal/api/server"
	"food/internal/api_admin/handler"
	"food/internal/api_admin/repository/postgres"
	"food/internal/api_admin/service"
	"food/internal/file_service/minio"
	"food/pkg/config"
	"food/pkg/database"
	objectstorage "food/pkg/object_storage"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "config/dev.toml", "path to config file")
}

// @title Food Admin API
// @version 1.0
// @description Admin API Server for FOOD Application

// @host localhost:8092
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	flag.Parse()
	logrus.SetFormatter(new(logrus.JSONFormatter))
	config, err := config.InitConfig(configPath)
	if err != nil {
		log.Fatal(err)
	}
	db, err := database.InitDB(config.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	minioClient, err := objectstorage.InitMinio(
		config.MinioEndpoint,
		config.MinioAccessKey,
		config.MinioSecretKey,
		config.MinioUseSSL,
	)

	repo := postgres.NewRepository(db)
	fileService := minio.NewFileServcie(minioClient)
	service := service.NewService(config, repo, fileService)
	handler := handler.NewHandler(config, service)

	srv := new(server.Server)
	go func() {
		if err := srv.Run(config.AdminAddr, handler.InitRoutes()); err != nil {
			log.Fatal(err)
		}
	}()
	logrus.Print("Admin API Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("Admin API Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}
}
