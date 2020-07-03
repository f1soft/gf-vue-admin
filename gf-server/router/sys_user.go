package router

import (
	"gf-server/global"
	"github.com/gogf/gf/net/ghttp"
)

type Controller struct{}

// RESTFul - GET
func (c *Controller) Get(r *ghttp.Request) {
	r.Response.Write("GET")
}

// RESTFul - POST
func (c *Controller) Post(r *ghttp.Request) {
	r.Response.Write("POST")
}

// RESTFul - DELETE
func (c *Controller) Delete(r *ghttp.Request) {
	r.Response.Write("DELETE")
}


func InitUserRouter(){

		//UserRouter := Router.Group("user").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
		//{
		//	UserRouter.POST("changePassword", v1.ChangePassword)     // 修改密码
		//	UserRouter.POST("uploadHeaderImg", v1.UploadHeaderImg)   // 上传头像
		//	UserRouter.POST("getUserList", v1.GetUserList)           // 分页获取用户列表
		//	UserRouter.POST("setUserAuthority", v1.SetUserAuthority) // 设置用户权限
		//	UserRouter.DELETE("deleteUser", v1.DeleteUser)           // 删除用户
		//}
	UserRouter := global.GFVA_SERVER.Group("/api.v2")
	{
		UserRouter.GET("/test", func(r *ghttp.Request) {
			r.Response.Write("test")
		})
	}
	//global.GFVA_SERVER.Group("/api.v2", func(group *ghttp.RouterGroup) {
	//	group.GET("/test", func(r *ghttp.Request) {
	//		r.Response.Write("test")
	//	})
	//})
	//c := Controller{}
	//s.BindObjectRest("/object", c)
	//return s
}
