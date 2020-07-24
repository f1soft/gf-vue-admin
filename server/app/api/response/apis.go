package response

import (
	"server/app/model/apis"

	"github.com/gogf/gf/os/gtime"
)

type ApiListResponse struct {
	Apis []*Apis `json:"apis"`
}

type ApiResponse struct {
	Api *apis.Entity `json:"api"`
}

// Entity is the golang structure for table apis.
type Apis struct {
	Id          uint        `orm:"id,primary"  json:"ID"`          // 自增ID
	CreateAt    *gtime.Time `orm:"create_at"   json:"CreateAt"`    // 创建时间
	UpdateAt    *gtime.Time `orm:"update_at"   json:"UpdateAt"`    // 更新时间
	DeleteAt    *gtime.Time `orm:"delete_at"   json:"DeleteAt"`    // 删除时间
	Path        string      `orm:"path"        json:"path"`        // api路径
	Description string      `orm:"description" json:"description"` // api中文描述
	ApiGroup    string      `orm:"api_group"   json:"apiGroup"`    // api组
	Method      string      `orm:"method"      json:"method"`      // 方法
}
