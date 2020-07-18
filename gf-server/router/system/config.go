package system

import (
	v1 "gf-server/app/api/v1"
	"gf-server/app/middleware"
	"gf-server/library/global"
)

// InitSystemRouter 注册system配置相关路由
func InitSystemRouter() {
	// TODO 缺少CasbinHandler中间件
	ConfigRouter := global.GFVA_SERVER.Group("system").Middleware(middleware.JwtAuth)
	{
		ConfigRouter.POST("getSystemConfig", v1.GetSystemConfig) // 获取配置文件内容
		ConfigRouter.POST("setSystemConfig", v1.SetSystemConfig) // 设置配置文件内容
	}
}
