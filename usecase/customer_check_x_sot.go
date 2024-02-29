package usecase

import (
	"context"
	"log"

	"bitbucket.org/Amartha/poc-chaining-usecase/model"
)

// CheckXSot implements port.CustomerUsecase.
func (*customerUsecase) CheckXSot(h model.UsecaseHandlerFunc) model.UsecaseHandlerFunc {
	return func(ctx context.Context, in interface{}, opts ...interface{}) (interface{}, error) {
		log.SetPrefix("port.CustomerUsecase.CheckXSot\t\t")

		log.Print("started")
		sot, ok := ctx.Value("x-sot").(string)
		if ok {
			log.Printf("x-sot is not empty, %s", sot)
			return h(context.WithValue(ctx, "prefix", sot), in, opts...)
		}

		log.Print("x-sot is empty")
		return h(ctx, in, opts...)
	}
}
