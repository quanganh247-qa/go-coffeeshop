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
			ID:        x.ID.String(),
			Name:      x.Name,
			Type:      int(x.Type),
			Price:     x.Price,
			Image:     x.Image,
			CreatedAt: x.CreatedAt.Time.String(),
			UpdatedAt: x.UpdatedAt.Time.String(),
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

func (p *productPostgresRepo) GetProductByID(ctx context.Context, id string) (*domain.ItemTypeDto, error) {
	querier := postgresql.New(p.pg.GetDB())

	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.Wrap(err, "uuid.Parse")
	}

	product, err := querier.GetProductByID(ctx, uid)
	if err != nil {
		return nil, errors.Wrap(err, "querier.GetProductByID")
	}

	return &domain.ItemTypeDto{
		ID:        product.ID.String(),
		Name:      product.Name,
		Type:      int(product.Type),
		Price:     product.Price,
		Image:     product.Image,
		CreatedAt: product.CreatedAt.Time.String(),
		UpdatedAt: product.UpdatedAt.Time.String(),
	}, nil
}

func (p *productPostgresRepo) Create(ctx context.Context, name string, itemType int32, price float64, image string) (string, error) {
	querier := postgresql.New(p.pg.GetDB())

	id, err := querier.CreateProduct(ctx, postgresql.CreateProductParams{
		Name:  name,
		Type:  itemType,
		Price: price,
		Image: image,
		Stock: 0,
	})
	if err != nil {
		return "", errors.Wrap(err, "querier.CreateProduct")
	}

	return id.String(), nil
}
