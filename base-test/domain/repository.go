package domain

import (
	"context"
	"customer/domain/model"
)

type ProductRepository interface {
	Get(ctx context.Context) *[]model.Product
}