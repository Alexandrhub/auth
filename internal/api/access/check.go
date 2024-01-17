package access

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
)

func (i *Implementation) Check(ctx context.Context, endpointAddress string) (*empty.Empty, error) {
	err := i.accessService.Check(ctx, endpointAddress)
	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}
