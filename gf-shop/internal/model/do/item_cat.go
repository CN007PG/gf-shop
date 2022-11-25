// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ItemCat is the golang structure of table item_cat for DAO operations like Where/Data.
type ItemCat struct {
	g.Meta    `orm:"table:item_cat, do:true"`
	Id        interface{} //
	ParentId  interface{} // 父类目id
	Name      interface{} // 类目名称
	Status    interface{} // 状态 1-正常销售，2-停售
	SortOrder interface{} // 排序号
	IsParent  interface{} // 是否为父类目
	CreateAt  *gtime.Time // 创建时间
	UpdateAt  *gtime.Time // 更新时间
}