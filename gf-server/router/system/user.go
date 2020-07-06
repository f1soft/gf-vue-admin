package system

import (
	"gf-server/global"
	"gf-server/library/response"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// EmptyRequest 空函数
func EmptyRequest(r *ghttp.Request) {
	response.Json(r, 0, "操作成功", g.Map{})
}

// InitUserRouter 注册用户路由
func InitUserRouter() {
	// TODO 缺少JWTAuth中间件与CasbinHandler中间件
	UserRouter := global.GFVA_SERVER.Group("user")
	{
		UserRouter.POST("changePassword", EmptyRequest)   // 修改密码
		UserRouter.POST("uploadHeaderImg", EmptyRequest)  // 上传头像
		UserRouter.POST("getUserList", EmptyRequest)      // 分页获取用户列表
		UserRouter.POST("setUserAuthority", EmptyRequest) // 设置用户权限
		UserRouter.DELETE("deleteUser", EmptyRequest)     // 删除用户
	}
}
