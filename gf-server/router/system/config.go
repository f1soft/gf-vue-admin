package system

import (
	"gf-server/global"
)

// InitSystemRouter 注册system配置相关路由
func InitSystemRouter() {
	// TODO 缺少JWTAuth中间件与CasbinHandler中间件
	ConfigRouter := global.GFVA_SERVER.Group("system")
	{
		ConfigRouter.POST("getSystemConfig", EmptyRequest) // 获取配置文件内容
		ConfigRouter.POST("setSystemConfig", EmptyRequest) // 设置配置文件内容
	}
}
