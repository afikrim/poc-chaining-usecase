package usecase

import (
	"bitbucket.org/Amartha/poc-chaining-usecase/port"
)

var _ port.Usecase = (*usecases)(nil)

type usecases struct {
	common usecase

	customerUsecase port.CustomerUsecase
}

type usecase struct {
	ucs *usecases
}

func NewUsecase() port.Usecase {
	uc := &usecases{}

	uc.customerUsecase = (*customerUsecase)(&uc.common)
	return uc
}
