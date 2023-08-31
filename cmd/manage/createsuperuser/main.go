package main

import (
	"flag"
	"fmt"
	"food/internal/api_admin/model"
	"food/internal/api_admin/repository/postgres"
	"food/internal/api_admin/service"
	"food/pkg/config"
	"food/pkg/database"
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
	repo := postgres.NewRepository(db)
	service := service.NewService(config, repo)

	loginData := &model.CreateUser{
		IsSuperuser: newBool(true),
		IsStaff:     newBool(true),
	}

	fmt.Print("Login: ")
	if _, err := fmt.Scanln(&loginData.Phone); err != nil {
		log.Fatal(err)
	}
	fmt.Print("Password: ")
	if _, err := fmt.Scanln(&loginData.Password); err != nil {
		log.Fatal(err)
	}
	fmt.Print("Password again: ")
	if _, err := fmt.Scanln(&loginData.Password2); err != nil {
		log.Fatal(err)
	}
	_, err = service.AuthService.Create(loginData)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Success!")
}

func newBool(b bool) *bool {
	return &b
}
