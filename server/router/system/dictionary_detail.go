package system

import (
	v1 "server/app/api/v1"
	"server/app/middleware"
	"server/library/global"
)

// InitDictionaryDetailRouter 注册字典详情管理路由
func InitDictionaryDetailRouter() {
	DictionaryDetailRouter := global.GFVA_SERVER.Group("DictionaryDetail").Middleware(middleware.CasbinMiddleware)
	{
		DictionaryDetailRouter.POST("createDictionaryDetail", v1.CreateDictionaryDetail)   // 新建DictionaryDetail
		DictionaryDetailRouter.DELETE("deleteDictionaryDetail", v1.DeleteDictionaryDetail) // 删除DictionaryDetail
		DictionaryDetailRouter.PUT("updateDictionaryDetail", v1.UpdateDictionaryDetail)    // 更新DictionaryDetail
		DictionaryDetailRouter.GET("findDictionaryDetail", v1.FindDictionaryDetail)        // 根据ID获取DictionaryDetail
		DictionaryDetailRouter.GET("getDictionaryDetailList", v1.GetDictionaryDetailList)  // 获取DictionaryDetail列表
	}
}
