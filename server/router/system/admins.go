package system

import (
	v1 "server/app/api/v1"
	"server/app/middleware"

	"github.com/gogf/gf/frame/g"
)

// InitAdminsRouter 注册管理员路由
func InitAdminsRouter() {
	UserRouter := g.Server().Group("user").Middleware(middleware.JwtAuth)
	//UserRouter := global.GFVA_SERVER.Group("user").Middleware(middleware.JwtAuth).Middleware(middleware.CasbinMiddleware)
	{
		UserRouter.POST("changePassword", v1.ChangePassword)     // 修改密码
		UserRouter.POST("uploadHeaderImg", v1.UploadHeaderImg)   // 上传头像
		UserRouter.POST("getUserList", v1.GetAdminList)          // 分页获取用户列表
		UserRouter.POST("setUserAuthority", v1.SetUserAuthority) // 设置用户权限
		UserRouter.DELETE("deleteUser", v1.DeleteAdmin)          // 删除用户
	}
}
