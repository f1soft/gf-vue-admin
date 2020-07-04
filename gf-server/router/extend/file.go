package extend

import (
	"gf-server/global"
)

// InitFileRouter 注册文件上传下载功能路由
func InitFileRouter() {
	FileRouter := global.GFVA_SERVER.Group("fileUploadAndDownload")
	{
		FileRouter.POST("upload", EmptyRequest)                   // 上传文件
		FileRouter.POST("getFileList", EmptyRequest)              // 获取上传文件列表
		FileRouter.POST("deleteFile", EmptyRequest)               // 删除指定文件
		FileRouter.POST("breakpointContinue", EmptyRequest)       // 断点续传
		FileRouter.GET("findFile", EmptyRequest)                  // 查询当前文件成功的切片
		FileRouter.POST("breakpointContinueFinish", EmptyRequest) // 查询当前文件成功的切片
		FileRouter.POST("removeChunk", EmptyRequest)              // 查询当前文件成功的切片
	}
}
