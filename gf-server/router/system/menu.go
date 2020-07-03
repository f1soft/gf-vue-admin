package system

import (
	"gf-server/global"
)

// InitMenuRouter 注册menu路由
func InitMenuRouter() {
	// TODO 缺少JWTAuth中间件与CasbinHandler中间件
	MenuRouter := global.GFVA_SERVER.Group("menu")
	{
		MenuRouter.POST("getMenu", EmptyRequest)          // 获取菜单树
		MenuRouter.POST("getMenuList", EmptyRequest)      // 分页获取基础menu列表
		MenuRouter.POST("addBaseMenu", EmptyRequest)      // 新增菜单
		MenuRouter.POST("getBaseMenuTree", EmptyRequest)  // 获取用户动态路由
		MenuRouter.POST("addMenuAuthority", EmptyRequest) //	增加menu和角色关联关系
		MenuRouter.POST("getMenuAuthority", EmptyRequest) // 获取指定角色menu
		MenuRouter.POST("deleteBaseMenu", EmptyRequest)   // 删除菜单
		MenuRouter.POST("updateBaseMenu", EmptyRequest)   // 更新菜单
		MenuRouter.POST("getBaseMenuById", EmptyRequest)  // 根据id获取菜单
	}
}
