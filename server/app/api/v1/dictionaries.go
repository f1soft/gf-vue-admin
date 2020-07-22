package v1

import (
	"fmt"
	"gf-server/app/api/request"
	"gf-server/app/api/response"
	"gf-server/app/service"
	"gf-server/library/global"

	"github.com/gogf/gf/frame/g"

	"github.com/gogf/gf/net/ghttp"
)

// @Tags Dictionary
// @Summary 创建Dictionary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Dictionary true "创建Dictionary"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysDictionary/createDictionary [post]
func CreateDictionary(r *ghttp.Request) {
	var createInfo request.CreateDictionary
	if err := r.Parse(&createInfo); err != nil {
		global.FailWithMessage(r, err.Error())
	}
	if err := service.CreateDictionary(&createInfo); err != nil {
		global.FailWithMessage(r, fmt.Sprintf("创建失败，err:%v", err))
		r.Exit()
	}
	global.OkWithMessage(r, "创建成功")
}

// @Tags Dictionary
// @Summary 删除Dictionary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Dictionary true "删除Dictionary"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysDictionary/deleteDictionary [delete]
func DeleteDictionary(r *ghttp.Request) {
	var deleteInfo request.DeleteDictionary
	if err := r.Parse(&deleteInfo); err != nil {
		global.FailWithMessage(r, err.Error())
	}
	if err := service.DeleteDictionary(&deleteInfo); err != nil {
		global.FailWithMessage(r, fmt.Sprintf("删除失败，err:%v", err))
		r.Exit()
	}
	global.OkWithMessage(r, "删除成功")
}

// @Tags Dictionary
// @Summary 更新Dictionary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Dictionary true "更新Dictionary"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sysDictionary/updateDictionary [put]
func UpdateDictionary(r *ghttp.Request) {
	var updateInfo request.UpdateDictionary
	if err := r.Parse(&updateInfo); err != nil {
		global.FailWithMessage(r, err.Error())
	}
	if err := service.UpdateDictionary(&updateInfo); err != nil {
		global.FailWithMessage(r, fmt.Sprintf("更新失败，err:%v", err))
		r.Exit()
	}
	global.OkWithMessage(r, "更新成功")
}

// @Tags Dictionary
// @Summary 用id查询Dictionary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Dictionary true "用id查询Dictionary"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sysDictionary/findDictionary [get]
func FindDictionary(r *ghttp.Request) {
	var findInfo request.FindDictionary
	if err := r.Parse(&findInfo); err != nil {
		global.FailWithMessage(r, err.Error())
	}
	dictionary, err := service.FindDictionary(&findInfo)
	if err != nil {
		global.FailWithMessage(r, fmt.Sprintf("查询失败，err:%v", err))
		r.Exit()
	}
	global.OkWithData(r, g.Map{"resysDictionary": dictionary})
}

// @Tags Dictionary
// @Summary 分页获取Dictionary列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.DictionarySearch true "分页获取SysDictionary列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysDictionary/getSysDictionaryList [get]
func GetDictionaryList(r *ghttp.Request) {
	var pageInfo request.DictionaryInfoList
	if err := r.Parse(&pageInfo); err != nil {
		global.FailWithMessage(r, err.Error())
	}
	list, total, err := service.GetDictionaryInfoList(&pageInfo)
	if err != nil {
		global.FailWithMessage(r, fmt.Sprintf("获取数据失败，err:%v", err))
		r.Exit()
	}
	global.OkWithData(r, response.PageResult{List: list, Total: total, Page: pageInfo.Page, PageSize: pageInfo.PageSize})
}
