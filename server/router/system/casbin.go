package system

import (
	v1 "gf-server/app/api/v1"
	"gf-server/app/middleware"
	"gf-server/library/global"
)

// InitCasbinRouter 注册权限相关路由
func InitCasbinRouter() {
	// TODO 缺少CasbinHandler中间件
	CasbinRouter := global.GFVA_SERVER.Group("casbin").Middleware(middleware.JwtAuth)
	{
		CasbinRouter.POST("updateCasbin", v1.UpdateCasbin)
		CasbinRouter.POST("getPolicyPathByAuthorityId", v1.GetPolicyPathByAuthorityId)
		CasbinRouter.GET("casbinTest/:pathParam", v1.CasbinTest)
	}
}
