package auth

import (
	"github.com/alexandrhub/auth/internal/service"
	pb "github.com/alexandrhub/auth/pkg/user_v1"
)

type Implementation struct {
	pb.UnimplementedUserV1Server
	authService service.AuthService
}

func NewImplementation(authService service.AuthService) *Implementation {
	return &Implementation{
		authService: authService,
	}
}
