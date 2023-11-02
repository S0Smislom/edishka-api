package main

import (
	"flag"
	"fmt"
	"food/internal/api/model"
	"food/internal/api/repository/postgres"
	"food/pkg/config"
	"food/pkg/database"
	"log"
	"math/rand"
	"time"
)

var (
	configPath string
	userId     int
	randItems  int
	products   = []*model.CreateProduct{
		{
			Title:         "Test",
			Slug:          "test",
			Calories:      999,
			Squirrels:     999,
			Fats:          999,
			Carbohydrates: 999,
		},
	}
)

func init() {
	flag.StringVar(&configPath, "config-path", "config/dev.toml", "path to config file")
	flag.IntVar(&userId, "user-id", 1, "user created by")
	flag.IntVar(&randItems, "rand-items", 10, "amount of random items")
}

func main() {
	rand.Seed(time.Now().UnixNano())
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
	for _, product := range products {
		product.CreatedById = userId

		for i := 0; i < randItems; i++ {
			index := rand.Intn(1001)
			product.Title = fmt.Sprintf("%s-%d", product.Title, index)
			product.Slug = fmt.Sprintf("%s-%d", product.Slug, index)
			id, err := repo.Product().Create(product)
			if err != nil {
				log.Printf("Error: %s", err)
			} else {
				log.Printf("Created Product(%d)\n", id)
			}
		}
	}

}
