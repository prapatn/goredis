package repositories

import (
	"fmt"
	"math/rand"
	"time"

	"gorm.io/gorm"
)

type ProductRepository interface {
	GetProducts() ([]product, error)
}

type product struct {
	ID       int
	Name     string
	Quantity int
}

func mockData(db *gorm.DB) error {
	var count int64
	db.Model(&product{}).Count(&count)
	if count > 0 {
		return nil
	}

	products := []product{}
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)
	for len(products) < 5000 {
		products = append(products, product{
			Name:     fmt.Sprintf("Product%v", len(products)+1),
			Quantity: random.Intn(100),
		})
	}
	return db.Create(products).Error
}
