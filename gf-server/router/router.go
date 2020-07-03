package router

import (
	"gf-server/app/api/hello"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func MiddlewareLog(r *ghttp.Request) {
	r.Middleware.Next()
	g.Log().Println(r.Response.Status, r.URL.Path)
}

var s *ghttp.Server

func init() {
	s = g.Server()
	s.Use(MiddlewareLog)
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.ALL("/", hello.Hello)
	})
}

func InitializeRouters() {
	group := s.Group("/api")
	group.ALL("/all", func(r *ghttp.Request) {
		r.Response.Write("all")
	})
	group.GET("/get", func(r *ghttp.Request) {
		r.Response.Write("get")
	})
	group.POST("/post", func(r *ghttp.Request) {
		r.Response.Write("post")
	})
}
