package auth

import (
	"context"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/alexandrhub/auth/internal/converter"
	pb "github.com/alexandrhub/auth/pkg/user_v1"
)

func (i *Implementation) Get(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	user, err := i.authService.Get(ctx, req.GetId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	log.Println("User found: ", user.UserCreate.UserUpdate.ID)

	return &pb.GetUserResponse{
		Info: converter.ToUserFromService(user),
	}, nil
}
