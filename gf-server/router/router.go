package router

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func MiddlewareLog(r *ghttp.Request) {
	r.Middleware.Next()
	g.Log().Println(r.Response.Status, r.URL.Path)
}

//var s *ghttp.Server
//
//func init() {
//	s = g.Server()
//	s.Use(MiddlewareLog)
//	s.Group("/", func(group *ghttp.RouterGroup) {
//		group.ALL("/", hello.Hello)
//	})
//}

func InitializeRouters() {
	InitUserRouter()
}
