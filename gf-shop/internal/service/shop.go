// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"gf-shop/internal/model"
)

type (
	IShop interface {
		Create(ctx context.Context, in model.UserCreateInput) (err error)
		ChangePassword(ctx context.Context, in model.UserChangePasswordInput) (err error)
		GetSignedInContextUser(ctx context.Context) (user *model.ContextUser)
		SignIn(ctx context.Context, in model.UserSignInInput) (err error)
		SignOut(ctx context.Context) (err error)
	}
)

var (
	localShop IShop
)

func Shop() IShop {
	if localShop == nil {
		panic("implement not found for interface IShop, forgot register?")
	}
	return localShop
}

func RegisterShop(i IShop) {
	localShop = i
}
