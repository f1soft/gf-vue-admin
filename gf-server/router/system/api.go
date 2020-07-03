package system

import "gf-server/global"

// InitApiRouter 注册功能api路由
func InitApiRouter() {
	// TODO 缺少JWTAuth中间件与CasbinHandler中间件
	ApiRouter := global.GFVA_SERVER.Group("api")
	{
		ApiRouter.POST("createApi", EmptyRequest)  // 创建Api
		ApiRouter.POST("deleteApi", EmptyRequest)  // 删除Api
		ApiRouter.POST("getApiList", EmptyRequest) // 获取Api列表
		ApiRouter.POST("getApiById", EmptyRequest) // 获取单条Api消息
		ApiRouter.POST("updateApi", EmptyRequest)  // 更新api
		ApiRouter.POST("getAllApis", EmptyRequest) // 获取所有api
	}
}
