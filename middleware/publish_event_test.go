package middleware

import (
	"context"
	"testing"

	"bitbucket.org/Amartha/poc-chaining-usecase/model"
	"github.com/stretchr/testify/assert"
)

func Test_PublishEvent(t *testing.T) {
	type args[Input, Output any] struct {
		ctx     context.Context
		in      Input
		info    *model.UsecaseInfo
		handler model.UsecaseHandlerFunc[Input, Output]
	}

	type test[Input, Output any] struct {
		name string
		args args[Input, Output]
		out  Output
		err  error
	}

	tests := []test[any, any]{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := PublishEvent[any, any](tt.args.ctx, tt.args.in, tt.args.info, tt.args.handler)

			assert.Equal(t, tt.out, out)
			assert.Equal(t, tt.err, err)
		})
	}
}
