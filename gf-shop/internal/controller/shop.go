package controller

import (
	"context"
	"gf-shop/internal/model"
	"gf-shop/internal/service"
	"github.com/gogf/gf/v2/frame/g"

	"gf-shop/api/v1"
)

var (
	Shop = cShop{}
)

type cShop struct{}

// 商城 get 主页监听
func (c *cShop) ShopGetIndex(ctx context.Context, req *v1.ShopGetIndexReq) (res *v1.ShopGetIndexRes, err error) {
	// 返回主页面
	user := service.Shop().GetSignedInContextUser(ctx)
	if user != nil {
		// 返回主面
		g.RequestFromCtx(ctx).Response.WriteTpl("shop/index.html", g.Map{
			"tips":           "欢迎您" + user.Nickname,
			"display":        "none",
			"logout_display": "inline-block",
		})
	} else {
		g.RequestFromCtx(ctx).Response.WriteTpl("shop/index.html", g.Map{
			"tips":           "",
			"display":        "inline-block",
			"logout_display": "none",
		})
	}
	return
}

// 商城 get 登录监听
func (c *cShop) ShopGetLogin(ctx context.Context, req *v1.ShopGetLoginReq) (res *v1.ShopGetLoginRes, err error) {
	user := service.Shop().GetSignedInContextUser(ctx)
	if user != nil {
		// 返回主面
		g.RequestFromCtx(ctx).Response.RedirectTo("/")
	} else {
		// 返回登录面
		g.RequestFromCtx(ctx).Response.WriteTpl("shop/login.html")
	}
	return
}

// 商城 get 修改密码监听
func (c *cShop) ShopGetChangepassword(ctx context.Context, req *v1.ShopGetChangepasswordReq) (res *v1.ShopGetChangepasswordRes, err error) {
	// 返回修改密码面
	g.RequestFromCtx(ctx).Response.WriteTpl("shop/changepassword.html")
	return
}

// // 商城 get 注册监听
func (c *cShop) ShopGetRegister(ctx context.Context, req *v1.ShopGetRegisterReq) (res *v1.ShopGetRegisterRes, err error) {
	// 返回注册面
	g.RequestFromCtx(ctx).Response.WriteTpl("shop/register.html")
	return
}

// 商城 post 主页监听
func (c *cShop) ShopPostIndex(ctx context.Context, req *v1.ShopPostIndexReq) (res *v1.ShopPostIndexRes, err error) {
	// 返回主页面
	g.RequestFromCtx(ctx).Response.RedirectTo("/")
	return
}

// 商城 post 登录监听
func (c *cShop) ShopPostLogin(ctx context.Context, req *v1.ShopPostLoginReq) (res *v1.ShopPostLoginRes, err error) {
	err = service.Shop().SignIn(ctx, model.UserSignInInput{
		Password: req.Password,
		Nickname: req.Nickname,
	})
	if err != nil {
		g.RequestFromCtx(ctx).Response.WriteTpl("shop/login.html", g.Map{
			"tips": err,
		})
	} else {
		g.RequestFromCtx(ctx).Response.RedirectTo("/")
	}
	return
}

// 商城 post 登出监听
func (c *cShop) ShopPostLogout(ctx context.Context, req *v1.ShopPostLogoutReq) (res *v1.ShopPostLogoutRes, err error) {
	err = service.Shop().SignOut(ctx)
	user := service.Shop().GetSignedInContextUser(ctx)
	// 返回主面
	if user != nil && err != nil {

		g.RequestFromCtx(ctx).Response.WriteTpl("shop/index.html", g.Map{
			"tips":           "欢迎您" + user.Nickname,
			"display":        "none",
			"logout_display": "inline-block",
			"err":            err,
		})
	} else {
		g.RequestFromCtx(ctx).Response.RedirectTo("/")
	}
	if err != nil {

	} else {
	}
	return
}

// 商城 post 修改密码监听
func (c *cShop) ShopPostChangePassword(ctx context.Context, req *v1.ShopPostChangePasswordReq) (res *v1.ShopPostChangePasswordRes, err error) {
	err = service.Shop().ChangePassword(ctx, model.UserChangePasswordInput{
		Password: req.Password,
		Nickname: req.Nickname,
	})
	if err != nil {
		g.RequestFromCtx(ctx).Response.WriteTpl("shop/changepassword.html", g.Map{
			"tips": "密码修改失败！",
		})
	} else {
		g.RequestFromCtx(ctx).Response.WriteTpl("shop/changepassword.html", g.Map{
			"tips": req.Nickname + " 密码修改成功！",
		})
	}
	return
}

// 商城 post 注册监听
func (c *cShop) ShopPostRegister(ctx context.Context, req *v1.ShopPostRegisterReq) (res *v1.ShopPostRegisterRes, err error) {
	err = service.Shop().Create(ctx, model.UserCreateInput{
		Phone:    req.Phone,
		Password: req.Password,
		Nickname: req.Nickname,
	})
	if err != nil {
		g.RequestFromCtx(ctx).Response.WriteTpl("shop/register.html", g.Map{
			"tips": err,
		})
	} else {
		g.RequestFromCtx(ctx).Response.WriteTpl("shop/register.html", g.Map{
			"tips":    req.Nickname + " 注册成功！",
			"display": "none",
		})
	}
	return
}
