package system

import (
	v1 "server/app/api/v1"
	"server/app/middleware"
	"server/library/global"
)

// InitWorkflowRouter 注册功能api路由
func InitWorkflowRouter() {
	WorkflowRouter := global.GFVA_SERVER.Group("workflow").Middleware(middleware.JwtAuth).Middleware(middleware.CasbinMiddleware)
	{
		WorkflowRouter.POST("createWorkFlow", v1.CreateWorkFlow) // 创建工作流
	}
}
