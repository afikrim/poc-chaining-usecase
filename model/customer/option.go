package customeroption

type Option func(o *Options)
type Options struct {
}

func NewOptions(opts ...Option) *Options {
	o := &Options{}
	for _, opt := range opts {
		opt(o)
	}

	return o
}
