package extend

import (
	v1 "gf-server/app/api/v1"
	"gf-server/app/middleware"
	"gf-server/library/global"
)

// InitCustomerRouter 注册功能api路由
func InitCustomerRouter() {
	// TODO 缺少CasbinHandler中间件
	CustomerRouter := global.GFVA_SERVER.Group("customer").Middleware(middleware.JwtAuth)
	{
		CustomerRouter.POST("customer", v1.CreateExaCustomer)     // 创建客户
		CustomerRouter.PUT("customer", v1.UpdateExaCustomer)      // 更新客户
		CustomerRouter.DELETE("customer", v1.DeleteExaCustomer)   // 删除客户
		CustomerRouter.GET("customer", v1.GetExaCustomer)         // 获取单一客户信息
		CustomerRouter.GET("customerList", v1.GetExaCustomerList) // 获取客户列表
	}
}
