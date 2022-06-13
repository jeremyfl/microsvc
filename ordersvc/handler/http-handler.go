package handler

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"ordersvc/domain"
	"ordersvc/domain/model"
)

type Handler struct {
	Services domain.Services
}

func (s *Handler) Handle(c *fiber.Ctx) error {
	ctx := context.Background()

	p := model.Order{
		ProductID:     0,
		TotalPrice:    0,
		TotalQuantity: 0,
	}

	if err := s.Services.StockService.CreateOrder(ctx, &p); err != nil {
		return c.SendString(err.Error())
	}

	return c.SendString("root")
}
