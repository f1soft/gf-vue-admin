package middleware

import (
	"gf-server/app/service"
	"gf-server/library/global"

	jwt "github.com/gogf/gf-jwt"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

// JwtAuth authHook is the HOOK function implements JWT logistics.
//func JwtAuth(r *ghttp.Request) {
//	service.GfJWTMiddleware.MiddlewareFunc()(r)
//	r.Middleware.Next()
//}

// JwtAuth 验证token有效性
func JwtAuth(r *ghttp.Request) {
	service.GfJWTMiddleware.MiddlewareFunc()(r)
	parseToken, err := service.GfJWTMiddleware.ParseToken(r) // 解析token
	if err != nil {
		if err == jwt.ErrExpiredToken {
			global.Result(r, global.ERROR, g.Map{"reload": true}, "授权已过期")
			r.ExitAll()
		}
		global.Result(r, global.ERROR, g.Map{"reload": true}, err.Error())
		r.ExitAll()
	}
	var token = parseToken.Raw
	if service.IsBlacklist(token) {
		global.Result(r, global.ERROR, g.Map{"reload": true}, "您的帐户异地登陆或令牌失效")
		r.Exit()
	}
	var claims = gconv.Map(parseToken.Claims)
	r.SetParam("uuid", claims["uuid"])
	if !service.ValidatorRedisToken(gconv.String(claims["uuid"]), token) {
		global.FailWithMessage(r, "Token鉴权失败")
	}
	r.Middleware.Next()
}
