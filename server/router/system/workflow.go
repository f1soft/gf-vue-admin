package system

import (
	v1 "gf-server/app/api/v1"
	"gf-server/app/middleware"
	"gf-server/library/global"
)

// InitWorkflowRouter 注册功能api路由
func InitWorkflowRouter() {
	// TODO 缺少CasbinHandler中间件
	WorkflowRouter := global.GFVA_SERVER.Group("workflow").Middleware(middleware.JwtAuth)
	{
		WorkflowRouter.POST("createWorkFlow", v1.CreateWorkFlow) // 创建工作流
	}
}
