package request

type CreateApi struct {
	Path        string `p:"path" v:"path@required|length:1,100#请输入api路径|api路径长度为:min到:max位"`
	Description string `p:"description" v:"description@required|length:1,100#请输入api中文描述|api中文描述长度为:min到:max位"`
	ApiGroup    string `p:"apiGroup" v:"apiGroup@required|length:1,100#请输入api组|api组描述长度为:min到:max位"`
	Method      string `p:"method" v:"method@required|length:1,100#请输入api请求方法|api请求方法长度为:min到:max位"`
}

type DeleteApi struct {
	ID int `p:"id" v:"method@required|length:1,1000#请输入ID|ID长度为:min到:max位"`
	Path        string `p:"path" v:"path@required|length:1,100#请输入api路径|api路径长度为:min到:max位"`
	Method      string `p:"method" v:"method@required|length:1,100#请输入api请求方法|api请求方法长度为:min到:max位"`
}

type UpdateApi struct {
	ID          float64 `p:"id" v:"method@required|length:1,1000#请输入ID|ID长度为:min到:max位"`
	Path        string  `p:"path" v:"path@required|length:1,100#请输入api路径|api路径长度为:min到:max位"`
	Description string  `p:"description" v:"description@required|length:1,100#请输入api中文描述|api中文描述长度为:min到:max位"`
	ApiGroup    string  `p:"apiGroup" v:"apiGroup@required|length:1,100#请输入api组|api组描述长度为:min到:max位"`
	Method      string  `p:"method" v:"method@required|length:1,100#请输入api请求方法|api请求方法长度为:min到:max位"`
}

type GetApiById struct {
	ID float64 `p:"id" v:"method@required|length:1,1000#请输入ID|ID长度为:min到:max位"`
}

// api分页条件查询及排序结构体
type GetApiList struct {
	ID          float64 `p:"id" v:"method@required|length:1,1000#请输入ID|ID长度为:min到:max位"`
	Path        string  `p:"path" v:"path@required|length:1,100#请输入api路径|api路径长度为:min到:max位"`
	Description string  `p:"description" v:"description@required|length:1,100#请输入api中文描述|api中文描述长度为:min到:max位"`
	ApiGroup    string  `p:"apiGroup" v:"apiGroup@required|length:1,100#请输入api组|api组描述长度为:min到:max位"`
	Method      string  `p:"method" v:"method@required|length:1,100#请输入api请求方法|api请求方法长度为:min到:max位"`
	PageInfo
	OrderKey string `p:"orderKey"`
	Desc     bool   `p:"desc"`
}
