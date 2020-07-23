package request

type CreateAuthority struct {
	AuthorityId   string `p:"authority_id" v:"required|length:1,1000#请输入角色id|角色id长度为:min到:max位"`
	AuthorityName string `p:"authority_name" v:"required|length:1,1000#请输入角色名字|角色名字长度为:min到:max位"`
	ParentId      string `p:"authority_id" v:"required|length:1,1000#请输入角色父id|角色父id长度为:min到:max位"`
}

type UpdateAuthority struct {
	AuthorityId   string `p:"authority_id" v:"required|length:1,1000#请输入角色id|角色id长度为:min到:max位"`
	AuthorityName string `p:"authority_name" v:"required|length:1,1000#请输入角色名字|角色名字长度为:min到:max位"`
	ParentId      string `p:"authority_id" v:"required|length:1,1000#请输入角色父id|角色父id长度为:min到:max位"`
}

type DeleteAuthority struct {
	AuthorityId string `p:"authority_id" v:"required|length:1,1000#请输入角色id|角色id长度为:min到:max位"`
}
