package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"

	"github.com/alexandrhub/auth/internal/api/auth"
	"github.com/alexandrhub/auth/internal/service"
	authServiceMocks "github.com/alexandrhub/auth/internal/service/mocks"
	pb "github.com/alexandrhub/auth/pkg/user_v1"
)

func TestDelete(t *testing.T) {
	t.Parallel()

	type authServiceMockFunc func(mc *minimock.Controller) service.AuthService

	type args struct {
		ctx context.Context
		req *pb.DeleteRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id = gofakeit.Int64()

		serviceErr = fmt.Errorf("service error")

		req = &pb.DeleteRequest{
			Id: id,
		}

		res any
	)

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
			want: res,
			err:  nil,
			authServiceMock: func(mc *minimock.Controller) service.AuthService {
				mock := authServiceMocks.NewAuthServiceMock(mc)
				mock.DeleteMock.Expect(ctx, id).Return(nil)
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
				mock.DeleteMock.Expect(ctx, id).Return(serviceErr)
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

				_, err := api.Delete(tt.args.ctx, tt.args.req)
				require.Equal(t, tt.err, err)
				require.Equal(t, tt.want, res)
			},
		)
	}

}
