package system

import (
	v1 "server/app/api/v1"
	"server/app/middleware"
	"server/library/global"
)

// InitOperationRouter 注册操作记录路由
func InitOperationRouter() {
	// TODO 缺少CasbinHandler中间件
	OperationRouter := global.GFVA_SERVER.Group("operation").Middleware(middleware.JwtAuth)
	{
		OperationRouter.POST("CreateOperation", v1.CreateOperation)     // 新建Operation
		OperationRouter.DELETE("deleteOperation", v1.DeleteOperation)   // 删除Operation
		OperationRouter.DELETE("deleteOperations", v1.DeleteOperations) // 批量删除Operation
		OperationRouter.PUT("updateOperation", v1.UpdateOperation)      // 更新Operation
		OperationRouter.GET("findOperation", v1.FindOperation)          // 根据ID获取Operation
		OperationRouter.GET("getOperationList", v1.GetOperationList)    // 获取Operation列表
	}
}
