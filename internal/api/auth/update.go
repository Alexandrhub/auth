package auth

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/alexandrhub/auth/internal/converter"
	desc "github.com/alexandrhub/auth/pkg/user_v1"
)

func (i *Implementation) Update(ctx context.Context, req *desc.UpdateRequest) (*empty.Empty, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}
	err = i.authService.Update(ctx, converter.ToUserFromDescUpdate(req.GetInfo()))
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
