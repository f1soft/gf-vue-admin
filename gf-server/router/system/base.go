package system

import (
	"gf-server/global"
)

// InitBaseRouter 注册基础功能路由
func InitBaseRouter() {
	BaseRouter := global.GFVA_SERVER.Group("base")
	{
		BaseRouter.GET("register", EmptyRequest) // 注册
		BaseRouter.GET("login", EmptyRequest)    // 登录
		BaseRouter.GET("captcha", EmptyRequest)  // 验证码
	}
}
