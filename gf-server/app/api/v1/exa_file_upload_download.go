package v1

import (
	"fmt"
	"gf-server/app/api/request"
	"gf-server/app/api/response"
	"gf-server/app/model/files"
	"gf-server/app/service"
	"gf-server/library/global"
	"gf-server/library/utils"
	"strings"

	"github.com/gogf/gf/net/ghttp"
)

// @Tags Files
// @Summary 上传文件示例
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true "上传文件示例"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"上传成功"}"
// @Router /fileUploadAndDownload/upload [post]
func UploadFile(r *ghttp.Request) {
	var (
		filePath string
		key      string
		file     files.Entity
	)
	noSave := r.GetQuery("noSave", "0")
	_, header, err := r.Request.FormFile("file")
	if err != nil {
		global.FailWithMessage(r, fmt.Sprintf("上传文件失败，%v", err))
		r.Exit()
	}
	filePath, key, err = utils.Upload(header)
	if err != nil {
		global.FailWithMessage(r, fmt.Sprintf("接收返回值失败，%v", err))
		r.Exit()
	}
	file.Url = filePath
	file.Name = header.Filename
	file.Url = filePath
	s := strings.Split(file.Name, ".")
	file.Tag = s[len(s)-1]
	file.Key = key
	file.Key = key
	if noSave == "0" {
		if err = service.UploadFile(&file); err != nil {
			global.FailWithMessage(r, fmt.Sprintf("修改数据库链接失败，%v", err))
			r.Exit()
		}
		global.OkDetailed(r, response.Files{File: &file}, "上传成功")
	}
}

// @Tags Files
// @Summary 删除文件
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body model.ExaFileUploadAndDownload true "传入文件里面id即可"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /fileUploadAndDownload/deleteFile [post]
func DeleteFile(r *ghttp.Request) {
	var d request.DeleteFile
	if err := r.Parse(&d); err != nil {
		global.FailWithMessage(r, err.Error())
		r.Exit()
	}
	if err := service.DeleteFile(&d); err != nil {
		global.FailWithMessage(r, fmt.Sprintf("删除失败，%v", err))
		r.Exit()
	}
	global.OkWithMessage(r, "删除成功")
}

// @Tags Files
// @Summary 分页文件列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取文件户列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /fileUploadAndDownload/getFileList [post]
func GetFileList(r *ghttp.Request) {
	var g request.PageInfo
	if err := r.Parse(&g); err != nil {
		global.FailWithMessage(r, err.Error())
		r.Exit()
	}
	list, total, err := service.GetFileList(&g)
	if err != nil {
		global.FailWithMessage(r, fmt.Sprintf("获取数据失败，%v", err))
		r.Exit()
	}
	global.OkDetailed(r, response.PageResult{List: list, Total: total, Page: g.Page, PageSize: g.PageSize}, "获取成功")
}
