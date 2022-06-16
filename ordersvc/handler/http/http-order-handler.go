package http

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/jeremylo/microsvc/ordersvc/domain"
	"gitlab.com/jeremylo/microsvc/ordersvc/domain/model"
	"net/http"
)

type Handler struct {
	Services domain.Services
}

func (s *Handler) Handle(c *fiber.Ctx) error {
	ctx, span := domain.Tracer.Start(c.UserContext(), "Handler")
	defer span.End()

	p := new(model.Order)
	if err := c.BodyParser(p); err != nil {
		return err
	}

	if err := s.Services.OrderService.CreateOrder(ctx, p); err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(http.StatusCreated).SendString("Order created")
}
