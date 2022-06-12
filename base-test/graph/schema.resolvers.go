package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"customer/graph/generated"
	"customer/graph/model"
	"strconv"
)

func (r *queryResolver) Products(ctx context.Context) ([]*model.Product, error) {
	productStore := r.Services.ProductService.FetchProduct(ctx)

	products := make([]*model.Product, 0)

	for _, product := range productStore {
		productId := strconv.FormatUint(uint64(product.ID), 10)

		products = append(products, &model.Product{
			ID:          &productId,
			Name:        product.Name,
			Price:       product.Price,
			Description: product.Description,
			Rating:      product.Rating,
			Category:    product.Category,
			Brand:       product.Brand,
			Thumbnail:   product.Thumbnail,
		})
	}

	return products, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
type mutationResolver struct{ *Resolver }
