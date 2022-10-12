package pubsub

import (
	"notifications/internal/useCase"

	"notifications/pkg/kafka/server"
	"notifications/pkg/logger"
	"time"
)

type Delivery struct {
	mailsUseCase useCase.Mails
	kServer      *server.KafkaServer
	logger       *logger.Logger
	options      Options
}

type Options struct {
	DefaultTimeout time.Duration
}

func New(mailsUseCase useCase.Mails, kServer *server.KafkaServer, logger *logger.Logger, o Options) *Delivery {
	var d = &Delivery{
		mailsUseCase: mailsUseCase,
		kServer:      kServer,
		logger:       logger,
	}

	d.SetOptions(o)
	return d
}

func (d *Delivery) SetOptions(options Options) {
	if d.options != options {
		d.options = options
	}
}
