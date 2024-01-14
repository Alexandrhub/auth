package auth

import (
	"context"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/alexandrhub/auth/internal/converter"
	desc "github.com/alexandrhub/auth/pkg/user_v1"
)

func (i *Implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	authObj, err := i.authService.Get(ctx, req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to get user: %v", err)
	}

	log.Printf("id: %d", authObj.UserCreate.UserUpdate.ID)

	return &desc.GetResponse{
		Info: converter.ToUserFromService(authObj),
	}, nil
}
