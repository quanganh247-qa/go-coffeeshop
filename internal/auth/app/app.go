package app

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/thangchung/go-coffeeshop/cmd/auth/config"
	"github.com/thangchung/go-coffeeshop/pkg/postgres"
	pkgConsumer "github.com/thangchung/go-coffeeshop/pkg/rabbitmq/consumer"
	pkgPublisher "github.com/thangchung/go-coffeeshop/pkg/rabbitmq/publisher"
)

type App struct {
	Cfg       *config.Config
	PG        postgres.DBEngine
	AMQPConn  *amqp.Connection
	Publisher pkgPublisher.EventPublisher
	Consumer  pkgConsumer.EventConsumer
}

func New(
	cfg *config.Config,
	pg postgres.DBEngine,
	amqpConn *amqp.Connection,
	publisher pkgPublisher.EventPublisher,
	consumer pkgConsumer.EventConsumer,

) *App {
	return &App{
		Cfg: cfg,

		PG:        pg,
		AMQPConn:  amqpConn,
		Publisher: publisher,
		Consumer:  consumer,
	}
}
