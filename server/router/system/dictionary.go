package system

import (
	v1 "server/app/api/v1"
	"server/app/middleware"
	"server/library/global"
)

// InitDictionaryDetailRouter 注册字典详情管理
func InitDictionaryRouter() {
	DictionaryRouter := global.GFVA_SERVER.Group("sysDictionary").Middleware(middleware.JwtAuth)
	//DictionaryRouter := global.GFVA_SERVER.Group("sysDictionary").Middleware(middleware.JwtAuth).Middleware(middleware.CasbinMiddleware)
	{
		DictionaryRouter.POST("createSysDictionary", v1.CreateDictionary)   // 新建Dictionary
		DictionaryRouter.DELETE("deleteSysDictionary", v1.DeleteDictionary) // 删除Dictionary
		DictionaryRouter.PUT("updateSysDictionary", v1.UpdateDictionary)    // 更新Dictionary
		DictionaryRouter.GET("findSysDictionary", v1.FindDictionary)        // 根据ID获取Dictionary
		DictionaryRouter.GET("getSysDictionaryList", v1.GetDictionaryList)  // 获取Dictionary列表
	}
}
