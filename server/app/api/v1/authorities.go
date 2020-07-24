package v1

import (
	"fmt"
	"server/app/api/request"
	"server/app/api/response"
	"server/app/model/authorities"
	"server/app/service"
	"server/library/global"

	"github.com/gogf/gf/net/ghttp"
)

// @Tags authority
// @Summary 创建角色
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysAuthority true "创建角色"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /authority/createAuthority [post]
func CreateAuthority(r *ghttp.Request) {
	var c request.CreateAuthority
	if err := r.Parse(&c); err != nil {
		global.FailWithMessage(r, err.Error())
		r.Exit()
	}
	authority, err := service.CreateAuthority(&c)
	if err != nil {
		global.FailWithMessage(r, fmt.Sprintf("创建失败，%v", err))
		r.Exit()
	}
	global.OkDetailed(r, response.Authority{Authority: authority}, "创建成功")
}

// @Tags authority
// @Summary 拷贝角色
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body response.SysAuthorityCopyResponse true "拷贝角色"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"拷贝成功"}"
// @Router /authority/copyAuthority [post]
func CopyAuthority(r *ghttp.Request) {
	var copyInfo request.AuthorityCopy
	if err := r.Parse(&copyInfo); err != nil {
		global.FailWithMessage(r, err.Error())
		r.Exit()
	}
	authority, err := service.CopyAuthority(&copyInfo)
	if err != nil {
		global.FailWithMessage(r, fmt.Sprintf("拷贝失败，%v", err))
		r.Exit()
	}
	global.OkWithData(r, response.CopyAuthority{Authority: authority})
}

// @Tags authority
// @Summary 删除角色
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysAuthority true "删除角色"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /authority/deleteAuthority [post]
func DeleteAuthority(r *ghttp.Request) {
	var d request.DeleteAuthority
	if err := r.Parse(&d); err != nil {
		global.FailWithMessage(r, err.Error())
		r.Exit()
	}
	if err := service.DeleteAuthority(&d); err != nil {
		global.FailWithMessage(r, fmt.Sprintf("删除失败，err:%v", err))
		r.Exit()
	}
	global.OkWithMessage(r, "删除成功")
}

// @Tags authority
// @Summary 设置角色资源权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysAuthority true "设置角色资源权限"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"设置成功"}"
// @Router /authority/updateAuthority [post]
func UpdateAuthority(r *ghttp.Request) {
	var u request.UpdateAuthority
	if err := r.Parse(&u); err != nil {
		global.FailWithMessage(r, err.Error())
		r.Exit()
	}
	authority, err := service.UpdateAuthority(&u)
	if err != nil {
		global.FailWithMessage(r, fmt.Sprintf("更改失败，err:%v", err))
		r.Exit()
	}
	global.OkDetailed(r, response.Authority{Authority: authority}, "更改成功")
}

// @Tags authority
// @Summary 分页获取角色列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取用户列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /authority/getAuthorityList [post]
func GetAuthorityList(r *ghttp.Request) {
	var pageInfo request.PageInfo
	if err := r.Parse(&pageInfo); err != nil {
		global.FailWithMessage(r, err.Error())
		r.Exit()
	}
	list, total, err := service.GetAuthorityInfoList(&pageInfo)
	if err != nil {
		global.FailWithMessage(r, fmt.Sprintf("获取数据失败，err:%v", err))
		r.Exit()
	}
	global.OkDetailed(r, response.PageResult{List: list, Total: total, Page: pageInfo.Page, PageSize: pageInfo.PageSize}, "获取成功")
}

// @Tags authority
// @Summary 设置角色资源权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysAuthority true "设置角色资源权限"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"设置成功"}"
// @Router /authority/setDataAuthority [post]
func SetDataAuthority(r *ghttp.Request) {
	var auth authorities.Authorities
	if err := r.Parse(&auth); err != nil {
		global.FailWithMessage(r, err.Error())
		r.Exit()
	}
	err := service.SetDataAuthority(&auth)
	if err != nil {
		global.FailWithMessage(r, fmt.Sprintf("设置关联失败，%v", err))
		r.Exit()
	}
	global.Ok(r)
}
