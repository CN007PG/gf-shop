// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure of table user for DAO operations like Where/Data.
type User struct {
	g.Meta     `orm:"table:user, do:true"`
	Id         interface{} //
	Name       interface{} // 用户名
	Password   interface{} // 密码
	Phone      interface{} // 电话
	BuyHistory interface{} // 购买记录
	BuyCart    interface{} // 购物车
	Token      interface{} // 用户唯一身份token
	CreateAt   *gtime.Time // 创建时间
	UpdateAt   *gtime.Time // 更新时间
}
