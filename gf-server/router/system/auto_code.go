package system

import (
	v1 "gf-server/app/api/v1"
	"gf-server/app/middleware"
	"gf-server/library/global"
)

// InitJwtRouter 注册jwt相关路由
func InitAutoCodeRouter() {
	// TODO 缺少CasbinHandler中间件
	AutoCodeRouter := global.GFVA_SERVER.Group("autoCode").Middleware(middleware.JwtAuth)
	{
		AutoCodeRouter.POST("createTemp", v1.CreateTemp) // 创建自动化代码
	}
}
