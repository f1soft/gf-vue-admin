package service

import (
	"github.com/gogf/gf/frame/g"
	"github.com/mojocn/base64Captcha"
)

var store = base64Captcha.DefaultMemStore

func Captcha() (id string, b64s string, err error) {
	// 字符,公式,验证码配置
	// 生成默认数字的driver
	driver := base64Captcha.NewDriverDigit(g.Cfg().GetInt("captcha.ImgHeight"), g.Cfg().GetInt("captcha.ImgWidth"), g.Cfg().GetInt("captcha.KeyLong"), 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err = cp.Generate()
	return id, b64s, err
}
