package handler

import (
	"goredis/services"

	"github.com/gofiber/fiber/v2"
)

type catalogHandler struct {
	service services.CatalogServices
}

func NewCatalogHandler(service services.CatalogServices) CatalogHandler {
	return catalogHandler{service: service}
}

func (h catalogHandler) GetProducts(c *fiber.Ctx) error {
	products, err := h.service.GetProducts()
	if err != nil {
		return err
	}
	res := fiber.Map{
		"status": fiber.StatusOK,
		"data":   products,
	}
	return c.JSON(res)
}
