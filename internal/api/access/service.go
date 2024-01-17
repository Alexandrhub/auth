package access

import (
	"github.com/alexandrhub/auth/internal/service"
	pb "github.com/alexandrhub/auth/pkg/access_v1"
)

type Implementation struct {
	pb.UnimplementedAccessV1Server
	accessService service.AccessService
}

func NewImplementation(accessService service.AccessService) *Implementation {
	return &Implementation{accessService: accessService}
}
