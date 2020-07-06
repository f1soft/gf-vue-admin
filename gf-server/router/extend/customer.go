package extend

import (
	"errors"
	"gf-server/global"
	"gf-server/library/response"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// EmptyRequest 空函数
func EmptyRequest(r *ghttp.Request) {
	global.GFVA_LOG.Error(errors.New("测试BUG打印"))
	response.Json(r, 0,"操作成功",g.Map{})
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
