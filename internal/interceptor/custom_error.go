package interceptor

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/alexandrhub/auth/internal/pkg/sys"
	"github.com/alexandrhub/auth/internal/pkg/sys/grpcCodes"
	"github.com/alexandrhub/auth/internal/pkg/sys/validation"
)

type GRPCStatusInterface interface {
	GRPCStatus() *status.Status
}

func ErrorCodesInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	resp, err := handler(ctx, req)
	if nil == err {
		return resp, nil
	}

	fmt.Printf("err: %v\n", err.Error())

	switch {
	case sys.IsCommonError(err):
		commEr := sys.GetCommonError(err)
		code := toGRPCCode(grpcCodes.Code(commEr.Code()))

		err = status.Error(code, commEr.Error())
	case validation.IsValidationError(err):
		err = status.Error(codes.InvalidArgument, err.Error())
	default:
		var se GRPCStatusInterface
		if errors.As(err, &se) {
			return nil, se.GRPCStatus().Err()
		} else {
			if errors.Is(err, context.DeadlineExceeded) {
				err = status.Error(codes.DeadlineExceeded, err.Error())
			} else if errors.Is(err, context.Canceled) {
				err = status.Error(codes.Canceled, err.Error())
			} else {
				err = status.Error(codes.Internal, "internal error")
			}
		}
	}

	return resp, err
}

func toGRPCCode(code grpcCodes.Code) codes.Code {
	var res codes.Code

	switch code {
	case grpcCodes.OK:
		res = codes.OK
	case grpcCodes.Canceled:
		res = codes.Canceled
	case grpcCodes.InvalidArgument:
		res = codes.InvalidArgument
	case grpcCodes.DeadlineExceeded:
		res = codes.DeadlineExceeded
	case grpcCodes.NotFound:
		res = codes.NotFound
	case grpcCodes.AlreadyExists:
		res = codes.AlreadyExists
	case grpcCodes.PermissionDenied:
		res = codes.PermissionDenied
	case grpcCodes.ResourceExhausted:
		res = codes.ResourceExhausted
	case grpcCodes.FailedPrecondition:
		res = codes.FailedPrecondition
	case grpcCodes.Aborted:
		res = codes.Aborted
	case grpcCodes.OutOfRange:
		res = codes.OutOfRange
	case grpcCodes.Unimplemented:
		res = codes.Unimplemented
	case grpcCodes.Internal:
		res = codes.Internal
	case grpcCodes.Unavailable:
		res = codes.Unavailable
	case grpcCodes.DataLoss:
		res = codes.DataLoss
	case grpcCodes.Unauthenticated:
		res = codes.Unauthenticated
	default:
		res = codes.Unknown
	}

	return res
}
