package event

import (
	"context"
	"encoding/json"
	"gitlab.com/jeremylo/microsvc/stocksvc/domain"
)

type Handler struct {
	Services domain.Services
}

type OrderCreatedPayload struct {
	ProductID int
}

func (h *Handler) GenericHandler(message []byte) error {
	return nil
}

func (h *Handler) OrderCreated(message []byte) error {
	ctx, span := domain.Tracer.Start(context.Background(), "CreateOrder")
	defer span.End()

	var ocp OrderCreatedPayload

	if err := json.Unmarshal(message, &ocp); err != nil {
		return err
	}

	if err := h.Services.DecreaseStock(ctx, ocp.ProductID); err != nil {
		return err
	}

	return nil
}

