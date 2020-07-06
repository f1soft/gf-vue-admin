package v1

import (
	"gf-server/global"
	"gf-server/library/response"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/mojocn/base64Captcha"
)

var store = base64Captcha.DefaultMemStore

// @Tags base
// @Summary 生成验证码
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /base/captcha [post]
func Captcha(r *ghttp.Request) {
	//字符,公式,验证码配置
	// 生成默认数字的driver
	driver := base64Captcha.NewDriverDigit(g.Cfg().GetInt("captcha.ImgHeight"), g.Cfg().GetInt("captcha.ImgWidth"), g.Cfg().GetInt("captcha.KeyLong"), 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := cp.Generate()
	if err != nil {
		global.GFVA_LOG.Errorf("获取数据失败，err:%v", err)
		response.FailWithMessage(r, "获取数据失败")
	}
	response.OkDetailed(r,g.Map{"captchaId":id, "picPath":b64s}, "验证码获取成功")
}
