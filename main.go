package main

import (
	"context"
	"fmt"

	"bitbucket.org/Amartha/poc-chaining-usecase/model"
	"bitbucket.org/Amartha/poc-chaining-usecase/usecase"
)

func main() {
	usecase := usecase.NewUsecase()

	getCustomerHandler := usecase.GetCustomerUsecase().GetCustomer
	checkXSotMiddleware := usecase.GetCustomerUsecase().CheckXSot
	publishEventMiddleware := usecase.GetCustomerUsecase().PublishEvent

	// directly call get customer without any middleware
	h := usecase.ChainHandler(getCustomerHandler)
	out, err := h(context.Background(), &model.GetCustomerIn{ID: "123091283"}, nil)
	fmt.Println(err)
	fmt.Println(out)

	// use check x sot middleware
	ctx := context.WithValue(context.Background(), "x-sot", "customer")
	h = usecase.ChainHandler(getCustomerHandler, checkXSotMiddleware)
	out, err = h(ctx, &model.GetCustomerIn{ID: "123091283"}, nil)
	fmt.Println(err)
	fmt.Println(out)

	// use check x sot middleware and publish event middleware
	h = usecase.ChainHandler(getCustomerHandler, checkXSotMiddleware, publishEventMiddleware)
	out, err = h(ctx, &model.GetCustomerIn{ID: "123091283"}, nil)
	fmt.Println(err)
	fmt.Println(out)

	// use invalid input
	h = usecase.ChainHandler(getCustomerHandler, checkXSotMiddleware)
	out, err = h(ctx, struct{}{}, nil)
	fmt.Println(err)
	fmt.Println(out)
}
