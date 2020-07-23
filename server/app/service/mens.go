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

func GetMenuAuthority(authorityId string) (menusReturn []model.AuthorityMenu, err error) {
	return menusReturn, nil
}
