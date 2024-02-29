package usecase

import (
	"bitbucket.org/Amartha/poc-chaining-usecase/model"
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

// ChainHandler implements port.Usecase.
func (u *usecases) ChainHandler(h model.UsecaseHandlerFunc, m ...model.UsecaseMiddlewareFunc) model.UsecaseHandlerFunc {
	// if at the end of the middleware then return handler
	if len(m) == 0 {
		return h
	}

	// else wrap handler with middleware
	return m[0](u.ChainHandler(h, m[1:cap(m)]...))
}
