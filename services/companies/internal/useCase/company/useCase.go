package company

import (
	"companies/internal/useCase/adapters/storage"
)

type UseCase struct {
	adaptersStorage storage.Company
	options         Options
}

type Options struct{}

func New(storage storage.Company, options Options) *UseCase {
	var uc = &UseCase{
		adaptersStorage: storage,
	}
	uc.SetOptions(options)
	return uc
}

func (uc *UseCase) SetOptions(options Options) {
	if uc.options != options {
		uc.options = options
		//log.Info("set new options", zap.Any("options", uc.options))
	}
}
