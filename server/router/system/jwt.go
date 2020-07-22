package system

import (
	v1 "server/app/api/v1"
	"server/app/middleware"
	"server/library/global"
)

// InitJwtRouter 注册jwt相关路由
func InitJwtRouter() {
	ApiRouter := global.GFVA_SERVER.Group("jwt").Middleware(middleware.JwtAuth).Middleware(middleware.CasbinMiddleware)
	{
		ApiRouter.POST("jsonInBlacklist", v1.JsonInBlacklist) // jwt加入黑名单
	}
}
