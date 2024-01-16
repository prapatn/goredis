package handler

import (
	"context"
	"encoding/json"
	"goredis/services"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

type catalogHandlerRedis struct {
	service     services.CatalogServices
	redisClient *redis.Client
}

func NewCatalogHandlerRedis(service services.CatalogServices, redisClient *redis.Client) CatalogHandler {
	return catalogHandlerRedis{service: service, redisClient: redisClient}
}

func (h catalogHandlerRedis) GetProducts(c *fiber.Ctx) error {
	key := "repository::GetProducts"
	//Redis chk for get
	dataChk, err := h.redisClient.Get(context.Background(), key).Result()
	if err == nil {
		log.Println("redis")
		c.Set("Content-Type", "application/json")
		return c.SendString(dataChk)
	}

	products, err := h.service.GetProducts()
	if err != nil {
		return err
	}
	res := fiber.Map{
		"status": fiber.StatusOK,
		"data":   products,
	}

	dataJson, err := json.Marshal(res)
	if err != nil {
		return err
	}

	//set
	h.redisClient.Set(context.Background(), key, string(dataJson), time.Second*10).Err()
	log.Println("database")
	return c.JSON(res)
}
