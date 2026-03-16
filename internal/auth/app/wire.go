//go:build wireinject
// +build wireinject

package app

import (
	"database/sql"

	"github.com/google/wire"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/thangchung/go-coffeeshop/cmd/auth/config"
	"github.com/thangchung/go-coffeeshop/internal/auth/app/router"
	"github.com/thangchung/go-coffeeshop/internal/auth/domain"
	"github.com/thangchung/go-coffeeshop/internal/auth/infras/hasher"
	"github.com/thangchung/go-coffeeshop/internal/auth/infras/repo"
	"github.com/thangchung/go-coffeeshop/internal/auth/infras/token"
	"github.com/thangchung/go-coffeeshop/internal/auth/usecases/users"
	"github.com/thangchung/go-coffeeshop/pkg/postgres"
	"github.com/thangchung/go-coffeeshop/pkg/rabbitmq"
	pkgConsumer "github.com/thangchung/go-coffeeshop/pkg/rabbitmq/consumer"
	pkgPublisher "github.com/thangchung/go-coffeeshop/pkg/rabbitmq/publisher"
	"google.golang.org/grpc"
)

func InitApp(
	cfg *config.Config,
	dbConnStr postgres.DBConnString,
	rabbitMQConnStr rabbitmq.RabbitMQConnStr,
	grpcServer *grpc.Server,
) (*App, func(), error) {
	panic(wire.Build(
		New,
		dbEngineFunc,
		sqlDBFunc,
		rabbitMQFunc,
		pkgPublisher.EventPublisherSet,
		pkgConsumer.EventConsumerSet,
		repo.NewUserRepository,
		repo.NewRefreshTokenRepository,
		hasherFunc,
		tokenGeneratorFunc,
		usecases.NewUserService,
		router.AuthGRPCServiceSet,
	))
}

func dbEngineFunc(url postgres.DBConnString) (postgres.DBEngine, func(), error) {
	db, err := postgres.NewPostgresDB(url)
	if err != nil {
		return nil, nil, err
	}
	return db, func() { db.Close() }, err
}

func sqlDBFunc(db postgres.DBEngine) *sql.DB {
	return db.GetDB()
}

func rabbitMQFunc(url rabbitmq.RabbitMQConnStr) (*amqp.Connection, func(), error) {
	conn, err := rabbitmq.NewRabbitMQConn(url)
	if err != nil {
		return nil, nil, err
	}
	return conn, func() { conn.Close() }, err
}

func hasherFunc() domain.PasswordHasher {
	return hasher.NewBcryptHasher(0)
}

func tokenGeneratorFunc(cfg *config.Config) domain.TokenGenerator {
	return token.NewJWTTokenGenerator(cfg.JWT.SecretKey)
}
