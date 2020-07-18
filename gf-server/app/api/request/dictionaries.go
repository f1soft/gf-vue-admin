package request

// CreateDictionary 创建 Dictionary
type CreateDictionary struct {
	Name   string `p:"name" v:"name@required|length:1,1000#请输入字典中文名|字典中文名长度为:min到:max位"`
	Type   string `p:"type" v:"type@required|length:1,1000#请输入字典英文名|字典英文名长度为:min到:max位"`
	Status int    `p:"status" v:"status@required|length:1,1000#请输入状态|状态长度为:min到:max位"`
	Desc   string `p:"desc" v:"desc@required|length:1,1000#请输入描述|描述长度为:min到:max位"`
}

// DeleteDictionary 更新 Dictionary
type DeleteDictionary struct {
	Id float64 `p:"id" v:"id@required|length:1,1000#请输入id|id长度为:min到:max位" orm:"id,primary"` // 自增ID
}

// UpdateDictionary 更新 Dictionary
type UpdateDictionary struct {
	Id     float64 `p:"id" v:"id@required|length:1,1000#请输入id|id长度为:min到:max位"` // 自增ID
	Name   string  `p:"name" v:"name@required|length:1,1000#请输入字典中文名|字典中文名长度为:min到:max位"`
	Type   string  `p:"type" v:"type@required|length:1,1000#请输入字典英文名|字典英文名长度为:min到:max位"`
	Status int     `p:"status" v:"status@required|length:1,1000#请输入状态|状态长度为:min到:max位"`
	Desc   string  `p:"desc" v:"desc@required|length:1,1000#请输入描述|描述长度为:min到:max位" `
}

type GetDictionary struct {
	ID   float64 `p:"id" v:"id@required|length:1,1000#请输入id|id长度为:min到:max位"`
	Type string  `p:"type" v:"type@required|length:1,1000#请输入字典英文名|字典英文名长度为:min到:max位"`
}

// FindDictionary 用id查询Dictionary
type FindDictionary struct {
	Id   float64 `p:"id" v:"id@required|length:1,1000#请输入id|id长度为:min到:max位"` // 自增ID
	Type string  `p:"type" v:"type@required|length:1,1000#请输入字典英文名|字典英文名长度为:min到:max位"`
}

// DictionaryInfoList 分页获取SysDictionary列表
type DictionaryInfoList struct {
	Name   string `p:"name" v:"name@required|length:1,1000#请输入字典中文名|字典中文名长度为:min到:max位"`
	Type   string `p:"type" v:"type@required|length:1,1000#请输入字典英文名|字典英文名长度为:min到:max位"`
	Status int    `p:"status" v:"status@required|length:1,1000#请输入状态|状态长度为:min到:max位"`
	Desc   string `p:"desc" v:"desc@required|length:1,1000#请输入描述|描述长度为:min到:max位"`
	PageInfo
}
