package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// 商城 get 主页路由
type ShopGetIndexReq struct {
	g.Meta `path:"/" tags:"shop" method:"get" summary:"shop get index"`
}

// 商城 get 主页返回
type ShopGetIndexRes struct {
	g.Meta `mime:"text/html" example:"string"`
}

// 商城 get 登录路由
type ShopGetLoginReq struct {
	g.Meta `path:"/login" tags:"shop" method:"get" summary:"shop get login"`
}

// 商城 get 登录返回
type ShopGetLoginRes struct {
	g.Meta `mime:"text/html" example:"string"`
}

// 商城 get 修改密码路由
type ShopGetChangepasswordReq struct {
	g.Meta `path:"/changepassword" tags:"shop" method:"get" summary:"shop get change password"`
}

// 商城 get 修改密码返回
type ShopGetChangepasswordRes struct {
	g.Meta `mime:"text/html" example:"string"`
}

// 商城 get 注册路由
type ShopGetRegisterReq struct {
	g.Meta `path:"/register" tags:"shop" method:"get" summary:"shop get register"`
}

// 商城 get 注册返回
type ShopGetRegisterRes struct {
	g.Meta `mime:"text/html" example:"string"`
}

// 商城 post 主页路由
type ShopPostIndexReq struct {
	g.Meta `path:"/" tags:"shop" method:"post" summary:"shop post index"`
	Token  string
}

// 商城 post 主页返回
type ShopPostIndexRes struct {
	g.Meta `mime:"text/html" example:"string"`
}

// 商城 post 登录路由
type ShopPostLoginReq struct {
	g.Meta   `path:"/login" tags:"shop" method:"post" summary:"shop post login"`
	Password string `v:"required"`
	Nickname string `v:"required"`
}

// 商城 post 登录返回
type ShopPostLoginRes struct {
	g.Meta `mime:"text/html" example:"string"`
}

// 商城 post 登出路由
type ShopPostLogoutReq struct {
	g.Meta `path:"/logout" tags:"shop" method:"post" summary:"shop post login"`
}

// 商城 post 登出返回
type ShopPostLogoutRes struct {
	g.Meta `mime:"text/html" example:"string"`
}

// 商城 post 修改密码路由
type ShopPostChangePasswordReq struct {
	g.Meta      `path:"/changepassword" tags:"shop" method:"post" summary:"shop post change password"`
	PasswordOld string `v:"required"`
	Password    string `v:"required"`
	Nickname    string `v:"required"`
}

// 商城 post 修改密码返回
type ShopPostChangePasswordRes struct {
	g.Meta `mime:"text/html" example:"string"`
}

// 商城 post 注册路由
type ShopPostRegisterReq struct {
	g.Meta   `path:"/register" tags:"shop" method:"post" summary:"shop post register"`
	Phone    string `v:"required"`
	Password string `v:"required"`
	Nickname string `v:"required"`
}

// 商城 post 注册返回
type ShopPostRegisterRes struct {
	g.Meta `mime:"text/html" example:"string"`
}
