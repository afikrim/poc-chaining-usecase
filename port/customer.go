package port

import (
	"context"

	"bitbucket.org/Amartha/poc-chaining-usecase/model"
)

//go:generate mockgen -source=customer.go -destination=mock/customer_mock.go -package=mock
type CustomerUsecase interface {
	CheckXSot(h model.UsecaseHandlerFunc) model.UsecaseHandlerFunc
	PublishEvent(h model.UsecaseHandlerFunc) model.UsecaseHandlerFunc

	GetCustomer(ctx context.Context, in any, opts any) (any, error)
}
