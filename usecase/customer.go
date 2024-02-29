package usecase

import (
	"context"
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
func (c *customerUsecase) GetCustomer(ctx context.Context, in *model.GetCustomerIn, opts ...customeroption.Option) (*model.GetCustomerOut, error) {
	o := customeroption.NewOptions(opts...)

	// use grpc approach when passing a handler to the chaining method
	f := func(_ context.Context, _ interface{}, _ ...interface{}) (interface{}, error) {
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

	// chaining handler
	h := c.ucs.ChainHandler(f, o.Middlewares...)

	out, err := h(ctx, in, opts)
	if err != nil {

	}

	return out.(*model.GetCustomerOut), err
}
