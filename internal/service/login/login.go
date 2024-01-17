package login

import (
	"context"
	"os"
	"time"

	"github.com/pkg/errors"

	"github.com/alexandrhub/auth/internal/model"
	"github.com/alexandrhub/auth/internal/utils"
)

const (
	refreshTokenExpiration = 60 * time.Minute
	accessTokenExpiration  = 15 * time.Minute
)

func (s *serverAuth) Login(ctx context.Context, info *model.UserClaims) (string, error) {
	refreshTokenSecretKey := os.Getenv("REFRESH_TOKEN_SECRET_KEY")
	r, err := s.loginRepository.GetUserRole(ctx)
	if err != nil {
		return "", err
	}
	refreshToken, err := utils.GenerateToken(
		model.UserInfo{
			Username: info.Username,
			Role:     r,
		}, []byte(refreshTokenSecretKey), refreshTokenExpiration,
	)
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return refreshToken, nil
}
