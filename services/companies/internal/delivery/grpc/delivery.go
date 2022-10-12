package grpc

import (
	company "companies/internal/delivery/grpc/interface"
	"companies/internal/useCase"
	"time"
)

type Delivery struct {
	company.UnimplementedCompanyServer
	ucCompany useCase.Company

	options Options
}

type Options struct {
	DefaultTimeout time.Duration
}

func New(ucCompany useCase.Company, o Options) *Delivery {
	var d = &Delivery{
		ucCompany: ucCompany,
	}

	d.SetOptions(o)
	return d
}

func (d *Delivery) SetOptions(options Options) {
	if d.options != options {
		d.options = options
	}
}
