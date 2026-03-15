//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	"github.com/thangchung/go-coffeeshop/cmd/product/config"
	"github.com/thangchung/go-coffeeshop/internal/product/app/router"
	"github.com/thangchung/go-coffeeshop/internal/product/infras/repo"
	productsUC "github.com/thangchung/go-coffeeshop/internal/product/usecases/products"
	"github.com/thangchung/go-coffeeshop/pkg/postgres"
	"google.golang.org/grpc"
)

func InitApp(
	cfg *config.Config,
	dbConnStr postgres.DBConnString,
	grpcServer *grpc.Server,
) (*App, func(), error) {
	panic(wire.Build(
		New,
		dbEngineFunc,
		router.ProductGRPCServerSet,
		repo.RepositorySet,
		productsUC.UseCaseSet,
	))
}

func dbEngineFunc(url postgres.DBConnString) (postgres.DBEngine, func(), error) {
	db, err := postgres.NewPostgresDB(url)
	if err != nil {
		return nil, nil, err
	}
	return db, func() { db.Close() }, nil
}
