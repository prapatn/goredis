package services

import (
	"context"
	"encoding/json"
	"goredis/repositories"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

type catalogServiceRedis struct {
	productRepo repositories.ProductRepository
	redisClient *redis.Client
}

func NewCatalogServiceRedis(repo repositories.ProductRepository, redis *redis.Client) CatalogServices {
	return catalogServiceRedis{productRepo: repo, redisClient: redis}
}

func (s catalogServiceRedis) GetProducts() (products []Product, err error) {
	key := "repository::GetProducts"
	//Redis chk for get
	dataChk, err := s.redisClient.Get(context.Background(), key).Result()
	if err == nil {
		err = json.Unmarshal([]byte(dataChk), &products)
		if err != nil {
			return nil, err
		}
		log.Println("redis")
		return products, nil
	}

	//get database
	productsDB, err := s.productRepo.GetProducts()
	if err != nil {
		return nil, err
	}

	for _, p := range productsDB {
		products = append(products, Product(p))
	}

	dataJson, err := json.Marshal(products)
	if err != nil {
		return nil, err
	}

	//set
	s.redisClient.Set(context.Background(), key, string(dataJson), time.Second*10).Err()
	log.Println("database")

	return products, nil
}
