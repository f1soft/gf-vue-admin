package middleware

import (
	"gf-server/app/service/auth"

	"github.com/gogf/gf/net/ghttp"
)

// MiddlewareAuth authHook is the HOOK function implements JWT logistics.
func MiddlewareAuth(r *ghttp.Request) {
	auth.GfJWTMiddleware.MiddlewareFunc()(r)
	r.Middleware.Next()
}
