package system

import (
	"gf-server/global"
)

// InitDictionaryDetailRouter 注册字典详情管理路由
func InitDictionaryDetailRouter() {
	// TODO 缺少JWTAuth中间件与CasbinHandler中间件
	DictionaryDetailRouter := global.GFVA_SERVER.Group("sysDictionaryDetail")
	{
		DictionaryDetailRouter.POST("createSysDictionaryDetail", EmptyRequest)   // 新建SysDictionaryDetail
		DictionaryDetailRouter.DELETE("deleteSysDictionaryDetail", EmptyRequest) // 删除SysDictionaryDetail
		DictionaryDetailRouter.PUT("updateSysDictionaryDetail", EmptyRequest)    // 更新SysDictionaryDetail
		DictionaryDetailRouter.GET("findSysDictionaryDetail", EmptyRequest)      // 根据ID获取SysDictionaryDetail
		DictionaryDetailRouter.GET("getSysDictionaryDetailList", EmptyRequest)   // 获取SysDictionaryDetail列表
	}
}
