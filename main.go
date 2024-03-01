package main

import (
	"context"
	"fmt"

	"bitbucket.org/Amartha/poc-chaining-usecase/middleware"
	"bitbucket.org/Amartha/poc-chaining-usecase/model"
	"bitbucket.org/Amartha/poc-chaining-usecase/usecase"
)

func main() {
	uc := usecase.NewUsecase()

	// directly call get customer without any middleware
	m := middleware.ChainMiddleware[*model.GetCustomerIn, *model.GetCustomerOut]()
	out, err := usecase.Customer_GetCustomer_Handler(uc, context.Background(), &model.GetCustomerIn{ID: "123091283"}, m)
	fmt.Println(err)
	fmt.Println(out)

	// use check x sot middleware
	ctx := context.WithValue(context.Background(), "x-sot", "customer")
	m = middleware.ChainMiddleware[*model.GetCustomerIn, *model.GetCustomerOut](
		middleware.CheckXSot[*model.GetCustomerIn, *model.GetCustomerOut](),
	)
	out, err = usecase.Customer_GetCustomer_Handler(uc, ctx, &model.GetCustomerIn{ID: "123091283"}, m)
	fmt.Println(err)
	fmt.Println(out)

	// use check x sot middleware and publish event middleware
	m = middleware.ChainMiddleware[*model.GetCustomerIn, *model.GetCustomerOut](
		middleware.CheckXSot[*model.GetCustomerIn, *model.GetCustomerOut](),
		middleware.PublishEvent[*model.GetCustomerIn, *model.GetCustomerOut](),
	)
	out, err = usecase.Customer_GetCustomer_Handler(uc, context.Background(), &model.GetCustomerIn{ID: "123091283"}, m)
	fmt.Println(err)
	fmt.Println(out)
}
