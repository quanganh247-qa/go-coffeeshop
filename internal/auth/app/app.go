package app

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/thangchung/go-coffeeshop/cmd/auth/config"
	usecases "github.com/thangchung/go-coffeeshop/internal/auth/usecases/users"
	"github.com/thangchung/go-coffeeshop/internal/pkg/auth"
	"github.com/thangchung/go-coffeeshop/pkg/postgres"
	pkgConsumer "github.com/thangchung/go-coffeeshop/pkg/rabbitmq/consumer"
	pkgPublisher "github.com/thangchung/go-coffeeshop/pkg/rabbitmq/publisher"
	"github.com/thangchung/go-coffeeshop/proto/gen"
)

type App struct {
	Cfg            *config.Config
	PG             postgres.DBEngine
	AMQPConn       *amqp.Connection
	Publisher      pkgPublisher.EventPublisher
	Consumer       pkgConsumer.EventConsumer
	UserUseCase    usecases.UseCase
	UserGRPCServer gen.AuthServiceServer
	AuthInterceptor auth.Interceptor
}

func New(
	cfg *config.Config,
	pg postgres.DBEngine,
	amqpConn *amqp.Connection,
	publisher pkgPublisher.EventPublisher,
	consumer pkgConsumer.EventConsumer,
	userUseCase usecases.UseCase,
	userGRPCServer gen.AuthServiceServer,
	authInterceptor auth.Interceptor,
) *App {
	return &App{
		Cfg:            cfg,
		PG:             pg,
		AMQPConn:       amqpConn,
		Publisher:      publisher,
		Consumer:       consumer,
		UserUseCase:    userUseCase,
		UserGRPCServer: userGRPCServer,
		AuthInterceptor: authInterceptor,
	}
}
