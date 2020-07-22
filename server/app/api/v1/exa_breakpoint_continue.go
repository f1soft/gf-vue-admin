package v1

import (
	"github.com/gogf/gf/net/ghttp"
)

// @Tags ExaFileUploadAndDownload
// @Summary 断点续传到服务器
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true "an example for breakpoint resume, 断点续传示例"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"上传成功"}"
// @Router /fileUploadAndDownload/breakpointContinue [post]
func BreakpointContinue(r *ghttp.Request) {
}

// @Tags ExaFileUploadAndDownload
// @Summary 查找文件
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true "Find the file, 查找文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查找成功"}"
// @Router /fileUploadAndDownload/findFile [post]
func FindFile(r *ghttp.Request) {
}

// @Tags ExaFileUploadAndDownload
// @Summary 查找文件
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true "上传文件完成"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"file uploaded, 文件创建成功"}"
// @Router /fileUploadAndDownload/findFile [post]
func BreakpointContinueFinish(r *ghttp.Request) {
}

// @Tags ExaFileUploadAndDownload
// @Summary 删除切片
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true "删除缓存切片"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查找成功"}"
// @Router /fileUploadAndDownload/removeChunk [post]
func RemoveChunk(r *ghttp.Request) {
	//fileMd5 := r.GetString("fileMd5")
	//fileName := r.GetString("fileName")
	//filePath := r.GetString("filePath")
	//err := utils.RemoveChunk(fileMd5)
	//err = service.DeleteFileChunk(fileMd5, fileName, filePath)
	//if err != nil {
	//	global.FailWithDetailed(r, global.ERROR, g.Map{"filePath": filePath}, fmt.Sprintf("缓存切片删除失败：%v", err))
	//	r.Exit()
	//}
	//global.OkDetailed(r, g.Map{"filePath": filePath}, "缓存切片删除成功")
}
