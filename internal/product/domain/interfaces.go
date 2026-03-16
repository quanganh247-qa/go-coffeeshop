package domain

import (
	"context"
)

type (
	ProductRepo interface {
		GetAll(context.Context) ([]*ItemTypeDto, error)
		GetByTypes(context.Context, []string) ([]*ItemDto, error)
		GetProductByID(context.Context, string) (*ItemTypeDto, error)
		Create(ctx context.Context, name string, itemType int32, price float64, image string) (string, error)
	}
)
