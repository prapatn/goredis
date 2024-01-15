package main

import (
	"goredis/repositories"
	"log"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db := initDatabase()
	redisClient := initRedis()

	productRepo := repositories.NewProductRepositoryRedis(db, redisClient)

	prod, err := productRepo.GetProducts()

	if err != nil {
		println(err)
		return
	}
	log.Println(prod)
	// app := fiber.New()

	// app.Get("/hello", func(c *fiber.Ctx) error {
	// 	time.Sleep(time.Millisecond * 10)
	// 	return c.SendString("Hello world")
	// })

	// app.Listen(":8000")

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
