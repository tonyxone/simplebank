package grpc

import (
	db "github.com/tonyxone/simplebank/db/sqlc"
	"github.com/tonyxone/simplebank/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertUser(user db.User) *proto.User {
	return &proto.User{
		Username:          user.Username,
		FullName:          user.FullName,
		Email:             user.Email,
		PasswordChangedAt: timestamppb.New(user.PasswordChangedAt),
		CreatedAt:         timestamppb.New(user.CreatedAt),
	}
}
