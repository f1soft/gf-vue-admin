package request

type UpdateBaseMenu struct {
	ParentId  string `orm:"parent_id"    json:"parent_id"` // 父菜单ID
	Path      string `orm:"path"         json:"path"`      // 路由path
	Name      string `orm:"name"         json:"name"`      // 路由name
	Component string `orm:"component"    json:"component"` // 前端文件路径
}
