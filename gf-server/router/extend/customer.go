package extend

import (
	"gf-server/global"
	"github.com/gogf/gf/net/ghttp"
)

// EmptyRequest 空函数
func EmptyRequest(r *ghttp.Request) {
	_ = r.Response.WriteJson("空函数, 功能开发ing")
}

// InitCustomerRouter 注册功能api路由
func InitCustomerRouter() {
	// TODO 缺少JWTAuth中间件与CasbinHandler中间件
	CustomerRouter := global.GFVA_SERVER.Group("customer")
	{
		CustomerRouter.POST("customer", EmptyRequest)    // 创建客户
		CustomerRouter.PUT("customer", EmptyRequest)     // 更新客户
		CustomerRouter.DELETE("customer", EmptyRequest)  // 删除客户
		CustomerRouter.GET("customer", EmptyRequest)     // 获取单一客户信息
		CustomerRouter.GET("customerList", EmptyRequest) // 获取客户列表
	}
}
