package converter

import (
	"github.com/alexandrhub/auth/internal/model"
	pb "github.com/alexandrhub/auth/pkg/auth_v1"
)

func ToUserClaimsFromLogin(req *pb.Login) *model.UserClaims {
	return &model.UserClaims{
		Username: req.GetUsername(),
	}
}
