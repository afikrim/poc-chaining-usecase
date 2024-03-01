package model

import (
	"context"
)

type UsecaseInfo struct {
	Uc any
}

type UsecaseHandlerFunc[Input, Output any] func(ctx context.Context, in Input) (Output, error)
type UsecaseMiddlewareFunc[Input, Output any] func(ctx context.Context, in Input, info *UsecaseInfo, handler UsecaseHandlerFunc[Input, Output]) (Output, error)
