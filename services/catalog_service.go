package services

import (
	"goredis/repositories"
)

type catalogService struct {
	productRepo repositories.ProductRepository
}

func NewCatalogService(repo repositories.ProductRepository) CatalogServices {
	return catalogService{productRepo: repo}
}

func (s catalogService) GetProducts() (products []Product, err error) {
	productsDB, err := s.productRepo.GetProducts()
	if err != nil {
		return nil, err
	}

	for _, p := range productsDB {
		products = append(products, Product{
			ID:       p.ID,
			Name:     p.Name,
			Quantity: p.Quantity,
		})
	}

	return products, nil
}
