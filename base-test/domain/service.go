package domain

import (
	"context"
	"customer/domain/model"
)

type Services struct {
	AuthService
	ProductService
}

type AuthService interface {
	Login(payload ...string) error
}

type ProductService interface {
	FetchProduct(ctx context.Context) *[]model.Product
}
