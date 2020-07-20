package request

// User login Structure
type LoginRequest struct {
	Username  string `p:"username" v:"username@required|length:1,30#请输入用户名称|您输入用户名称长度非法"`
	Password  string `p:"password" v:"password@required|length:6,30#请输入密码|您输入用户名称长度非法"`
	Captcha   string `json:"captcha" valid:"captcha@required#请输入正确的验证码"`
	CaptchaId string `json:"captchaId" valid:"captchaId@required|length:20#请输入captchaId|您输入captchaId长度非法"`
}

// User Register Structure
type RegisterRequest struct {
	Username    string `p:"username" v:"username@required|length:1,30#请输入用户名称|账号长度为:min到:max位"`
	Password    string `p:"password" v:"password@required|length:6,30#请输入密码|密码长度为:min到:max位"`
	Nickname    string `p:"nickname" v:"nickname@required|length:1,30#请输入昵称|昵称长度为:min到:max位"`
	HeaderImg   string `p:"header_img" v:"header_img@required|length:6,30#请输入用户头像|头像地址长度为:min到:max位"`
	AuthorityId string `p:"authority_id" v:"authority_id@required|length:1,100#请输入密码|authority_id长度为:min到:max位"`
}
