package v1

import (
	"gf-server/app/api/request"
	"gf-server/app/service"
	"gf-server/library/global"
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
	var change request.ChangePasswordRequest
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
}

// @Tags Admins
// @Summary 删除用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "删除用户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /user/deleteUser [delete]
func DeleteUser(r *ghttp.Request) {
}
