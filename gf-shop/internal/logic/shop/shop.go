package shop

import (
	"context"
	"gf-shop/internal/model"
	"gf-shop/internal/service"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
)

type (
	sShop struct{}
)

func init() {
	service.RegisterShop(New())
}

func New() *sShop {
	return &sShop{}
}

// Create  user.
func (s *sShop) Create(ctx context.Context, in model.UserCreateInput) (err error) {
	//存入数据库
	return service.User().Create(ctx, in)
}

// ChangePassword  user.
func (s *sShop) ChangePassword(ctx context.Context, in model.UserChangePasswordInput) (err error) {
	//数据库修改密码
	return service.User().ChangePassword(ctx, in.Nickname, in.Password)
}
func (s *sShop) GetSignedInContextUser(ctx context.Context) (user *model.ContextUser) {
	if v := service.BizCtx().Get(ctx); v != nil && v.User != nil {
		return v.User
	}
	return nil
}

// SignIn user.
func (s *sShop) SignIn(ctx context.Context, in model.UserSignInInput) (err error) {
	user, err := service.User().Get(ctx, gconv.String(in.Nickname))
	if err != nil {
		return err
	}
	if user == nil || user.Password != in.Password {
		return gerror.New(`用户名或密码出错`)
	}
	if err = service.Session().SetUser(ctx, user); err != nil {
		return err
	}
	service.BizCtx().SetUser(ctx, &model.ContextUser{
		Id:       user.Id,
		Nickname: user.Name,
	})
	return nil
}

// SignOut user.
func (s *sShop) SignOut(ctx context.Context) (err error) {
	return service.Session().RemoveUser(ctx)
}
