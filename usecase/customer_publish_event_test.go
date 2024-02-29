package usecase

import (
	"context"
	"testing"

	"bitbucket.org/Amartha/poc-chaining-usecase/model"
	"github.com/stretchr/testify/assert"
)

func Test_PublishEvent(t *testing.T) {
	type args struct {
		ctx context.Context
		in  interface{}
	}

	tests := []struct {
		name string
		h    model.UsecaseHandlerFunc
		args args
		out  interface{}
		err  error
	}{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			usecase := NewUsecase()

			handler := usecase.GetCustomerUsecase().PublishEvent(tt.h)
			out, err := handler(tt.args.ctx, tt.args.in)

			assert.Equal(t, tt.out, out)
			assert.Equal(t, tt.err, err)
		})
	}
}
