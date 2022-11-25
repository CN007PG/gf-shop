// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure for table user.
type User struct {
	Id         int64       `json:"id"         description:""`
	Name       string      `json:"name"       description:"用户名"`
	Password   string      `json:"password"   description:"密码"`
	Phone      string      `json:"phone"      description:"电话"`
	BuyHistory string      `json:"buyHistory" description:"购买记录"`
	BuyCart    string      `json:"buyCart"    description:"购物车"`
	Token      string      `json:"token"      description:"用户唯一身份token"`
	CreateAt   *gtime.Time `json:"createAt"   description:"创建时间"`
	UpdateAt   *gtime.Time `json:"updateAt"   description:"更新时间"`
}
