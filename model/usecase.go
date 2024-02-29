package model

import "context"

type UsecaseHandlerFunc func(ctx context.Context, in interface{}, opts ...interface{}) (out interface{}, err error)
type UsecaseMiddlewareFunc func(h UsecaseHandlerFunc) UsecaseHandlerFunc
