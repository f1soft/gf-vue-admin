package system

import (
	v1 "server/app/api/v1"
	"server/app/middleware"
	"server/library/global"
)

// InitCasbinRouter 注册权限相关路由
func InitCasbinRouter() {
	CasbinRouter := global.GFVA_SERVER.Group("casbin").Middleware(middleware.JwtAuth).Middleware(middleware.CasbinMiddleware)
	{
		CasbinRouter.POST("updateCasbin", v1.UpdateCasbin)
		CasbinRouter.POST("getPolicyPathByAuthorityId", v1.GetPolicyPathByAuthorityId)
		CasbinRouter.GET("casbinTest/:pathParam", v1.CasbinTest)
	}
}
