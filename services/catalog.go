package services

type CatalogServices interface {
	GetProducts() ([]Product, error)
}

type Product struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}
