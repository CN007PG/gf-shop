package item

import (
	"context"
	"gf-shop/internal/dao"
	"gf-shop/internal/model"
	"gf-shop/internal/model/do"
	"gf-shop/internal/model/entity"
	"gf-shop/internal/service"
	"github.com/gogf/gf/v2/database/gdb"
)

type (
	sItemCat struct{}
)

func init() {
	service.RegisterItemCat(New())
}

func New() *sItemCat {
	return &sItemCat{}
}

// Create.
func (s *sItemCat) Create(ctx context.Context, in model.ItemCatCreateInput) (err error) {
	return dao.ItemCat.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err = dao.ItemCat.Ctx(ctx).Data(do.ItemCat{
			ParentId:  in.ParentId,
			Name:      in.Name,
			Status:    in.Status,
			SortOrder: in.SortOrder,
			IsParent:  in.IsParent,
		}).Insert()
		return err
	})
}

// Remove.
func (s *sItemCat) Delete(ctx context.Context, id int64) (err error) {
	return dao.ItemCat.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err = dao.ItemCat.Ctx(ctx).Where("id", id).Delete()
		return err
	})
}

// Get.
func (s *sItemCat) Get(ctx context.Context, id int64) (data *entity.ItemCat, err error) {
	var itemCat *entity.ItemCat
	err = dao.ItemCat.Ctx(ctx).Where(do.ItemCat{
		Id: id,
	}).Scan(&itemCat)
	return itemCat, err
}
