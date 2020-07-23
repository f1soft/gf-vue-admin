package service

import (
	"errors"
	"server/app/api/request"
	"server/app/model"
	"server/app/model/menus"
	"strconv"

	"github.com/gogf/gf/frame/g"
)

// getMenuTreeMap Gets the route total tree Map
// getMenuTreeMap 获取路由总树map
func getMenuTreeMap(authorityId string) (treeMap map[string][]*model.AuthorityMenu, err error) {
	authorityMenus := ([]*model.AuthorityMenu)(nil)
	treeMap = make(map[string][]*model.AuthorityMenu)
	err = g.DB("default").Table("menus m").Safe().RightJoin("authority_menu a", "m.id=a.menu_id").Where(g.Map{"authority_id": authorityId}).Structs(&authorityMenus)
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

// getBaseMenuTreeMap 获取路由总树map
func getBaseMenuTreeMap() (treeMap map[string][]*model.BaseMenu, err error) {
	allMenus := ([]*model.BaseMenu)(nil)
	treeMap = make(map[string][]*model.BaseMenu)
	db := g.DB("default").Table("menus").Safe()
	err = db.Order("sort").Structs(&allMenus)
	for _, v := range allMenus {
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}
	return treeMap, err
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

// getBaseChildrenList get children of menu
// getBaseChildrenList 获取菜单的子菜单
func getBaseChildrenList(menu *model.BaseMenu, treeMap map[string][]*model.BaseMenu) (err error) {
	menu.Children = treeMap[strconv.Itoa(int(menu.Id))]
	for i := 0; i < len(menu.Children); i++ {
		err = getBaseChildrenList(menu.Children[i], treeMap)
	}
	return err
}

// GetMenuList Get routing pages
// GetMenuList 获取路由分页
func GetMenuList() (list interface{}, total int, err error) {
	var (
		treeMap  map[string][]*model.BaseMenu
		menuList []*model.BaseMenu
	)
	treeMap, err = getBaseMenuTreeMap()
	menuList = treeMap["0"]
	for i := 0; i < len(menuList); i++ {
		err = getBaseChildrenList(menuList[i], treeMap)
	}
	return menuList, total, err
}

// DeleteBaseMenu Delete the underlying route
// DeleteBaseMenu 删除基础路由
func DeleteBaseMenu(delete *request.GetById) (err error) {
	var menu *menus.Entity
	if menu, err = menus.FindOne(g.Map{"parent_id": delete.Id}); err != nil {
		return err
	}
	if menu.Id != 0 {
		_, err = menus.Delete(g.Map{"id": delete.Id})
		db := g.DB("default").Table("authority_menu").Safe()
		_, err = db.Where(g.Map{"menu_id": delete.Id}).Delete()
	}
	return errors.New("此菜单存在子菜单不可删除")
}

// UpdateBaseMenu Update the routing
// UpdateBaseMenu 更新路由
func UpdateBaseMenu(update *request.UpdateBaseMenu) (err error) {
	condition := g.Map{"id": update.Id}
	updateDate := g.Map{}
	_, err = menus.Update(updateDate, condition)
	return err
}

// GetBaseMenuById get current menus
// GetBaseMenuById 返回当前选中menu
func GetBaseMenuById(idInfo *request.GetById) (menu *model.BaseMenu, err error) {
	menu = (*model.BaseMenu)(nil)
	db := g.DB("default").Table("menus").Safe()
	err = db.Where(g.Map{"id": idInfo.Id}).Struct(&menu)
	return menu, err
}
