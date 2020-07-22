package v1

import (
	"fmt"
	"server/app/api/request"
	"server/app/api/response"
	"server/app/service"
	"server/library/global"

	"github.com/gogf/gf/net/ghttp"
)

// @Tags Operations
// @Summary 创建Operation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.CreateOperation true "创建Operation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /operation/createSysOperationRecord [post]
func CreateOperation(r *ghttp.Request) {
	var c request.CreateOperation
	if err := r.Parse(&c); err != nil {
		global.FailWithMessage(r, err.Error())
		r.Exit()
	}
	if err := service.CreateOperation(&c); err != nil {
		global.FailWithMessage(r, fmt.Sprintf("创建失败，%v", err))
		r.Exit()
	}
	global.OkWithMessage(r, "创建成功")
}

// @Tags Operations
// @Summary 删除Operation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.DeleteOperation true "删除Operation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /operation/deleteSysOperationRecord [delete]
func DeleteOperation(r *ghttp.Request) {
	var d request.DeleteOperation
	if err := r.Parse(&d); err != nil {
		global.FailWithMessage(r, err.Error())
		r.Exit()
	}
	if err := service.DeleteOperation(&d); err != nil {
		global.FailWithMessage(r, fmt.Sprintf("删除失败，%v", err))
		r.Exit()
	}
	global.OkWithMessage(r, "删除成功")
}

// @Tags Operations
// @Summary 批量删除Operation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.DeleteOperations true "批量删除Operations"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /operation/deleteOperations [delete]
func DeleteOperations(r *ghttp.Request) {
	var d request.DeleteOperations
	if err := r.Parse(&d); err != nil {
		global.FailWithMessage(r, err.Error())
		r.Exit()
	}
	if err := service.DeleteOperations(&d); err != nil {
		global.FailWithMessage(r, fmt.Sprintf("批量删除失败，%v", err))
		r.Exit()
	}
	global.OkWithMessage(r, "批量删除成功")
}

// @Tags Operations
// @Summary 更新Operation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.UpdateOperation true "更新Operation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /operation/updateOperation [put]
func UpdateOperation(r *ghttp.Request) {
	var u request.UpdateOperation
	if err := r.Parse(&u); err != nil {
		global.FailWithMessage(r, err.Error())
		r.Exit()
	}
	if err := service.UpdateOperation(&u); err != nil {
		global.FailWithMessage(r, fmt.Sprintf("更新失败，%v", err))
		r.Exit()
	}
	global.OkWithMessage(r, "更新成功")
}

// @Tags Operations
// @Summary 用id查询Operation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysOperationRecord true "用id查询Operation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /operation/findOperation [get]
func FindOperation(r *ghttp.Request) {
	var u request.UpdateOperation
	if err := r.Parse(&u); err != nil {
		global.FailWithMessage(r, err.Error())
		r.Exit()
	}
	if err := service.UpdateOperation(&u); err != nil {
		global.FailWithMessage(r, fmt.Sprintf("获取失败，%v", err))
		r.Exit()
	}
	global.OkWithMessage(r, "获取成功")
}

// @Tags Operations
// @Summary 分页获取Operation列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetOperationList true "分页获取Operation列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /operation/getOperationList [get]
func GetOperationList(r *ghttp.Request) {
	var g request.GetOperationList
	if err := r.Parse(&g); err != nil {
		global.FailWithMessage(r, err.Error())
		r.Exit()
	}
	list, total, err := service.GetOperationList(&g)
	if err != nil {
		global.FailWithMessage(r, fmt.Sprintf("获取数据失败，%v", err))
		r.Exit()
	}
	global.OkWithData(r, response.PageResult{
		List:     list,
		Total:    total,
		Page:     g.Page,
		PageSize: g.PageSize,
	})
}
