package user

import (
	"context"
	"gf-shop/internal/dao"
	"gf-shop/internal/model"
	"gf-shop/internal/model/do"
	"gf-shop/internal/model/entity"
	"gf-shop/internal/service"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

type (
	sUser struct{}
)

func init() {
	service.RegisterUser(New())
}

func New() *sUser {
	return &sUser{}
}

// Create.
func (s *sUser) Create(ctx context.Context, in model.UserCreateInput) (err error) {
	return dao.User.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err = dao.User.Ctx(ctx).Data(do.User{
			Phone:    in.Phone,
			Password: in.Password,
			Name:     in.Nickname,
		}).Insert()
		return err
	})
}

// ChangePassword.
func (s *sUser) ChangePassword(ctx context.Context, name string, password string) (err error) {
	return dao.User.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err = dao.User.Ctx(ctx).Data(g.Map{"password": password}).Where("name", name).Update()
		return err
	})
}

// Update.
func (s *sUser) Update(ctx context.Context, in model.UserUpdateInput) (err error) {
	return dao.User.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err = dao.User.Ctx(ctx).Data(do.User{
			Phone:      in.Phone,
			Password:   in.Password,
			Name:       in.Name,
			BuyHistory: in.BuyHistory,
			BuyCart:    in.BuyCart,
			Token:      in.Token,
		}).Update()
		return err
	})
}

// Get.
func (s *sUser) Get(ctx context.Context, name string) (data *entity.User, err error) {
	var user *entity.User
	err = dao.User.Ctx(ctx).Where(do.User{
		Name: name,
	}).Scan(&user)
	return user, err
}
