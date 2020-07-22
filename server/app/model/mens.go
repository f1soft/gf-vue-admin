package model

import "gf-server/app/model/menus"

type AuthorityMenus struct {
	menus.Entity
	MenuId      string         `json:"menuId" gorm:"comment:'菜单ID'"`
	AuthorityId string         `json:"-" gorm:"comment:'角色ID'"`
	Children    []menus.Entity `json:"children"`
}
