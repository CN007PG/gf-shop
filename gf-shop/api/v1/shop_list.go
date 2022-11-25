package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// 商城列表 get 主页路由
type ShopListIndexReq struct {
	g.Meta `path:"/" tags:"shopList" method:"get" summary:"shopList get index"`
}

// 商城列表 get 主页返回
type ShopListIndexRes struct {
	g.Meta `mime:"text/html" example:"string"`
}
