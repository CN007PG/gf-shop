package cmd

import (
	"context"
	"gf-shop/internal/controller"
	"gf-shop/internal/service"
	"gf-shop/utility"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "shop",
		Usage: "shop",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			//
			//商城手机分类服务
			utility.BindStaticServer("")
			//商城购买
			utility.BindStaticServer("serverBuy")

			//商城服务
			utility.BindDynamicServer("serverShop", controller.Shop, service.Middleware().Ctx)
			//商城列表
			utility.BindDynamicServer("serverList", controller.ShopList, service.Middleware().Auth)
			//商城结算
			utility.BindDynamicServer("serverAccount", controller.ShopAccount, service.Middleware().Auth)
			//
			g.Wait()
			return nil
		},
	}
)
