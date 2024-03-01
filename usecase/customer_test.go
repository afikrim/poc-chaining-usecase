package usecase

import (
	"context"
	"testing"

	"bitbucket.org/Amartha/poc-chaining-usecase/model"
	"github.com/stretchr/testify/assert"
)

func Test_GetCustomer(t *testing.T) {
	type args struct {
		ctx context.Context
		in  *model.GetCustomerIn
	}

	tests := []struct {
		name string
		args args
		out  *model.GetCustomerOut
		err  error
	}{
		{
			name: "Given a valid get customer input, When getting customer, Should return valid customer",
			args: args{
				ctx: context.Background(),
				in:  &model.GetCustomerIn{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			usecase := NewUsecase()

			out, err := usecase.GetCustomerUsecase().GetCustomer(tt.args.ctx, tt.args.in)
			assert.Equal(t, tt.err, err)
			assert.Equal(t, tt.out, out)
		})
	}
}
