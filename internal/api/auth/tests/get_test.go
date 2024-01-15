package tests

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/alexandrhub/auth/internal/api/auth"
	"github.com/alexandrhub/auth/internal/model"
	"github.com/alexandrhub/auth/internal/service"
	authServiceMocks "github.com/alexandrhub/auth/internal/service/mocks"
	pb "github.com/alexandrhub/auth/pkg/user_v1"
)

func TestGet(t *testing.T) {
	t.Parallel()

	type authServiceMockFunc func(mc *minimock.Controller) service.AuthService

	type args struct {
		ctx context.Context
		req *pb.GetRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id        = gofakeit.Int64()
		name      = gofakeit.Animal()
		email     = gofakeit.Email()
		password  = gofakeit.HackerPhrase()
		createdAt = gofakeit.Date()
		updatedAt = gofakeit.Date()

		serviceErr = fmt.Errorf("service error")

		req = &pb.GetRequest{
			Id: id,
		}

		getRes = &model.User{
			UserCreate: model.UserCreate{
				UserUpdate: model.UserUpdate{
					ID:    id,
					Name:  name,
					Email: email,
					Role:  0,
				},
				Password: password,
			},
			CreatedAt: createdAt,
			UpdatedAt: sql.NullTime{Time: updatedAt, Valid: true},
		}

		res = &pb.GetResponse{
			Info: &pb.User{
				UserCreate: &pb.UserCreate{
					UserUpdate: &pb.UserUpdate{
						Id:    id,
						Name:  name,
						Email: email,
						Role:  0,
					},
					Password: password,
				},
				CreatedAt: timestamppb.New(createdAt),
				UpdatedAt: timestamppb.New(updatedAt),
			},
		}
	)

	defer t.Cleanup(mc.Finish)

	tests := []struct {
		name            string
		args            args
		want            *pb.GetResponse
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
				mock.GetMock.Expect(ctx, id).Return(getRes, nil)
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
				mock.GetMock.Expect(ctx, id).Return(nil, serviceErr)
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

				newID, err := api.Get(tt.args.ctx, tt.args.req)
				require.Equal(t, tt.err, err)
				require.Equal(t, tt.want, newID)
			},
		)
	}

}
