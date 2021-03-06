package system

import (
	v1 "server/app/api/v1"
	"server/app/middleware"
	"server/library/global"
)

// InitAuthorityRouter 注册角色路由组
func InitAuthorityRouter() {
	AuthorityRouter := global.GFVA_SERVER.Group("authority").Middleware(middleware.JwtAuth).Middleware(middleware.CasbinMiddleware)
	{
		AuthorityRouter.POST("createAuthority", v1.CreateAuthority)   // 创建角色
		AuthorityRouter.POST("deleteAuthority", v1.DeleteAuthority)   // 删除角色
		AuthorityRouter.PUT("deleteAuthority", v1.UpdateAuthority)    // 更新角色
		AuthorityRouter.POST("copyAuthority", v1.CopyAuthority)       // 更新角色
		AuthorityRouter.POST("getAuthorityList", v1.GetAuthorityList) // 获取角色列表
		AuthorityRouter.POST("setDataAuthority", v1.SetDataAuthority) // 设置角色资源权限
	}
}
