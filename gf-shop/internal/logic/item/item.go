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
	sItem struct{}
)

func init() {
	service.RegisterItem(New())
}

func New() *sItem {
	return &sItem{}
}

// Create.
func (s *sItem) Create(ctx context.Context, in model.ItemCreateInput) (err error) {
	return dao.Item.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err = dao.Item.Ctx(ctx).Data(do.Item{
			Title:     in.Title,
			SellPoint: in.SellPoint,
			Price:     in.Price,
			Num:       in.Num,
			Barcode:   in.Barcode,
			Image:     in.Image,
			Cid:       in.Cid,
			Status:    in.Status,
		}).Insert()
		return err
	})
}

// Remove.
func (s *sItem) Delete(ctx context.Context, id int64) (err error) {
	return dao.Item.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err = dao.Item.Ctx(ctx).Where("id", id).Delete()
		return err
	})
}

// Get.
func (s *sItem) Get(ctx context.Context, id int64) (data *entity.Item, err error) {
	var item *entity.Item
	err = dao.Item.Ctx(ctx).Where(do.Item{
		Id: id,
	}).Scan(&item)
	return item, err
}
