package v1

import (
	"fmt"
	"gf-server/app/api/request"
	"gf-server/app/model/admins"
	"gf-server/app/service"
	"gf-server/library/global"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// @Tags Base
// @Summary 用户注册账号
// @Produce  application/json
// @Param data body model.SysUser true "用户注册接口"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"注册成功"}"
// @Router /base/register [post]
func Register(r *ghttp.Request) {
	var R request.AdminRegister
	if err := r.Parse(&R); err != nil {
		global.FailWithMessage(r, err.Error())
	}
	u := &admins.Entity{
		Username:    R.Username,
		Password:    R.Password,
		Nickname:    R.Nickname,
		HeaderImg:   R.HeaderImg,
		AuthorityId: R.AuthorityId,
	}
	if err := service.Register(u); err != nil {
		global.FailWithMessage(r, err.Error())
		r.ExitAll()
	}
	global.OkDetailed(r, g.Map{}, "注册成功!")
}

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
		global.FailWithMessage(r, fmt.Sprintf("获取数据失败，err:%v", err))
	}
	global.OkDetailed(r, g.Map{"captchaId": id, "picPath": b64s}, "验证码获取成功")
}
