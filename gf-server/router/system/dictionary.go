package system

import (
	"gf-server/global"
)

// InitDictionaryDetailRouter 注册字典详情管理
func InitDictionaryRouter() {
	// TODO 缺少JWTAuth中间件与CasbinHandler中间件
	DictionaryRouter := global.GFVA_SERVER.Group("sysDictionary")
	{
		DictionaryRouter.POST("createSysDictionary", EmptyRequest)   // 新建SysDictionary
		DictionaryRouter.DELETE("deleteSysDictionary", EmptyRequest) // 删除SysDictionary
		DictionaryRouter.PUT("updateSysDictionary", EmptyRequest)    // 更新SysDictionary
		DictionaryRouter.GET("findSysDictionary", EmptyRequest)      // 根据ID获取SysDictionary
		DictionaryRouter.GET("getSysDictionaryList", EmptyRequest)   // 获取SysDictionary列表
	}
}
