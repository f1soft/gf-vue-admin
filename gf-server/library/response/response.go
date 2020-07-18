package response

import (
	"gf-server/library/global"

	"github.com/gogf/gf/net/ghttp"
)

const (
	ERROR   = 7
	SUCCESS = 0
)

// 数据返回通用JSON数据结构
type JsonResponse struct {
	Code    int         `json:"code"` // 错误码((0:成功, 1:失败, >1:错误码))
	Data    interface{} `json:"data"` // 返回数据(业务接口定义具体数据结构)
	Message string      `json:"msg"`  // 提示信息
}

func Result(r *ghttp.Request, code int, data interface{}, message string) {
	if err := r.Response.WriteJson(JsonResponse{
		code,
		data,
		message,
	}); err != nil {
		global.GFVA_LOG.Error(err)
	}
}

func Ok(r *ghttp.Request) {
	Result(r, SUCCESS, map[string]interface{}{}, "操作成功")
}

func OkWithMessage(r *ghttp.Request, message string) {
	Result(r, SUCCESS, map[string]interface{}{}, message)
}

func OkWithData(r *ghttp.Request, data interface{}) {
	Result(r, SUCCESS, data, "操作成功")
}

func OkDetailed(r *ghttp.Request, data interface{}, message string) {
	Result(r, SUCCESS, data, message)
}

func Fail(r *ghttp.Request) {
	Result(r, ERROR, map[string]interface{}{}, "操作失败")
}

func FailWithMessage(r *ghttp.Request, message string) {
	Result(r, ERROR, map[string]interface{}{}, message)
}

func FailWithDetailed(r *ghttp.Request, code int, data interface{}, message string) {
	Result(r, code, data, message)
}

// 标准返回结果数据结构封装。
func Json(r *ghttp.Request, code int, message string, data ...interface{}) {
	responseData := interface{}(nil)
	if len(data) > 0 {
		responseData = data[0]
	}
	_ = r.Response.WriteJson(JsonResponse{
		Code:    code,
		Message: message,
		Data:    responseData,
	})
}

// 返回JSON数据并退出当前HTTP执行函数。
func JsonExit(r *ghttp.Request, err int, msg string, data ...interface{}) {
	Json(r, err, msg, data...)
	r.Exit()
}
