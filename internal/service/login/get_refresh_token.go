package login

import (
	"context"
	"os"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/alexandrhub/auth/internal/model"
	"github.com/alexandrhub/auth/internal/utils"
)

func (s *serverAuth) GetRefreshToken(ctx context.Context, token string) (string, error) {
	refreshTokenSecretKey := os.Getenv("REFRESH_TOKEN_SECRET_KEY")

	claims, err := utils.VerifyToken(token, []byte(refreshTokenSecretKey))
	if err != nil {
		return "", status.Errorf(codes.Aborted, "invalid refresh token")
	}

	r, err := s.loginRepository.GetUserRole(ctx)
	if err != nil {
		return "", err
	}

	refreshToken, err := utils.GenerateToken(
		model.UserInfo{
			Username: claims.Username,
			Role:     r,
		},
		[]byte(refreshTokenSecretKey),
		refreshTokenExpiration,
	)
	if err != nil {
		return "", err
	}

	return refreshToken, nil
}
