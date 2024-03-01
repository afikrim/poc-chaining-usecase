package port

import (
	"context"

	"bitbucket.org/Amartha/poc-chaining-usecase/model"
)

//go:generate mockgen -source=customer.go -destination=mock/customer_mock.go -package=mock
type CustomerUsecase interface {
	GetCustomer(ctx context.Context, in *model.GetCustomerIn) (*model.GetCustomerOut, error)
}
