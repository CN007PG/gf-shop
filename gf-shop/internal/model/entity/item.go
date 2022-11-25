// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Item is the golang structure for table item.
type Item struct {
	Id        int64       `json:"id"        description:"商品id,商品编号"`
	Title     string      `json:"title"     description:"商品标题"`
	SellPoint string      `json:"sellPoint" description:"商品卖点"`
	Price     int64       `json:"price"     description:"商品价格单位元"`
	Num       int         `json:"num"       description:"库存数量"`
	Barcode   string      `json:"barcode"   description:"商品条形码"`
	Image     string      `json:"image"     description:"商品图片"`
	Cid       int64       `json:"cid"       description:"商品类目"`
	Status    int         `json:"status"    description:"商品状态 1-正常销售，2-停售"`
	CreateAt  *gtime.Time `json:"createAt"  description:"创建时间"`
	UpdateAt  *gtime.Time `json:"updateAt"  description:"更新时间"`
}
