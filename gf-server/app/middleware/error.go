package middleware

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func Error(r *ghttp.Request) {
	r.Middleware.Next()
	if err := r.GetError(); err != nil {
		// 记录到自定义错误日志文件
		g.Log("exception").Error(err)
		//返回固定的友好信息
		r.Response.ClearBuffer()
		r.Response.Writeln("服务器居然开小差了，请稍后再试吧！")
	}
}