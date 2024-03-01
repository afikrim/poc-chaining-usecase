package middleware

import (
	"context"

	"bitbucket.org/Amartha/poc-chaining-usecase/model"
)

func ChainMiddleware[Input, Output any](middlewares ...model.UsecaseMiddlewareFunc[Input, Output]) model.UsecaseMiddlewareFunc[Input, Output] {
	n := len(middlewares)

	if n == 0 {
		return func(ctx context.Context, in Input, info *model.UsecaseInfo, handler model.UsecaseHandlerFunc[Input, Output]) (Output, error) {
			return handler(ctx, in)
		}
	}

	if n == 1 {
		return middlewares[0]
	}

	return func(ctx context.Context, in Input, info *model.UsecaseInfo, handler model.UsecaseHandlerFunc[Input, Output]) (Output, error) {
		currHandler := handler

		for i := n - 1; i > 0; i-- {
			innerHandler, i := currHandler, i
			currHandler = func(currCtx context.Context, currIn Input) (Output, error) {
				return middlewares[i](currCtx, currIn, info, innerHandler)
			}
		}

		return middlewares[0](ctx, in, info, currHandler)
	}
}
