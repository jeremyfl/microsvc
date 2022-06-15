package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"stocksvc/domain"
)

type OrderCreatedHandler struct {
	Services domain.Services
}

type OrderCreatedPayload struct {
	ProductID int
}

func (oc *OrderCreatedHandler) Handle(ctx context.Context, message []byte) error {
	var ocp OrderCreatedPayload

	if err := json.Unmarshal(message, &ocp); err != nil {
		return err
	}

	fmt.Println(ocp)

	if err := oc.Services.DecreaseStock(ctx, ocp.ProductID); err != nil {
		return err
	}

	return nil
}
