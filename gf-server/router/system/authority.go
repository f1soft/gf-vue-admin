package system

import "gf-server/global"

// InitAuthorityRouter 注册角色路由组
func InitAuthorityRouter() {
	// TODO 缺少JWTAuth中间件与CasbinHandler中间件
	AuthorityRouter := global.GFVA_SERVER.Group("authority")
	{
		AuthorityRouter.POST("createAuthority", EmptyRequest)  // 创建角色
		AuthorityRouter.POST("deleteAuthority", EmptyRequest)  // 删除角色
		AuthorityRouter.PUT("updateAuthority", EmptyRequest)   // 更新角色
		AuthorityRouter.POST("copyAuthority", EmptyRequest)    // 更新角色
		AuthorityRouter.POST("getAuthorityList", EmptyRequest) // 获取角色列表
		AuthorityRouter.POST("setDataAuthority", EmptyRequest) // 设置角色资源权限
	}
}
