package system

import (
	v1 "gf-server/app/api/v1"
	"gf-server/app/service"
	"gf-server/library/global"
)

// InitBaseRouter 注册基础功能路由
func InitBaseRouter() {
	BaseRouter := global.GFVA_SERVER.Group("base")
	{
		BaseRouter.POST("register", v1.Register)                       // 注册
		BaseRouter.POST("login", service.GfJWTMiddleware.LoginHandler) // 登录
		BaseRouter.POST("captcha", v1.Captcha)                         // 验证码
	}
}
