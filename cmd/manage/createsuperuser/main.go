package main

import (
	"flag"
	"fmt"
	"food/internal/api_admin/model"
	"food/internal/api_admin/repository/postgres"
	"food/internal/api_admin/service"
	"food/internal/file_service/minio"
	"food/pkg/config"
	"food/pkg/database"
	objectstorage "food/pkg/object_storage"
	"log"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "config/dev.toml", "path to config file")
}

func main() {
	flag.Parse()
	config, err := config.InitConfig(configPath)
	if err != nil {
		log.Fatal(err)
	}
	db, err := database.InitDB(config.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	minioClient, err := objectstorage.InitMinio(config.MinioEndpoint, config.MinioAccessKey, config.MinioSecretKey, config.MinioUseSSL)
	if err != nil {
		log.Fatal(err)
	}
	repo := postgres.NewRepository(db)
	fileService := minio.NewFileServcie(minioClient)
	service := service.NewService(config, repo, fileService)

	loginData := &model.CreateUser{
		IsSuperuser: newBool(true),
		IsStaff:     newBool(true),
	}

	fmt.Print("Login: ")
	if _, err := fmt.Scanln(&loginData.Phone); err != nil {
		log.Fatal(err)
	}
	// TODO hide user input
	fmt.Print("Password: ")
	fmt.Print("\033[8m") // Hide input
	if _, err := fmt.Scanln(&loginData.Password); err != nil {
		log.Fatal(err)
	}
	fmt.Print("\033[28m") // Show input
	// TODO hide user input
	fmt.Print("Password again: ")
	fmt.Print("\033[8m") // Hide input
	if _, err := fmt.Scanln(&loginData.Password2); err != nil {
		log.Fatal(err)
	}
	fmt.Print("\033[28m") // Show input
	_, err = service.AuthService.Create(loginData)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Success!")
}

func newBool(b bool) *bool {
	return &b
}
