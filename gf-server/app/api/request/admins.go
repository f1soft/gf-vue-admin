package request

type ChangePasswordRequest struct {
	Uuid        string
	Username    string `p:"username" v:"username@required|length:1,30#请输入用户名称|您输入用户名称长度非法"`
	Password    string `p:"password" v:"password@required|length:6,30#请输入旧密码|旧密码长度为:min到:max位"`
	NewPassword string `p:"password" v:"password@required|length:6,30#请输入新密码|新密码长度为:min到:max位"`
}
