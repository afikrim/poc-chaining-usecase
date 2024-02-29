package main

import (
	"context"
	"fmt"

	"bitbucket.org/Amartha/poc-chaining-usecase/model"
	customeroption "bitbucket.org/Amartha/poc-chaining-usecase/model/customer"
	"bitbucket.org/Amartha/poc-chaining-usecase/usecase"
)

func main() {
	usecase := usecase.NewUsecase()

	getCustomerHandler := usecase.GetCustomerUsecase().GetCustomer
	checkXSotMiddleware := usecase.GetCustomerUsecase().CheckXSot
	publishEventMiddleware := usecase.GetCustomerUsecase().PublishEvent

	// directly call get customer without any middleware
	out, err := getCustomerHandler(context.Background(), &model.GetCustomerIn{ID: "123091283"})
	fmt.Println(err)
	fmt.Println(out)

	// use check x sot middleware
	ctx := context.WithValue(context.Background(), "x-sot", "customer")
	out, err = getCustomerHandler(ctx, &model.GetCustomerIn{ID: "123091283"}, customeroption.WithMiddleware(checkXSotMiddleware))
	fmt.Println(err)
	fmt.Println(out)

	// use check x sot middleware and publish event middleware
	out, err = getCustomerHandler(context.Background(), &model.GetCustomerIn{ID: "123091283"}, customeroption.WithMiddleware(checkXSotMiddleware), customeroption.WithMiddleware(publishEventMiddleware))
	fmt.Println(err)
	fmt.Println(out)
}
