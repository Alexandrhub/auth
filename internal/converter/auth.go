package converter

import (
	"database/sql"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/alexandrhub/auth/internal/model"
	pb "github.com/alexandrhub/auth/pkg/user_v1"
)

func ToUserFromService(user *model.User) *pb.User {
	var updatedAt *timestamppb.Timestamp
	if user.UpdatedAt.Valid {
		updatedAt = timestamppb.New(user.UpdatedAt.Time)
	}

	return &pb.User{
		UserCreate: ToUserCreateFromService(&user.UserCreate),
		CreatedAt:  timestamppb.New(user.CreatedAt),
		UpdatedAt:  updatedAt,
	}
}

func ToUserCreateFromService(user *model.UserCreate) *pb.UserCreate {
	return &pb.UserCreate{
		UserUpdate: ToUserUpdateFromService(&user.UserUpdate),
		Password:   user.Password,
	}
}

func ToUserUpdateFromService(user *model.UserUpdate) *pb.UserUpdate {
	return &pb.UserUpdate{
		Id:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  pb.Role(user.Role),
	}
}

func ToUserFromDesc(user *pb.User) *model.User {
	var updatedAt sql.NullTime
	if user.UpdatedAt.CheckValid() == nil {
		updatedAt.Valid = true
		updatedAt.Time = user.UpdatedAt.AsTime()
	}

	userCreate := ToUserFromDescCreate(user.GetUserCreate())
	return &model.User{
		UserCreate: *userCreate,
		CreatedAt:  user.CreatedAt.AsTime(),
		UpdatedAt:  updatedAt,
	}
}

func ToUserFromDescCreate(user *pb.UserCreate) *model.UserCreate {
	userUpdate := ToUserFromDescUpdate(user.GetUserUpdate())
	return &model.UserCreate{
		UserUpdate: *userUpdate,
		Password:   user.Password,
	}
}

func ToUserFromDescUpdate(user *pb.UserUpdate) *model.UserUpdate {
	return &model.UserUpdate{
		ID:    user.Id,
		Name:  user.Name,
		Email: user.Email,
		Role:  int(user.GetRole()),
	}
}
