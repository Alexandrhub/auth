package login

import (
	"github.com/alexandrhub/auth/internal/service"
	pb "github.com/alexandrhub/auth/pkg/auth_v1"
)

type Implementation struct {
	pb.UnimplementedAuthV1Server
	loginService service.LoginService
}

func NewImplementation(loginService service.LoginService) *Implementation {
	return &Implementation{loginService: loginService}
}
