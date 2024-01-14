package auth

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/alexandrhub/auth/internal/converter"
	pb "github.com/alexandrhub/auth/pkg/user_v1"
)

func (i *Implementation) Update(ctx context.Context, req *pb.UpdateUserRequest) (*empty.Empty, error) {
	err := i.authService.Update(ctx, converter.ToUserFromDescUpdate(req.GetInfo()))
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &emptypb.Empty{}, nil
}
