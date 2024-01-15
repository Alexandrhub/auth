package tests

import (
	"context"
	"database/sql"
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

func TestGet(t *testing.T) {
	t.Parallel()

	type authRepositoryMockFunc func(mc *minimock.Controller) repository.AuthRepository

	type args struct {
		ctx context.Context
		req int64
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

		repoErr = fmt.Errorf("repository error")

		res = &model.User{
			UserCreate: model.UserCreate{
				UserUpdate: model.UserUpdate{ID: id, Name: name, Email: email, Role: 0},
				Password:   password,
			},
			CreatedAt: createdAt,
			UpdatedAt: sql.NullTime{Time: updatedAt, Valid: true},
		}
	)

	tests := []struct {
		name               string
		args               args
		want               *model.User
		err                error
		authRepositoryMock authRepositoryMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: id,
			},
			want: res,
			err:  nil,
			authRepositoryMock: func(mc *minimock.Controller) repository.AuthRepository {
				mock := repoMocks.NewAuthRepositoryMock(mc)
				mock.GetMock.Expect(ctx, id).Return(res, nil)
				return mock
			},
		},
		{
			name: "repository error case",
			args: args{
				ctx: ctx,
				req: id,
			},
			want: nil,
			err:  repoErr,
			authRepositoryMock: func(mc *minimock.Controller) repository.AuthRepository {
				mock := repoMocks.NewAuthRepositoryMock(mc)
				mock.GetMock.Expect(ctx, id).Return(nil, repoErr)
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
				resID, err := service.Get(tt.args.ctx, tt.args.req)
				require.Equal(t, err, tt.err)
				require.Equal(t, tt.want, resID)
			},
		)
	}
}
