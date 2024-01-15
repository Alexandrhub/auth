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

func TestCreate(t *testing.T) {
	t.Parallel()

	type authServiceMockFunc func(mc *minimock.Controller) service.AuthService

	type args struct {
		ctx context.Context
		req *pb.CreateRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id       = gofakeit.Int64()
		title    = gofakeit.Animal()
		email    = gofakeit.Email()
		password = gofakeit.HackerPhrase()

		serviceErr = fmt.Errorf("service error")

		req = &pb.CreateRequest{
			Info: &pb.UserCreate{
				UserUpdate: &pb.UserUpdate{
					Id:    id,
					Name:  title,
					Email: email,
					Role:  0,
				},
				Password: password,
			},
		}

		info = &model.UserCreate{
			UserUpdate: model.UserUpdate{
				ID:    id,
				Name:  title,
				Email: email,
				Role:  0,
			},
			Password: password,
		}

		res = &pb.CreateResponse{
			Id: id,
		}
	)

	tests := []struct {
		name            string
		args            args
		want            *pb.CreateResponse
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
				mock.CreateMock.Expect(ctx, info).Return(id, nil)
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
				mock.CreateMock.Expect(ctx, info).Return(id, serviceErr)
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

				newID, err := api.Create(tt.args.ctx, tt.args.req)
				require.Equal(t, tt.err, err)
				require.Equal(t, tt.want, newID)
			},
		)
	}

}
