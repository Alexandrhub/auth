package login

import (
	"context"

	pb "github.com/alexandrhub/auth/pkg/auth_v1"
)

func (i *Implementation) GetAccessToken(ctx context.Context, req *pb.GetAccessTokenRequest) (*pb.GetAccessTokenResponse, error) {
	accessT, err := i.loginService.GetAccessToken(ctx, req.GetRefreshToken())
	if err != nil {
		return nil, err
	}

	return &pb.GetAccessTokenResponse{
		AccessToken: accessT,
	}, nil
}
