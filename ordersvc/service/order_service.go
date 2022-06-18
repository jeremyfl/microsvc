package service

import (
	"context"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"gitlab.com/jeremylo/microsvc/ordersvc/domain"
	"gitlab.com/jeremylo/microsvc/ordersvc/domain/model"
)

type OrderServiceImpl struct {
	Repository    domain.OrderRepository
	MessageBroker domain.MessageBroker
}

func (cs *OrderServiceImpl) CreateOrder(ctx context.Context, payload *model.Order) error {
	_, span := domain.Tracer.Start(ctx, "CreateOrder")
	defer span.End()

	p, err := cs.Repository.Create(ctx, payload)
	if err != nil {
		log.WithError(err).Errorln("error when creating to db")
		return err
	}

	if err := cs.MessageBroker.Publish(ctx, "order.created", p); err != nil {
		log.WithError(err).Errorln("error when publishing order created")
		return err
	}

	log.WithField("product_id", p.ProductID).Infoln("new order created")

	return nil
}

func (cs *OrderServiceImpl) CancelOrder(ctx context.Context, payload *model.Order) error {
	_, span := domain.Tracer.Start(ctx, "CreateOrder")
	defer span.End()

	p := model.Order{
		Model:         gorm.Model{
			ID: payload.ID,
		},
	}
	if err := cs.Repository.Update(ctx, &p, &model.Order{ProductID: payload.ProductID, IsCanceled: true}); err != nil {
		log.WithError(err).Errorln("error when update collection to db")

		return err
	}

	log.WithField("product_id", payload.ProductID).Infoln("order canceled")

	return nil
}
