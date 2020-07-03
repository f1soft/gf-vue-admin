package system

import (
	"gf-server/global"
)

// InitWorkflowRouter 注册功能api路由
func InitWorkflowRouter() {
	// TODO 缺少JWTAuth中间件与CasbinHandler中间件
	WorkflowRouter := global.GFVA_SERVER.Group("workflow")
	{
		WorkflowRouter.POST("createWorkFlow", EmptyRequest) // 创建工作流
	}
}