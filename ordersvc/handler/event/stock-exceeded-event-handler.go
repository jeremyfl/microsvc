package event

import (
	"gitlab.com/jeremylo/microsvc/ordersvc/domain"
	"log"
)

type Handler struct {
	Services domain.Services
}

func (h *Handler) ProductDeleted(message []byte) error {
	log.Println("product.deleted consumer")

	return nil
}

func (h *Handler) StockExceededConsumer(message []byte) error {
	log.Println("stock.exceeded-amount consumer")

	return nil
}