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
	IItemCat interface {
		Create(ctx context.Context, in model.ItemCatCreateInput) (err error)
		Delete(ctx context.Context, id int64) (err error)
		Get(ctx context.Context, id int64) (data *entity.ItemCat, err error)
	}
)

var (
	localItemCat IItemCat
)

func ItemCat() IItemCat {
	if localItemCat == nil {
		panic("implement not found for interface IItemCat, forgot register?")
	}
	return localItemCat
}

func RegisterItemCat(i IItemCat) {
	localItemCat = i
}