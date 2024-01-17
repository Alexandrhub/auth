package access

import (
	"context"
	"strings"

	"github.com/pkg/errors"
	"google.golang.org/grpc/metadata"

	"github.com/alexandrhub/auth/internal/utils"
)

// move to .env
const (
	grpcPort   = 50051
	authPrefix = "Bearer"

	refreshTokenSecretKey = "refresh_token_secret_key"
	accessTokenSecretKey  = "access_token_secret_key"
)

var accessibleRoles map[string]string

func (s *serverAccess) Check(ctx context.Context, endpointAddress string) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return errors.New("no metadata from incoming context")
	}

	authHeader, ok := md["authorization"]
	if !ok || len(authHeader) == 0 {
		return errors.New("no authorization header")
	}

	if !strings.HasPrefix(authHeader[0], authPrefix) {
		return errors.New("invalid authorization header")
	}

	accessToken := strings.TrimPrefix(authHeader[0], authPrefix)

	claims, err := utils.VerifyToken(accessToken, []byte(accessTokenSecretKey))
	if err != nil {
		return errors.New("invalid access token")
	}

	accessibleMap, err := s.accessibleRoles(ctx)
	if err != nil {
		return errors.New("failed to get accessible roles")
	}

	role, ok := accessibleMap[endpointAddress]
	if !ok {
		return errors.New("endpoint is not accessible")
	}

	if role == claims.Role {
		return nil
	}

	return errors.New("access denied")
}

// Возвращает мапу с адресом эндпоинта и ролью, которая имеет доступ к нему
func (s *serverAccess) accessibleRoles(ctx context.Context) (map[string]string, error) {
	if accessibleRoles == nil {
		Roles, err := s.accessRepository.Roles(ctx)
		if err != nil {
			return nil, nil
		}
		accessibleRoles = Roles
	}

	return accessibleRoles, nil
}
