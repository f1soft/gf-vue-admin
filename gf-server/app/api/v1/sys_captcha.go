package v1

import (
	"gf-server/app/service"
	"gf-server/library/global"
	"gf-server/library/response"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// @Tags base
// @Summary 生成验证码
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /base/captcha [post]
func Captcha(r *ghttp.Request) {
	id, b64s, err := service.Captcha()
	if err != nil {
		global.GFVA_LOG.Errorf("获取数据失败，err:%v", err)
		response.FailWithMessage(r, "获取数据失败")
	}
	response.OkDetailed(r, g.Map{"captchaId": id, "picPath": b64s}, "验证码获取成功")
}
