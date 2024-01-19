package auth

import (
	"context"

	"google.golang.org/grpc/codes"

	"github.com/alexandrhub/auth/internal/converter"
	"github.com/alexandrhub/auth/internal/pkg/sys"
	"github.com/alexandrhub/auth/internal/pkg/sys/validation"
	desc "github.com/alexandrhub/auth/pkg/user_v1"
)

func (i *Implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	err := validation.Validate(
		ctx,
		validateID(req.GetId()),
		otherValidateID(req.GetId()),
	)
	if err != nil {
		return nil, err
	}

	if req.GetId() > 100 {
		return nil, sys.NewCommonError("id must be less then 100", codes.ResourceExhausted)
	}

	authObj, err := i.authService.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	// log.Printf("id: %d", authObj.UserCreate.UserUpdate.ID)

	return &desc.GetResponse{
		Info: converter.ToUserFromService(authObj),
	}, nil
}

func validateID(id int64) validation.Condition {
	return func(ctx context.Context) error {
		if id <= 0 {
			return validation.NewValidationErrors("id must be greater than 0")
		}

		return nil
	}
}

func otherValidateID(id int64) validation.Condition {
	return func(ctx context.Context) error {
		if id <= 100 {
			return validation.NewValidationErrors("id must be greater than 100")
		}

		return nil
	}
}
