package v1

import (
	"fmt"
	"gf-server/app/api/request"
	"gf-server/app/api/response"
	"gf-server/app/model/admins"
	"gf-server/app/service"
	"gf-server/library/global"
	"gf-server/library/utils"
	"mime/multipart"

	"github.com/gogf/gf/util/gconv"

	"github.com/gogf/gf/net/ghttp"
)

// @Tags Admins
// @Summary 用户修改密码
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body request.ChangePasswordStruct true "用户修改密码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /user/changePassword [put]
func ChangePassword(r *ghttp.Request) {
	var change request.ChangePassword
	if err := r.Parse(&change); err != nil {
		global.FailWithMessage(r, err.Error())
		r.Exit()
	}
	if err := service.ChangePassword(&change); err != nil {
		global.OkWithMessage(r, "修改失败，请检查用户名密码")
		r.Exit()
	}
	global.OkWithMessage(r, "修改成功")
}

// @Tags Admins
// @Summary 用户上传头像
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param headerImg formData file true "用户上传头像"
// @Param username formData string true "用户上传头像"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"上传成功"}"
// @Router /user/uploadHeaderImg [post]
func UploadHeaderImg(r *ghttp.Request) {
	var (
		err      error
		filePath string
		header   *multipart.FileHeader
		admin    *admins.Entity
	)
	userUuid := gconv.String(r.GetParam("uuid"))
	if _, header, err = r.Request.FormFile("headerImg"); err != nil {
		global.FailWithMessage(r, fmt.Sprintf("上传文件失败，%v", err))
	}
	if filePath, _, err = utils.Upload(header); err != nil {
		global.FailWithMessage(r, fmt.Sprintf("接收返回值失败，%v", err))
	}
	// 修改数据库后得到修改后的user并且返回供前端使用
	admin, err = service.UploadHeaderImg(userUuid, filePath)
	if err != nil {
		global.FailWithMessage(r, fmt.Sprintf("修改数据库链接失败，%v", err))
	} else {
		global.OkDetailed(r, response.AdminResponse{User: admin}, "上传成功")
	}
}

// @Tags Admins
// @Summary 分页获取用户列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取用户列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /user/getUserList [post]
func GetUserList(r *ghttp.Request) {
	var pageInfo request.PageInfo
	if err := r.Parse(&pageInfo); err != nil {
		global.FailWithMessage(r, err.Error())
		r.Exit()
	}
	list, total, err := service.GetAdminInfoList(&pageInfo)
	if err != nil {
		global.FailWithMessage(r, fmt.Sprintf("获取数据失败，%v", err))
		r.Exit()
	}
	global.OkDetailed(r, response.PageResult{List: list, Total: total, Page: pageInfo.Page, PageSize: pageInfo.PageSize}, "获取成功")
}

// @Tags Admins
// @Summary 设置用户权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SetUserAuth true "设置用户权限"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /user/setUserAuthority [post]
func SetUserAuthority(r *ghttp.Request) {
	var set request.SetAdminAuthority
	if err := r.Parse(&set); err != nil {
		global.FailWithMessage(r, err.Error())
		r.Exit()
	}
	if err := service.SetUserAuthority(&set); err != nil {
		global.FailWithMessage(r, fmt.Sprintf("修改失败，%v", err))
		r.Exit()
	}
	global.OkWithMessage(r, "修改成功")

}

// @Tags Admins
// @Summary 删除用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "删除用户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /user/deleteUser [delete]
func DeleteAdmin(r *ghttp.Request) {
	var D request.DeleteAdmin
	if err := r.Parse(&D); err != nil {
		global.FailWithMessage(r, err.Error())
		r.Exit()
	}
	if err := service.DeleteAdmin(&D); err != nil {
		global.FailWithMessage(r, fmt.Sprintf("删除成功, err:%v", err))
	}
	global.OkWithMessage(r, "删除成功")
}
