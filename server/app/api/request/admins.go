package request

// ChangePassword request Structure
type ChangePassword struct {
	Uuid        string
	Username    string `p:"username" v:"username@required|length:1,30#请输入用户名称|您输入用户名称长度非法"`
	Password    string `p:"password" v:"password@required|length:6,30#请输入旧密码|旧密码长度为:min到:max位"`
	NewPassword string `p:"password" v:"new_password@required|length:6,30#请输入新密码|新密码长度为:min到:max位"`
}

// SetAdminAuthority request Structure
type SetAdminAuthority struct {
	Uuid        string `p:"uuid" v:"uuid@required|length:16#请输入管理员UUID|您输入管理员UUID长度非法"`
	AuthorityId string `p:"authority_id" v:"uuid@required|length:1, 100#请输入角色ID|角色ID长度为:min到:max位"`
}

// DeleteAdmin request Structure
type DeleteAdmin struct {
	Id float64 `p:"id" v:"id@required|length:1,1000#请输入id|id长度为:min到:max位"`
}
