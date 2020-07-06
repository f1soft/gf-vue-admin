package request


// User login structure
type LoginStruct struct {
	Username  string `json:"username" valid:"username@required|length:1,30#请输入用户名称|您输入用户名称长度非法"`
	Password  string `json:"password" valid:"password@required|length:6,30#请输入密码|您输入用户名称长度非法"`
	Captcha   string `json:"captcha" valid:"captcha@required#请输入正确的验证码"`
	CaptchaId string `json:"captchaId" valid:"captchaId@required|length:20#请输入captchaId|您输入captchaId长度非法"`
}