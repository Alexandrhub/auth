package login

import (
	"context"

	"github.com/alexandrhub/auth/internal/converter"
	pb "github.com/alexandrhub/auth/pkg/auth_v1"
)

func (i *Implementation) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	rToken, err := i.loginService.Login(ctx, converter.ToUserClaimsFromLogin(req.GetLogin()))
	if err != nil {
		return nil, err
	}

	return &pb.LoginResponse{
		RefreshToken: rToken,
	}, nil
}
