package app

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"gorm.io/gorm"
)

type AppCommons struct {
	DB               *gorm.DB
	Logger           *Logger
	Validator        *Validator
	RabbitConnection *amqp.Connection
}

func NewAppCommons(db *gorm.DB,
	logger *Logger,
	validator *Validator,
	rabbitConnection *amqp.Connection,
) *AppCommons {
	return &AppCommons{
		DB:               db,
		Logger:           logger,
		Validator:        validator,
		RabbitConnection: rabbitConnection,
	}
}
