package request

import (
	"server/app/model"

	"github.com/gogf/gf/os/gtime"
)

type UpdateBaseMenu struct {
	Id        uint        `r:"ID"`
	CreateAt  *gtime.Time `r:"CreateAt"`
	UpdateAt  *gtime.Time `r:"UpdateAt"`
	DeleteAt  *gtime.Time `r:"DeleteAt"`
	ParentId  string      `r:"parentId" v:"required|length:1,1000#请输入父菜单ID|父菜单ID长度为:min到:max位"`
	Path      string      `r:"path" v:"required|length:1,1000#请输入路由path|路由path长度为:min到:max位"`
	Name      string      `r:"name" v:"required|length:1,1000#请输入路由Name|路由Name长度为:min到:max位"`
	Hidden    bool        `r:"hidden" v:"required|length:1,1000#请输入是否在列表隐藏|是否在列表隐藏长度为:min到:max位"`
	Component string      `r:"component" v:"required|length:1,1000#请输入前端文件路径|前端文件路径长度为:min到:max位"`
	Sort      int         `r:"sort" v:"required|length:1,1000#请输入排序标记|排序标记长度为:min到:max位"`
	Meta      `r:"meta"`  // 附加属性
}

type Meta struct {
	Title       string `r:"title" v:"required|length:1,1000#请输入菜单名|id长度为:min到:max位"`
	Icon        string `r:"icon" v:"required|length:1,1000#请输入菜单图标|id长度为:min到:max位"`
	KeepAlive   int    `r:"keepAlive" v:"required|length:1,1000#请输入是否缓存|是否缓存长度为:min到:max位"`
	DefaultMenu bool   `r:"defaultMenu" v:"required|length:1,1000#请输入是否是基础路由(开发中)|是否是基础路由(开发中)长度为:min到:max位"`
}

type CreateBaseMenu struct {
	ParentId  string `r:"parentId" v:"required|length:1,1000#请输入父菜单ID|父菜单ID长度为:min到:max位"`
	Path      string `r:"path" v:"required|length:1,1000#请输入路由path|路由path长度为:min到:max位"`
	Name      string `r:"name" v:"required|length:1,1000#请输入路由Name|路由Name长度为:min到:max位"`
	Hidden    bool   `r:"hidden" v:"required|length:1,1000#请输入是否在列表隐藏|是否在列表隐藏长度为:min到:max位"`
	Component string `r:"component" v:"required|length:1,1000#请输入前端文件路径|前端文件路径长度为:min到:max位"`
	Sort      int    `r:"sort" v:"required|length:1,1000#请输入排序标记|排序标记长度为:min到:max位"`
}

// Get role by id structure
type AuthorityIdInfo struct {
	AuthorityId string `r:"authority_id" v:"required|length:1,1000#请输入权限ID|权限ID长度为:min到:max位"`
}

// Add menu authority info structure
type AddMenuAuthorityInfo struct {
	Menus       []model.BaseMenu
	AuthorityId string
}
