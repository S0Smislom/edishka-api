package main

import (
	"context"
	"flag"
	"food/internal/api/handler"
	"food/internal/api/repository/postgres"
	"food/internal/api/server"
	"food/internal/api/service"
	"food/pkg/config"
	"food/pkg/database"
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

// @title Food API
// @version 1.0
// @description API Server for FOOD Application

// @host localhost:8080
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

	repo := postgres.NewRepository(db)
	service := service.NewService(repo)
	handler := handler.NewHandler(service)
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
