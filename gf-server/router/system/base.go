package system

import (
	v1 "gf-server/app/api/v1"
	"gf-server/global"
)

// InitBaseRouter 注册基础功能路由
func InitBaseRouter() {
	BaseRouter := global.GFVA_SERVER.Group("base")
	{
		BaseRouter.GET("register", v1.Register) // 注册
		BaseRouter.GET("login", v1.Login)       // 登录
		BaseRouter.GET("captcha", v1.Captcha)       // 验证码
	}
}
