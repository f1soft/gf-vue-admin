package v1

import (
	"fmt"
	"server/app/model/jwts"
	"server/app/service"
	"server/library/global"
	"strings"

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
	token := r.Request.Header.Get("Authorization")
	parts := strings.SplitN(token, " ", 2)
	if err := service.JsonInBlacklist(&jwts.Entity{Jwt: parts[1]}); err != nil {
		global.FailWithMessage(r, fmt.Sprintf("jwt作废失败，%v", err))
		r.Exit()
	}
	global.OkWithMessage(r, "jwt作废成功")
}
