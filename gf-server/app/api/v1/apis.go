package v1

import (
	"fmt"
	"gf-server/app/api/request"
	"gf-server/app/api/response"
	"gf-server/app/service"
	"gf-server/library/global"

	"github.com/gogf/gf/net/ghttp"
)

// @Tags SysApi
// @Summary 创建基础api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysApi true "创建api"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/createApi [post]
func CreateApi(r *ghttp.Request) {
	var C request.CreateApi
	if err := r.Parse(&C); err != nil {
		global.FailWithMessage(r, err.Error())
		r.Exit()
	}
	if err := service.CreateApi(&C); err != nil {
		global.FailWithMessage(r, fmt.Sprintf("创建失败，err:%v", err))
		r.Exit()
	}
	global.OkWithMessage(r, "创建成功")
}

// @Tags SysApi
// @Summary 创建基础api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysApi true "创建api"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/updateApi [post]
func UpdateApi(r *ghttp.Request) {
	var U request.UpdateApi
	if err := r.Parse(&U); err != nil {
		global.FailWithMessage(r, fmt.Sprintf("删除失败，err:%v", err))
		r.Exit()
	}
	if err := service.UpdateApi(&U); err != nil {
		global.FailWithMessage(r, fmt.Sprintf("修改数据失败，%v", err))
		r.Exit()
	}
	global.OkWithMessage(r, "修改数据成功")
}

// @Tags SysApi
// @Summary 删除指定api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysApi true "删除api"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/deleteApi [post]
func DeleteApi(r *ghttp.Request) {
	var D request.DeleteApi
	if err := r.Parse(&D); err != nil {
		global.FailWithMessage(r, err.Error())
		r.Exit()
	}
	if err := service.DeleteApi(&D); err != nil {
		global.FailWithMessage(r, fmt.Sprintf("删除失败，err:%v", err))
		r.Exit()
	}
	global.OkWithMessage(r, "删除成功")
}

// @Tags SysApi
// @Summary 根据id获取api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "根据id获取api"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/getApiById [post]
func GetApiById(r *ghttp.Request) {
	var G request.GetApiById
	if err := r.Parse(&G); err != nil {
		global.FailWithMessage(r, err.Error())
		r.Exit()
	}
	apiReturn, err := service.GetApiById(&G)
	if err != nil {
		global.FailWithMessage(r, fmt.Sprintf("获取失败, err:%v", err))
		r.Exit()
	}
	global.OkDetailed(r, response.ApiResponse{Api: apiReturn}, "获取成功")
}

// @Tags SysApi
// @Summary 获取所有的Api 不分页
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/getAllApis [post]
func GetAllApis(r *ghttp.Request) {
	apisReturn, err := service.GetAllApis()
	if err != nil {
		global.FailWithMessage(r, "获取失败")
		r.Exit()
	}
	global.OkDetailed(r, response.ApiListResponse{Apis: apisReturn}, "获取成功")
}

// 条件搜索后端看此api

// @Tags SysApi
// @Summary 分页获取API列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SearchApiParams true "分页获取API列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/getApiList [post]
func GetApiList(r *ghttp.Request) {
	var sp request.GetApiList // 此结构体仅本方法使用
	if err := r.Parse(&sp); err != nil{
		global.FailWithMessage(r, err.Error())
		r.Exit()
	}
	list, total, err := service.GetApiInfoList(&sp)
	if err != nil{
		global.FailWithMessage(r, fmt.Sprintf("获取数据失败，%v", err))
		r.Exit()
	}
	global.OkWithData(r, response.PageResult{
		List:     list,
		Total:    total,
		Page:     sp.Page,
		PageSize: sp.PageSize,
	})
}
