package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"

	"github.com/alexandrhub/auth/internal/model"
	"github.com/alexandrhub/auth/internal/repository"
	repoMocks "github.com/alexandrhub/auth/internal/repository/mocks"
	"github.com/alexandrhub/auth/internal/service/auth"
)

func TestUpdate(t *testing.T) {
	t.Parallel()

	type authRepositoryMockFunc func(mc *minimock.Controller) repository.AuthRepository

	type args struct {
		ctx context.Context
		req *model.UserUpdate
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id    = gofakeit.Int64()
		name  = gofakeit.Animal()
		email = gofakeit.Email()

		repoErr = fmt.Errorf("repository error")

		req = &model.UserUpdate{
			ID:    id,
			Name:  name,
			Email: email,
			Role:  0,
		}

		res any
	)

	tests := []struct {
		name               string
		args               args
		want               any
		err                error
		authRepositoryMock authRepositoryMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: res,
			err:  nil,
			authRepositoryMock: func(mc *minimock.Controller) repository.AuthRepository {
				mock := repoMocks.NewAuthRepositoryMock(mc)
				mock.UpdateMock.Expect(ctx, req).Return(nil)
				return mock
			},
		},
		{
			name: "repository error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  repoErr,
			authRepositoryMock: func(mc *minimock.Controller) repository.AuthRepository {
				mock := repoMocks.NewAuthRepositoryMock(mc)
				mock.UpdateMock.Expect(ctx, req).Return(repoErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(
			tt.name, func(t *testing.T) {
				t.Parallel()

				authRepositoryMock := tt.authRepositoryMock(mc)
				service := auth.NewMockService(authRepositoryMock)
				err := service.Update(tt.args.ctx, tt.args.req)
				require.Equal(t, err, tt.err)
				require.Equal(t, tt.want, res)
			},
		)
	}
}
