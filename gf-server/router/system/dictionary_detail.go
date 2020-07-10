package system

import (
	v1 "gf-server/app/api/v1"
	"gf-server/library/global"
)

// InitDictionaryDetailRouter 注册字典详情管理路由
func InitDictionaryDetailRouter() {
	// TODO 缺少CasbinHandler中间件
	DictionaryDetailRouter := global.GFVA_SERVER.Group("sysDictionaryDetail")
	{
		DictionaryDetailRouter.POST("createSysDictionaryDetail", v1.CreateSysDictionaryDetail)   // 新建SysDictionaryDetail
		DictionaryDetailRouter.DELETE("deleteSysDictionaryDetail", v1.CreateSysDictionaryDetail) // 删除SysDictionaryDetail
		DictionaryDetailRouter.PUT("updateSysDictionaryDetail", v1.CreateSysDictionaryDetail)    // 更新SysDictionaryDetail
		DictionaryDetailRouter.GET("findSysDictionaryDetail", v1.CreateSysDictionaryDetail)      // 根据ID获取SysDictionaryDetail
		DictionaryDetailRouter.GET("getSysDictionaryDetailList", v1.CreateSysDictionaryDetail)   // 获取SysDictionaryDetail列表
	}
}
