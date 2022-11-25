package controller

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"

	"gf-shop/api/v1"
)

var (
	ShopAccount = cShopAccount{}
)

type cShopAccount struct{}

// 商城结算 get 主页监听
func (c *cShopAccount) ShopAccountGetIndex(ctx context.Context, req *v1.ShopAccountIndexReq) (res *v1.ShopAccountIndexRes, err error) {
	cid := g.RequestFromCtx(ctx).Get("cid")

	g.RequestFromCtx(ctx).Response.WriteTpl("account/cart.html", g.Map{
		"tips": cid.String(),
		"cid":  cid.String(),
		"num":  1,
	})
	return
}
