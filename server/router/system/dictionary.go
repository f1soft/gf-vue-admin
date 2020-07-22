package system

import (
	v1 "server/app/api/v1"
	"server/app/middleware"
	"server/library/global"
)

// InitDictionaryDetailRouter 注册字典详情管理
func InitDictionaryRouter() {
	DictionaryRouter := global.GFVA_SERVER.Group("Dictionary").Middleware(middleware.JwtAuth).Middleware(middleware.CasbinMiddleware)
	{
		DictionaryRouter.POST("createDictionary", v1.CreateDictionary)   // 新建Dictionary
		DictionaryRouter.DELETE("deleteDictionary", v1.DeleteDictionary) // 删除Dictionary
		DictionaryRouter.PUT("updateDictionary", v1.UpdateDictionary)    // 更新Dictionary
		DictionaryRouter.GET("findDictionary", v1.FindDictionary)        // 根据ID获取Dictionary
		DictionaryRouter.GET("getDictionaryList", v1.GetDictionaryList)  // 获取Dictionary列表
	}
}
