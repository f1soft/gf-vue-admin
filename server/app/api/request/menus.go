package request

type UpdateBaseMenu struct {
	Id        int    `p:"id" v:"required|length:1,1000#请输入id|id长度为:min到:max位"`
	ParentId  string `orm:"parent_id"    json:"parent_id"` // 父菜单ID
	Path      string `orm:"path"         json:"path"`      // 路由path
	Name      string `orm:"name"         json:"name"`      // 路由name
	Component string `orm:"component"    json:"component"` // 前端文件路径
}
