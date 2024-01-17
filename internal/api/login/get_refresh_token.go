package login

import (
	"context"

	pb "github.com/alexandrhub/auth/pkg/auth_v1"
)

func (i *Implementation) GetRefreshToken(ctx context.Context, req *pb.GetRefreshTokenRequest) (*pb.GetRefreshTokenResponse, error) {
	rToken, err := i.loginService.GetRefreshToken(ctx, req.GetRefreshToken())
	if err != nil {
		return nil, err
	}

	return &pb.GetRefreshTokenResponse{
		RefreshToken: rToken,
	}, nil
}
