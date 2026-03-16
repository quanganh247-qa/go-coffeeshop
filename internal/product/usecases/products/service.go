package products

import (
	"context"
	"strings"

	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/thangchung/go-coffeeshop/internal/product/domain"
)

var _ UseCase = (*service)(nil)

var UseCaseSet = wire.NewSet(NewService)

type service struct {
	repo domain.ProductRepo
}

func NewService(repo domain.ProductRepo) UseCase {
	return &service{
		repo: repo,
	}
}

func (s *service) GetItemTypes(ctx context.Context) ([]*domain.ItemTypeDto, error) {
	results, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "service.GetItemTypes")
	}

	return results, nil
}

func (s *service) GetItemsByType(ctx context.Context, itemTypes string) ([]*domain.ItemDto, error) {
	types := strings.Split(itemTypes, ",")

	results, err := s.repo.GetByTypes(ctx, types)
	if err != nil {
		return nil, errors.Wrap(err, "service.GetItemsByType")
	}

	return results, nil
}

func (s *service) GetItemDetailByID(ctx context.Context, id string) (*domain.ItemTypeDto, error) {
	result, err := s.repo.GetProductByID(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "service.GetItemDetailByID")
	}

	return result, nil
}

func (s *service) CreateProduct(ctx context.Context, name string, itemType int32, price float64, image string) (string, error) {
	id, err := s.repo.Create(ctx, name, itemType, price, image)
	if err != nil {
		return "", errors.Wrap(err, "service.CreateProduct")
	}

	return id, nil
}
