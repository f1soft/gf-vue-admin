package v1

import "github.com/gogf/gf/net/ghttp"

// @Tags SysApi
// @Summary 自动代码模板
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AutoCodeStruct true "创建自动代码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /autoCode/createTemp [post]
func CreateTemp(r *ghttp.Request) {
}
