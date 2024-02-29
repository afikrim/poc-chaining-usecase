package usecase

import (
	"context"
	"errors"
	"fmt"
	"log"

	"bitbucket.org/Amartha/poc-chaining-usecase/model"
	customeroption "bitbucket.org/Amartha/poc-chaining-usecase/model/customer"
	"bitbucket.org/Amartha/poc-chaining-usecase/port"
)

var _ port.CustomerUsecase = (*customerUsecase)(nil)

type customerUsecase usecase

func (u *usecases) GetCustomerUsecase() port.CustomerUsecase {
	return u.customerUsecase
}

// GetCustomer implements port.CustomerUsecase.
func (c *customerUsecase) GetCustomer(ctx context.Context, in any, opts any) (any, error) {
	log.SetPrefix("port.CustomerUsecase.GetCustomer\t")

	getCustomerIn, ok := in.(*model.GetCustomerIn)
	if !ok {
		log.Printf("input is not a valid type")
		return nil, errors.New("input is not a valid type")
	}

	_, ok = opts.([]customeroption.Option)
	if opts != nil && !ok {
		log.Printf("option is not a valid type")
		return nil, errors.New("option is not a valid type")
	}

	if prefix, ok := ctx.Value("prefix").(string); ok {
		log.Printf("prefix is setted, %s", prefix)
		getCustomerIn.ID = fmt.Sprintf("%s-%s", prefix, getCustomerIn.ID)
	}

	log.Printf("return customer, %s", getCustomerIn.ID)
	return &model.GetCustomerOut{
		ID:   getCustomerIn.ID,
		Name: fmt.Sprintf("Dummy: %s", getCustomerIn.ID),
	}, nil
}
