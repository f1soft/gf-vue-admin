package model

import (
	"github.com/gogf/gf/os/gtime"
)

type AuthorityMenu struct {
	BaseMenu
	MenuId      string           `orm:"menu_id" json:"menuId"` //菜单ID
	AuthorityId string           `orm:"authority_id" json:"-"` // 角色ID
	Children    []*AuthorityMenu `orm:"children" json:"children"`
}

type BaseMenu struct {
	Id        uint          `orm:"id,primary"   json:"id"`        // 自增ID
	CreateAt  *gtime.Time   `orm:"create_at"   json:"create_at"`  // 创建时间
	UpdateAt  *gtime.Time   `orm:"update_at"   json:"update_at"`  // 更新时间
	DeleteAt  *gtime.Time   `orm:"delete_at"   json:"delete_at"`  // 删除时间
	MenuLevel uint          `orm:"menu_level"   json:"-"`         // 菜单等级(预留字段)
	ParentId  string        `orm:"parent_id"    json:"parent_id"` // 父菜单ID
	Path      string        `orm:"path"         json:"path"`      // 路由path
	Name      string        `orm:"name"         json:"name"`      // 路由name
	Hidden    bool          `orm:"hidden"       json:"hidden"`    // 是否在列表隐藏
	Component string        `orm:"component"    json:"component"` // 前端文件路径
	Sort      int           `orm:"sort"         json:"sort"`      // 排序标记
	Meta      `json:"meta"` // 附加属性
	Children  []*BaseMenu   `orm:"children" json:"children"`
}

type Meta struct {
	Title       string `orm:"title"        json:"title"`        // 菜单名
	Icon        string `orm:"icon"         json:"icon"`         // 菜单图标
	KeepAlive   int    `orm:"keep_alive"   json:"keep_alive"`   // 是否缓存
	DefaultMenu int    `orm:"default_menu" json:"default_menu"` // 是否是基础路由(开发中)
}
