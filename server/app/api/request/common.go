package request

// Paging common input parameter structure
type PageInfo struct {
	Page     int `p:"page" v:"required|length:1,1000#请输入页数|页数长度为:min到:max位" json:"page" form:"page"`
	PageSize int `p:"pageSize" v:"required|length:1,1000#请输入每页大小|每页大小为:min到:max位" json:"pageSize" form:"pageSize"`
}

// Find by id structure
type GetById struct {
	Id int `p:"id" v:"required|length:1,1000#请输入id|id长度为:min到:max位"`
}
