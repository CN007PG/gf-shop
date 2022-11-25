// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"gf-shop/internal/model"
	"gf-shop/internal/model/entity"
)

type (
	IUser interface {
		Create(ctx context.Context, in model.UserCreateInput) (err error)
		ChangePassword(ctx context.Context, name string, password string) (err error)
		Update(ctx context.Context, in model.UserUpdateInput) (err error)
		Get(ctx context.Context, name string) (data *entity.User, err error)
	}
)

var (
	localUser IUser
)

func User() IUser {
	if localUser == nil {
		panic("implement not found for interface IUser, forgot register?")
	}
	return localUser
}

func RegisterUser(i IUser) {
	localUser = i
}
