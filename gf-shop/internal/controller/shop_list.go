package controller

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"

	"gf-shop/api/v1"
)

var (
	ShopList = cShopList{}
)

type cShopList struct{}

// 商城购买 get 主页监听
func (c *cShopList) ShopBuyGetIndex(ctx context.Context, req *v1.ShopListIndexReq) (res *v1.ShopListIndexRes, err error) {
	// 返回主页面
	g.RequestFromCtx(ctx).Response.Writeln("hello list phone!")
	return
}
