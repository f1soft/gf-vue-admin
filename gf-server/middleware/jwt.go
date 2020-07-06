package middleware

import (
	"gf-server/app/service/auth"

	"github.com/gogf/gf/net/ghttp"
)

// JwtAuth authHook is the HOOK function implements JWT logistics.
func JwtAuth(r *ghttp.Request) {
	auth.GfJWTMiddleware.MiddlewareFunc()(r)
	r.Middleware.Next()
}
