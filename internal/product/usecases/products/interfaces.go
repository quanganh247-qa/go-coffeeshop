package products

import (
	"context"

	"github.com/thangchung/go-coffeeshop/internal/product/domain"
)

type UseCase interface {
	GetItemTypes(context.Context) ([]*domain.ItemTypeDto, error)
	GetItemsByType(context.Context, string) ([]*domain.ItemDto, error)
	GetItemDetailByID(context.Context, string) (*domain.ItemTypeDto, error)
	CreateProduct(ctx context.Context, name string, itemType int32, price float64, image string) (string, error)
}
