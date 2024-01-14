package auth

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/alexandrhub/auth/pkg/user_v1"
)

func (i *Implementation) Delete(ctx context.Context, req *pb.DeleteUserRequest) (*empty.Empty, error) {
	err := i.authService.Delete(ctx, req.GetId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &emptypb.Empty{}, nil
}
