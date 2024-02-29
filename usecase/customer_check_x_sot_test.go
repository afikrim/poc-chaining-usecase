package usecase

import (
	"context"
	"testing"

	"bitbucket.org/Amartha/poc-chaining-usecase/model"
	"github.com/stretchr/testify/assert"
)

func Test_CheckXSot(t *testing.T) {
	type args struct {
		ctx  context.Context
		in   any
		opts any
	}

	tests := []struct {
		name string
		h    model.UsecaseHandlerFunc
		args args
		out  any
		err  error
	}{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			usecase := NewUsecase()

			handler := usecase.GetCustomerUsecase().CheckXSot(tt.h)
			out, err := handler(tt.args.ctx, tt.args.in, tt.args.opts)

			assert.Equal(t, tt.out, out)
			assert.Equal(t, tt.err, err)
		})
	}
}
