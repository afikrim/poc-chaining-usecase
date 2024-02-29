package customeroption

import "bitbucket.org/Amartha/poc-chaining-usecase/model"

type Option func(o *Options)
type Options struct {
	Middlewares []model.UsecaseMiddlewareFunc
}

func NewOptions(opts ...Option) *Options {
	o := &Options{}
	for _, opt := range opts {
		opt(o)
	}

	return o
}

func WithMiddleware(m model.UsecaseMiddlewareFunc) Option {
	return func(o *Options) {
		if o.Middlewares == nil {
			o.Middlewares = make([]model.UsecaseMiddlewareFunc, 0)
		}

		o.Middlewares = append(o.Middlewares, m)
	}
}
