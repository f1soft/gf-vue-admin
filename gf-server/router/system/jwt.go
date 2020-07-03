package system

import (
	"gf-server/global"
)

// InitJwtRouter 注册jwt相关路由
func InitJwtRouter() {
	// TODO 缺少JWTAuth中间件与CasbinHandler中间件
	ApiRouter := global.GFVA_SERVER.Group("jwt")
	{
		ApiRouter.POST("jsonInBlacklist", EmptyRequest) // jwt加入黑名单
	}
}
