package usecase

import (
	"context"

	"bitbucket.org/Amartha/poc-chaining-usecase/model"
	"bitbucket.org/Amartha/poc-chaining-usecase/port"
)

func Customer_GetCustomer_Handler(uc port.Usecase, ctx context.Context, in *model.GetCustomerIn, m model.UsecaseMiddlewareFunc[*model.GetCustomerIn, *model.GetCustomerOut]) (*model.GetCustomerOut, error) {
	if m == nil {
		return uc.GetCustomerUsecase().GetCustomer(ctx, in)
	}

	info := &model.UsecaseInfo{
		Uc: uc,
	}

	handler := func(ctx context.Context, in *model.GetCustomerIn) (*model.GetCustomerOut, error) {
		return uc.GetCustomerUsecase().GetCustomer(ctx, in)
	}

	return m(ctx, in, info, handler)
}
