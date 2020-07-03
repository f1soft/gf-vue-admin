package system

import (
	"gf-server/global"
)

// InitJwtRouter 注册jwt相关路由
func InitAutoCodeRouter() {
	// TODO 缺少JWTAuth中间件与CasbinHandler中间件
	AutoCodeRouter := global.GFVA_SERVER.Group("autoCode")
	{
		AutoCodeRouter.POST("createTemp", EmptyRequest) // 创建自动化代码
	}
}
