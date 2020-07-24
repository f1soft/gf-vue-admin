package request

type CreateDictionaryDetail struct {
	Label        string `p:"label" v:"label@required|length:1,1000#请输入展示值|展示值长度为:min到:max位"`         //
	Value        int    `p:"value" v:"value@required|length:1,1000#请输入展示值|展示值长度为:min到:max位"`         // 字典值
	Status       bool   `p:"status" v:"boolean@required|length:1,1000#请输入展示值|展示值长度为:min到:max位"`      // 启用状态
	Sort         int    `p:"sort" v:"label@required|length:1,1000#请输入展示值|展示值长度为:min到:max位"`          // 排序标记
	DictionaryId int    `p:"dictionary_id" v:"label@required|length:1,1000#请输入展示值|展示值长度为:min到:max位"` // 关联标记
}

type DeleteDictionaryDetail struct {
	Id float64 `p:"id" v:"id@required|length:1,1000#请输入id|id长度为:min到:max位"` // 自增ID
}

type UpdateDictionaryDetail struct {
	Id           float64 `p:"id" v:"id@required|length:1,1000#请输入id|id长度为:min到:max位"`                 // 自增ID
	Label        string  `p:"label" v:"label@required|length:1,1000#请输入展示值|展示值长度为:min到:max位"`         //
	Value        int     `p:"value" v:"value@required|length:1,1000#请输入展示值|展示值长度为:min到:max位"`         // 字典值
	Status       bool    `p:"status" v:"boolean@required|length:1,1000#请输入展示值|展示值长度为:min到:max位"`      // 启用状态
	Sort         int     `p:"sort" v:"label@required|length:1,1000#请输入展示值|展示值长度为:min到:max位"`          // 排序标记
	DictionaryId int     `p:"dictionary_id" v:"label@required|length:1,1000#请输入展示值|展示值长度为:min到:max位"` // 关联标记
}

type FindDictionaryDetail struct {
	Id           float64 `p:"id" v:"id@required|length:1,1000#请输入id|id长度为:min到:max位"`                 // 自增ID
	Label        string  `p:"label" v:"label@required|length:1,1000#请输入展示值|展示值长度为:min到:max位"`         //
	Value        int     `p:"value" v:"value@required|length:1,1000#请输入展示值|展示值长度为:min到:max位"`         // 字典值
	Status       bool    `p:"status" v:"boolean@required|length:1,1000#请输入展示值|展示值长度为:min到:max位"`      // 启用状态
	Sort         int     `p:"sort" v:"label@required|length:1,1000#请输入展示值|展示值长度为:min到:max位"`          // 排序标记
	DictionaryId int     `p:"dictionary_id" v:"label@required|length:1,1000#请输入展示值|展示值长度为:min到:max位"` // 关联标记
}

type GetDictionaryDetailList struct {
	Label        string `p:"label"`
	Value        int    `p:"value"`         // 字典值
	Status       bool   `p:"status"`        // 启用状态
	Sort         int    `p:"sort"`          // 排序标记
	DictionaryId int    `p:"dictionary_id"` // 关联标记
	PageInfo
}
