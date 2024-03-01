package middleware

import (
	"context"
	"log"

	"bitbucket.org/Amartha/poc-chaining-usecase/model"
)

func PublishEvent[Input, Output any]() model.UsecaseMiddlewareFunc[Input, Output] {
	return func(ctx context.Context, in Input, info *model.UsecaseInfo, handler model.UsecaseHandlerFunc[Input, Output]) (Output, error) {
		out, err := handler(ctx, in)

		log.SetPrefix("port.CustomerUsecase.PublishEvent\t")
		if err != nil {
			log.Printf("process failed, %s", err.Error())
			return out, err
		}

		log.Printf("event published")
		return out, err
	}
}
