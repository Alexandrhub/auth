package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"

	"github.com/alexandrhub/auth/internal/api/auth"
	"github.com/alexandrhub/auth/internal/model"
	"github.com/alexandrhub/auth/internal/service"
	authServiceMocks "github.com/alexandrhub/auth/internal/service/mocks"
	pb "github.com/alexandrhub/auth/pkg/user_v1"
)

func TestUpdate(t *testing.T) {
	t.Parallel()

	type authServiceMockFunc func(mc *minimock.Controller) service.AuthService

	type args struct {
		ctx context.Context
		req *pb.UpdateRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id    = gofakeit.Int64()
		name  = gofakeit.Animal()
		email = gofakeit.Email()

		serviceErr = fmt.Errorf("service error")

		req = &pb.UpdateRequest{
			Info: &pb.UserUpdate{
				Id:    id,
				Name:  name,
				Email: email,
				Role:  0,
			},
		}

		updateRes any

		info = &model.UserUpdate{
			ID:    id,
			Name:  name,
			Email: email,
			Role:  0,
		}
	)
	defer t.Cleanup(mc.Finish)

	tests := []struct {
		name            string
		args            args
		want            any
		err             error
		authServiceMock authServiceMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: updateRes,
			err:  nil,
			authServiceMock: func(mc *minimock.Controller) service.AuthService {
				mock := authServiceMocks.NewAuthServiceMock(mc)
				mock.UpdateMock.Expect(ctx, info).Return(nil)
				return mock
			},
		},
		{
			name: "service error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  serviceErr,
			authServiceMock: func(mc *minimock.Controller) service.AuthService {
				mock := authServiceMocks.NewAuthServiceMock(mc)
				mock.UpdateMock.Expect(ctx, info).Return(serviceErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(
			tt.name, func(t *testing.T) {
				t.Parallel()

				authServiceMock := tt.authServiceMock(mc)
				api := auth.NewImplementation(authServiceMock)

				_, err := api.Update(tt.args.ctx, tt.args.req)
				require.Equal(t, tt.err, err)
				require.Equal(t, tt.want, updateRes)
			},
		)
	}

}
