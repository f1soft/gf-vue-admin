package system

import (
	v1 "gf-server/app/api/v1"
	"gf-server/app/middleware"
	"gf-server/library/global"
)

// InitDictionaryDetailRouter 注册字典详情管理
func InitDictionaryRouter() {
	// TODO 缺少CasbinHandler中间件
	DictionaryRouter := global.GFVA_SERVER.Group("sysDictionary").Middleware(middleware.JwtAuth)
	{
		DictionaryRouter.POST("createSysDictionary", v1.CreateSysDictionary)   // 新建SysDictionary
		DictionaryRouter.DELETE("deleteSysDictionary", v1.DeleteSysDictionary) // 删除SysDictionary
		DictionaryRouter.PUT("updateSysDictionary", v1.UpdateSysDictionary)    // 更新SysDictionary
		DictionaryRouter.GET("findSysDictionary", v1.FindSysDictionary)        // 根据ID获取SysDictionary
		DictionaryRouter.GET("getSysDictionaryList", v1.GetSysDictionaryList)  // 获取SysDictionary列表
	}
}
