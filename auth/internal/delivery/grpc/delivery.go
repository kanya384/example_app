package grpc

import (
	auth "auth/internal/delivery/grpc/interface"
	"auth/internal/useCase"
	"time"
)

type Delivery struct {
	auth.UnimplementedAuthServer
	ucAuth  useCase.UseCase
	salt    string
	options Options
}

type Options struct {
	DefaultTimeout time.Duration
}

func New(ucAuth useCase.UseCase, salt string, o Options) *Delivery {
	var d = &Delivery{
		ucAuth: ucAuth,
		salt:   salt,
	}

	d.SetOptions(o)
	return d
}

func (d *Delivery) SetOptions(options Options) {
	if d.options != options {
		d.options = options
	}
}
