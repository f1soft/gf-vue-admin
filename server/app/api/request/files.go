package request

type DeleteFile struct {
	Id float64 `p:"id" v:"id@required|length:1,1000#请输入id|id长度为:min到:max位"` // 自增ID
}
