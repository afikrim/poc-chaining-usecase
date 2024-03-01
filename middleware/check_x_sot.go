package middleware

import (
	"context"
	"log"

	"bitbucket.org/Amartha/poc-chaining-usecase/model"
)

func CheckXSot[Input, Output any]() model.UsecaseMiddlewareFunc[Input, Output] {
	return func(ctx context.Context, in Input, info *model.UsecaseInfo, handler model.UsecaseHandlerFunc[Input, Output]) (Output, error) {
		log.SetPrefix("port.CustomerUsecase.CheckXSot\t\t")

		log.Print("started")
		sot, ok := ctx.Value("x-sot").(string)
		if ok {
			log.Printf("x-sot is not empty, %s", sot)
			return handler(context.WithValue(ctx, "prefix", sot), in)
		}

		log.Print("x-sot is empty")
		return handler(ctx, in)
	}
}
