package access

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"

	pb "github.com/alexandrhub/auth/pkg/access_v1"
)

func (i *Implementation) Check(ctx context.Context, req *pb.CheckRequest) (*empty.Empty, error) {
	err := i.accessService.Check(ctx, req.GetEndpointAddress())
	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}
