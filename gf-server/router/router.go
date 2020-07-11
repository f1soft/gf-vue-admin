package router

import (
	"gf-server/router/system"
)

// InitializeRouters 初始化总路由
func InitializeRouters() {
	system.InitUserRouter() // 初始化用户路由
	system.InitBaseRouter() // 初始化基础功能路由 不做鉴权
	//system.InitMenuRouter()             // 初始化menu路由
	//system.InitAuthorityRouter()        // 初始化角色路由
	//system.InitApiRouter()              // 初始化功能api路由
	//system.InitWorkflowRouter()         // 初始化工作流相关路由
	//system.InitCasbinRouter()           // 初始化权限相关路由
	//system.InitJwtRouter()              // 初始化jwt相关路由
	//system.InitSystemRouter()           // 初始化system配置相关路由
	//system.InitAutoCodeRouter()         // 创建自动化代码
	//system.InitDictionaryRouter()       // 初始化字典管理路由
	//system.InitDictionaryDetailRouter() // 初始化字典详情管理路由
	//extend.InitCustomerRouter()         // 初始化客户路由
	//extend.InitFileRouter()             // 初始化文件上传下载功能路由
}
