package usecase

import (
	"context"
	"fmt"
	"log"

	"bitbucket.org/Amartha/poc-chaining-usecase/model"
	"bitbucket.org/Amartha/poc-chaining-usecase/port"
)

var _ port.CustomerUsecase = (*customerUsecase)(nil)

type customerUsecase usecase

func (u *usecases) GetCustomerUsecase() port.CustomerUsecase {
	return u.customerUsecase
}

// GetCustomer implements port.CustomerUsecase.
func (c *customerUsecase) GetCustomer(ctx context.Context, in *model.GetCustomerIn) (*model.GetCustomerOut, error) {
	log.SetPrefix("port.CustomerUsecase.GetCustomer\t")

	if prefix, ok := ctx.Value("prefix").(string); ok {
		log.Printf("prefix is setted, %s", prefix)
		in.ID = fmt.Sprintf("%s-%s", prefix, in.ID)
	}

	log.Printf("return customer, %s", in.ID)
	return &model.GetCustomerOut{
		ID:   in.ID,
		Name: fmt.Sprintf("Dummy: %s", in.ID),
	}, nil
}
