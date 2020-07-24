package v1

import (
	"fmt"
	"server/app/api/request"
	"server/app/api/response"
	"server/app/service"
	"server/library/global"

	"github.com/gogf/gf/frame/g"

	"github.com/gogf/gf/net/ghttp"
)

// @Tags DictionaryDetail
// @Summary 创建DictionaryDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DictionaryDetail true "创建DictionaryDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /DictionaryDetail/createDictionaryDetail [post]
func CreateDictionaryDetail(r *ghttp.Request) {
	var createInfo request.CreateDictionaryDetail
	if err := r.Parse(&createInfo); err != nil {
		global.FailWithMessage(r, err.Error())
	}
	if err := service.CreateDictionaryDetail(&createInfo); err != nil {
		global.FailWithMessage(r, fmt.Sprintf("创建失败，err:%v", err))
		r.Exit()
	}
	global.OkWithMessage(r, "创建成功")
}

// @Tags DictionaryDetail
// @Summary 删除DictionaryDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DictionaryDetail true "删除DictionaryDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /DictionaryDetail/deleteDictionaryDetail [delete]
func DeleteDictionaryDetail(r *ghttp.Request) {
	var deleteRequest request.DeleteDictionaryDetail
	if err := r.Parse(&deleteRequest); err != nil {
		global.FailWithMessage(r, err.Error())
	}
	if err := service.DeleteDictionaryDetail(&deleteRequest); err != nil {
		global.FailWithMessage(r, fmt.Sprintf("删除失败，err:%v", err))
		r.Exit()
	}
	global.OkWithMessage(r, "删除成功")
}

// @Tags DictionaryDetail
// @Summary 更新DictionaryDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DictionaryDetail true "更新DictionaryDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /DictionaryDetail/updateDictionaryDetail [put]
func UpdateDictionaryDetail(r *ghttp.Request) {
	var updateRequest request.UpdateDictionaryDetail
	if err := r.Parse(&updateRequest); err != nil {
		global.FailWithMessage(r, err.Error())
	}
	if err := service.UpdateDictionaryDetail(&updateRequest); err != nil {
		global.FailWithMessage(r, fmt.Sprintf("更新失败，err:%v", err))
		r.Exit()
	}
	global.OkWithMessage(r, "更新成功")
}

// @Tags DictionaryDetail
// @Summary 用id查询DictionaryDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DictionaryDetail true "用id查询DictionaryDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /DictionaryDetail/findDictionaryDetail [get]
func FindDictionaryDetail(r *ghttp.Request) {
	var findRequest request.FindDictionaryDetail
	if err := r.Parse(&findRequest); err != nil {
		global.FailWithMessage(r, err.Error())
	}
	dataReturn, err := service.FindDictionaryDetail(&findRequest)
	if err != nil {
		global.FailWithMessage(r, fmt.Sprintf("查询失败，err:%v", err))
		r.Exit()
	}
	global.OkDetailed(r, g.Map{"DictionaryDetail": dataReturn}, "查询成功")
}

// @Tags DictionaryDetail
// @Summary 分页获取DictionaryDetail列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.DictionaryDetailSearch true "分页获取DictionaryDetail列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /DictionaryDetail/getDictionaryDetailList [get]
func GetDictionaryDetailList(r *ghttp.Request) {
	var pageInfoList request.GetDictionaryDetailList
	if err := r.Parse(&pageInfoList); err != nil {
		global.FailWithMessage(r, err.Error())
		r.Exit()
	}
	list, total, err := service.GetDictionaryDetailList(&pageInfoList)
	if err != nil {
		global.FailWithMessage(r, fmt.Sprintf("获取数据失败，err:%v", err))
		r.Exit()
	}
	global.OkWithData(r, response.PageResult{List: list, Total: total, Page: pageInfoList.Page, PageSize: pageInfoList.PageSize})
}
