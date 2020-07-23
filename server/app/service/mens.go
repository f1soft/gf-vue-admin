package service

import (
	"server/app/model"

	"github.com/gogf/gf/frame/g"
)

// getMenuTreeMap Gets the route total tree Map
// getMenuTreeMap 获取路由总树map
func getMenuTreeMap(authorityId string) (treeMap map[string][]*model.AuthorityMenu, err error) {
	authorityMenus := ([]*model.AuthorityMenu)(nil)
	treeMap = make(map[string][]*model.AuthorityMenu)
	err = g.DB("default").Table("menus m").RightJoin("authority_menu a", "m.id=a.menu_id").Where(g.Map{"authority_id": authorityId}).Structs(&authorityMenus)
	for _, v := range authorityMenus {
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}
	return treeMap, err
}

// getChildrenList Get submenu
// getChildrenList 获取子菜单
func getChildrenList(menu *model.AuthorityMenu, treeMap map[string][]*model.AuthorityMenu) (err error) {
	menu.Children = treeMap[menu.MenuId]
	for i := 0; i < len(menu.Children); i++ {
		err = getChildrenList(menu.Children[i], treeMap)
	}
	return err
}

// GetMenuTree Gets the dynamic menu tree
// GetMenuTree 获取动态菜单树
func GetMenuTree(authorityId string) (menus []*model.AuthorityMenu, err error) {
	var menuTree map[string][]*model.AuthorityMenu
	menuTree, err = getMenuTreeMap(authorityId)
	menus = menuTree["0"]
	for i := 0; i < len(menus); i++ {
		err = getChildrenList(menus[i], menuTree)
	}
	return menus, err
}
