package port

import "bitbucket.org/Amartha/poc-chaining-usecase/model"

//go:generate mockgen -source=usecase.go -destination=mock/usecase_mock.go -package=mock
type Usecase interface {
	ChainHandler(h model.UsecaseHandlerFunc, m ...model.UsecaseMiddlewareFunc) model.UsecaseHandlerFunc

	GetCustomerUsecase() CustomerUsecase
}
