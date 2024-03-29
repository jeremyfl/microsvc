package event

import (
	"context"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"gitlab.com/jeremylo/microsvc/ordersvc/domain"
	"gitlab.com/jeremylo/microsvc/ordersvc/domain/model"
)

type Handler struct {
	Services domain.Services
}

func (h *Handler) GenericHandler(message []byte) error {
	return nil
}

func (h *Handler) StockExceededConsumer(message []byte) error {
	ctx, span := domain.Tracer.Start(context.Background(), "CreateOrder")
	defer span.End()

	var canceledOrder model.Order
	if err := json.Unmarshal(message, &canceledOrder); err != nil {
		return err
	}

	log.WithField("product_id", canceledOrder.ProductID).Infoln("there's canceled order event")

	if err := h.Services.CancelOrder(ctx, &canceledOrder); err != nil {
		return err
	}

	return nil
}