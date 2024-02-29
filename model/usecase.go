package model

import "context"

type UsecaseHandlerFunc func(ctx context.Context, in any, opts any) (out any, err error)
type UsecaseMiddlewareFunc func(h UsecaseHandlerFunc) UsecaseHandlerFunc
