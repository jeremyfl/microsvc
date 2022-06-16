package event

import (
	"context"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"gitlab.com/jeremylo/microsvc/stocksvc/domain"
	"gitlab.com/jeremylo/microsvc/stocksvc/domain/model"
)

type Handler struct {
	Services domain.Services
}

func (h *Handler) GenericHandler(message []byte) error {
	return nil
}

func (h *Handler) OrderCreated(message []byte) error {
	ctx, span := domain.Tracer.Start(context.Background(), "CreateOrder")
	defer span.End()

	var op model.OrderPayload
	if err := json.Unmarshal(message, &op); err != nil {
		log.WithError(err).Errorln("error when do unmarshall")

		return err
	}

	log.WithField("product_id", op.ProductID).Infoln("there's order created event, decreasing a stock..")

	if err := h.Services.DecreaseStock(ctx, &op); err != nil {
		log.WithError(err).Errorln("error when decreasing a stock")

		return err
	}

	return nil
}
