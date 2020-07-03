package system

import (
	"gf-server/global"
)

// InitCasbinRouter 注册权限相关路由
func InitCasbinRouter() {
	// TODO 缺少JWTAuth中间件与CasbinHandler中间件
	CasbinRouter := global.GFVA_SERVER.Group("casbin")
	{
		CasbinRouter.POST("updateCasbin", EmptyRequest)
		CasbinRouter.POST("getPolicyPathByAuthorityId", EmptyRequest)
		CasbinRouter.GET("casbinTest/:pathParam", EmptyRequest)
	}
}