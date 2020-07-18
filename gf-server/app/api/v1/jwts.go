package v1

import (
	"fmt"
	"gf-server/app/model/jwts"
	"gf-server/app/service"
	"gf-server/library/global"

	"github.com/gogf/gf/net/ghttp"
)

// @Tags jwt
// @Summary jwt加入黑名单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"拉黑成功"}"
// @Router /jwt/jsonInBlacklist [post]
func JsonInBlacklist(r *ghttp.Request) {
	token := r.Request.Header.Get("x-token")
	if err := service.JsonInBlacklist(&jwts.Entity{Jwt: token}); err != nil {
		global.FailWithMessage(r, fmt.Sprintf("jwt作废失败，%v", err))
		r.Exit()
	}
	global.OkWithMessage(r, "jwt作废成功")
}
