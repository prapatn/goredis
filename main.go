package main

import (
	"goredis/handler"
	"goredis/repositories"
	"goredis/services"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db := initDatabase()
	redisClient := initRedis()

	productRepo := repositories.NewProductRepositoryDB(db)
	productService := services.NewCatalogService(productRepo)
	productHandler := handler.NewCatalogHandlerRedis(productService, redisClient)

	app := fiber.New()

	app.Get("/products", productHandler.GetProducts)

	app.Listen(":8000")

}

func initDatabase() *gorm.DB {
	dial := mysql.Open("root:example@tcp(localhost:3306)/example")
	db, err := gorm.Open(dial, &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func initRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}
