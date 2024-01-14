package auth

import (
	"context"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/alexandrhub/auth/internal/converter"
	pb "github.com/alexandrhub/auth/pkg/user_v1"
)

func (i *Implementation) Create(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	id, err := i.authService.Create(ctx, converter.ToUserFromDescCreate(req.Info))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	log.Println("User created: ", id)

	return &pb.CreateUserResponse{
		Id: id,
	}, nil
}
