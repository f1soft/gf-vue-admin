package v1

import (
	"fmt"
	"server/app/api/request"
	"server/app/api/response"
	"server/app/service"
	"server/library/global"

	"github.com/gogf/gf/net/ghttp"
)

// @Tags Apis
// @Summary 创建基础api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysApi true "创建api"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/createApi [post]
func CreateApi(r *ghttp.Request) {
	var c request.CreateApi
	if err := r.Parse(&c); err != nil {
		global.FailWithMessage(r, err.Error())
		r.Exit()
	}
	if err := service.CreateApi(&c); err != nil {
		global.FailWithMessage(r, fmt.Sprintf("创建失败，err:%v", err))
		r.Exit()
	}
	global.OkWithMessage(r, "创建成功")
}

// @Tags Apis
// @Summary 更新基础api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysApi true "更新api"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /api/updateApi [post]
func UpdateApi(r *ghttp.Request) {
	var u request.UpdateApi
	if err := r.Parse(&u); err != nil {
		global.FailWithMessage(r, err.Error())
		r.Exit()
	}
	if err := service.UpdateApi(&u); err != nil {
		global.FailWithMessage(r, fmt.Sprintf("更新失败，%v", err))
		r.Exit()
	}
	global.OkWithMessage(r, "更新成功")
}

// @Tags Apis
// @Summary 删除指定api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysApi true "删除api"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /api/deleteApi [post]
func DeleteApi(r *ghttp.Request) {
	var d request.DeleteApi
	if err := r.Parse(&d); err != nil {
		global.FailWithMessage(r, err.Error())
		r.Exit()
	}
	if err := service.DeleteApi(&d); err != nil {
		global.FailWithMessage(r, fmt.Sprintf("删除失败，err:%v", err))
		r.Exit()
	}
	global.OkWithMessage(r, "删除成功")
}

// @Tags Apis
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

// @Tags Apis
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

// @Tags Apis
// @Summary 分页获取API列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SearchApiParams true "分页获取API列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/getApiList [post]
func GetApiList(r *ghttp.Request) {
	var pageInfo request.GetApiList // 此结构体仅本方法使用
	if err := r.Parse(&pageInfo); err != nil {
		global.FailWithMessage(r, err.Error())
		r.Exit()
	}
	list, total, err := service.GetApiInfoList(&pageInfo)
	if err != nil {
		global.FailWithMessage(r, fmt.Sprintf("获取数据失败，%v", err))
		r.Exit()
	}
	global.OkDetailed(r, response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功")
}
