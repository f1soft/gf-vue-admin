package request

// Paging common input parameter structure
type PageInfo struct {
	Page     int `p:"page" v:"page@required|length:1,1000#请输入页数|页数长度为:min到:max位" json:"page" form:"page"`
	PageSize int `p:"pageSize" v:"pageSize@required|length:1,1000#请输入每页大小|每页大小为:min到:max位" json:"pageSize" form:"pageSize"`
}
