package repo

import (
	"context"

	"github.com/google/uuid"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/samber/lo"
	"github.com/thangchung/go-coffeeshop/internal/product/domain"
	"github.com/thangchung/go-coffeeshop/internal/product/infras/postgresql"
	"github.com/thangchung/go-coffeeshop/pkg/postgres"
)

type productPostgresRepo struct {
	pg postgres.DBEngine
}

var _ domain.ProductRepo = (*productPostgresRepo)(nil)

var RepositorySet = wire.NewSet(NewProductRepo)

func NewProductRepo(pg postgres.DBEngine) domain.ProductRepo {
	return &productPostgresRepo{pg: pg}
}

func (p *productPostgresRepo) GetAll(ctx context.Context) ([]*domain.ItemTypeDto, error) {
	querier := postgresql.New(p.pg.GetDB())

	products, err := querier.ListProducts(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "querier.ListProducts")
	}

	return lo.Map(products, func(x postgresql.ProductProduct, _ int) *domain.ItemTypeDto {
		return &domain.ItemTypeDto{
			Name:  x.Name,
			Type:  int(x.Type),
			Price: x.Price,
			Image: x.Image,
		}
	}), nil
}

func (p *productPostgresRepo) GetByTypes(ctx context.Context, itemTypes []string) ([]*domain.ItemDto, error) {
	querier := postgresql.New(p.pg.GetDB())

	products, err := querier.GetProductsByNames(ctx, itemTypes)
	if err != nil {
		return nil, errors.Wrap(err, "querier.GetProductsByNames")
	}

	return lo.Map(products, func(x postgresql.ProductProduct, _ int) *domain.ItemDto {
		return &domain.ItemDto{
			Price: x.Price,
			Type:  int(x.Type),
		}
	}), nil
}

func (p *productPostgresRepo) GetProductByID(ctx context.Context, id string) (*domain.ItemDto, error) {
	querier := postgresql.New(p.pg.GetDB())

	idUUID, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.Wrap(err, "uuid.Parse")
	}

	product, err := querier.GetProductByID(ctx, idUUID)
	if err != nil {
		return nil, errors.Wrap(err, "querier.GetProductByID")
	}

	createdAt := product.CreatedAt.Time.Format("2006-01-02 15:04:05")
	updatedAt := product.UpdatedAt.Time.Format("2006-01-02 15:04:05")

	return &domain.ItemDto{
		Price:     product.Price,
		Type:      int(product.Type),
		Name:      product.Name,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}, nil
}
