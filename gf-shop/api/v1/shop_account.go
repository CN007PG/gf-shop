package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// 商城结算 get 主页路由
type ShopAccountIndexReq struct {
	g.Meta `path:"/" tags:"shopAccount" method:"get" summary:"shopAccount get index"`
}

// 商城结算 get 主页返回
type ShopAccountIndexRes struct {
	g.Meta `mime:"text/html" example:"string"`
}
