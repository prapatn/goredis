package repositories

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type productRepositoryRedis struct {
	db          *gorm.DB
	redisClient *redis.Client
}

func NewProductRepositoryRedis(db *gorm.DB, redisClient *redis.Client) ProductRepository {
	db.AutoMigrate(&product{})
	mockData(db)
	return productRepositoryRedis{db: db, redisClient: redisClient}
}

func (r productRepositoryRedis) GetProducts() (products []product, err error) {
	key := "repository::GetProducts"
	//chk
	dataChk, err := r.redisClient.Get(context.Background(), key).Result()
	if err == nil {
		err = json.Unmarshal([]byte(dataChk), &products)
		if err != nil {
			return nil, err
		}
		log.Println("redis")
		return products, nil
	}

	err = r.db.Order("quantity DESC").Limit(30).Find(&products).Error
	if err != nil {
		return nil, err
	}

	dataJson, err := json.Marshal(products)
	if err != nil {
		return nil, err
	}

	//set
	err = r.redisClient.Set(context.Background(), key, string(dataJson), time.Second*10).Err()
	if err != nil {
		return nil, err
	}

	log.Println("database")

	return products, nil
}
