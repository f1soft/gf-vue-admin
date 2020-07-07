package middleware

import (
	"gf-server/app/service"
	"github.com/gogf/gf/net/ghttp"
)

// JwtAuth authHook is the HOOK function implements JWT logistics.
func JwtAuth(r *ghttp.Request) {
	service.GfJWTMiddleware.MiddlewareFunc()(r)
	r.Middleware.Next()
}