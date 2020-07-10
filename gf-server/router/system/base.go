package system

import (
	v1 "gf-server/app/api/v1"
	"gf-server/library/global"

	"github.com/gogf/gf-jwt/example/auth"
)

// InitBaseRouter 注册基础功能路由
func InitBaseRouter() {
	BaseRouter := global.GFVA_SERVER.Group("base")
	{
		BaseRouter.POST("register", v1.Register) // 注册
		//BaseRouter.GET("register", v1.Register) // 注册
		BaseRouter.GET("login", auth.GfJWTMiddleware.LoginHandler) // 登录
		BaseRouter.GET("captcha", v1.Captcha)                      // 验证码
	}
}
