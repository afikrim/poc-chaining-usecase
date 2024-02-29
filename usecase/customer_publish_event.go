package usecase

import (
	"context"
	"log"

	"bitbucket.org/Amartha/poc-chaining-usecase/model"
)

// PublishEvent implements port.CustomerUsecase.
func (*customerUsecase) PublishEvent(h model.UsecaseHandlerFunc) model.UsecaseHandlerFunc {
	return func(ctx context.Context, in interface{}, opts ...interface{}) (interface{}, error) {
		out, err := h(ctx, in, opts...)

		log.SetPrefix("port.CustomerUsecase.PublishEvent\t")
		if err != nil {
			log.Printf("process failed, %s", err.Error())
			return out, err
		}

		log.Printf("event published")
		return out, err
	}
}
