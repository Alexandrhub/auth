package auth

import (
	"context"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/alexandrhub/auth/internal/converter"
	desc "github.com/alexandrhub/auth/pkg/user_v1"
)

func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	id, err := i.authService.Create(ctx, converter.ToUserFromDescCreate(req.GetInfo()))
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to create user: %v", err)
	}

	log.Printf("inserted auth with id: %d", id)

	return &desc.CreateResponse{
		Id: id,
	}, nil
}
