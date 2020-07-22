package request

type CreateOperation struct {
	Ip           string `p:"ip"            v:"required | length:1, 20#请输入ip|ip长度为:min到max"`
	Method       string `p:"method"        v:"required | length:1, 20#请输入请求方法|请求方法长度为:min到max"`
	Path         string `p:"path"          v:"required | length:1, 20#请输入请求路由|请求路由长度为:min到max"`
	Status       int    `p:"status"        v:"required | length:1, 20#请输入状态|状态长度为:min到max"`
	Latency      int64  `p:"latency"       v:"required | length:1, 20#请输入延迟|延迟长度为:min到max"`
	Agent        string `p:"agent"         v:"required | length:1, 20#请输入代理|代理长度为:min到max"`
	ErrorMessage string `p:"error_message" v:"required | length:1, 20#请输入报错信息|报错信息长度为:min到max"`
	Request      string `p:"request"       v:"required | length:1, 20#请输入请求Body|请求Body长度为:min到max"`
	UserId       int    `p:"user_id"       v:"required | length:1, 20#请输入用户id|用户id长度为:min到max"`
	Response     string `p:"response"      v:"required | length:1, 20#请输入响应Body|响应Body长度为:min到max"`
}

type DeleteOperation struct {
	Id int `p:"id" v:"required|length:1,1000#请输入id|id长度为:min到:max位"`
}

type DeleteOperations struct {
	Ids []int `p:"ids" v:"required|length:1,1000#请输入ids|ids长度为:min到:max位"`
}

type UpdateOperation struct {
	Id           []int  `p:"ids" v:"required|length:1,1000#请输入id|id长度为:min到:max位"`
	Ip           string `p:"ip"            v:"required | length:1, 20#请输入ip|ip长度为:min到max"`
	Method       string `p:"method"        v:"required | length:1, 20#请输入请求方法|请求方法长度为:min到max"`
	Path         string `p:"path"          v:"required | length:1, 20#请输入请求路由|请求路由长度为:min到max"`
	Status       int    `p:"status"        v:"required | length:1, 20#请输入状态|状态长度为:min到max"`
	Latency      int64  `p:"latency"       v:"required | length:1, 20#请输入延迟|延迟长度为:min到max"`
	Agent        string `p:"agent"         v:"required | length:1, 20#请输入代理|代理长度为:min到max"`
	ErrorMessage string `p:"error_message" v:"required | length:1, 20#请输入报错信息|报错信息长度为:min到max"`
	Request      string `p:"request"       v:"required | length:1, 20#请输入请求Body|请求Body长度为:min到max"`
	UserId       int    `p:"user_id"       v:"required | length:1, 20#请输入用户id|用户id长度为:min到max"`
	Response     string `p:"response"      v:"required | length:1, 20#请输入响应Body|响应Body长度为:min到max"`
}

type FindOperation struct {
	Id int `p:"id" v:"required|length:1,1000#请输入id|id长度为:min到:max位"`
}

type GetOperationList struct {
	Method string `p:"method"        v:"required | length:1, 20#请输入请求方法|请求方法长度为:min到max"`
	Path   string `p:"path"          v:"required | length:1, 20#请输入请求路由|请求路由长度为:min到max"`
	Status int    `p:"status"        v:"required | length:1, 20#请输入状态|状态长度为:min到max"`
	PageInfo
}
