package router

import (
	v1 "server/app/api/v1"

	"github.com/gogf/gf/frame/g"
)

// InitFileRouter 注册文件上传下载功能路由
func InitFileRouter() {
	FileRouter := g.Server().Group("fileUploadAdDownload")
	{
		FileRouter.POST("upload", v1.UploadFile)                                 // 上传文件
		FileRouter.POST("getFileList", v1.GetFileList)                           // 获取上传文件列表
		FileRouter.POST("deleteFile", v1.DeleteFile)                             // 删除指定文件
		FileRouter.POST("breakpointContinue", v1.BreakpointContinue)             // 断点续传
		FileRouter.GET("findFile", v1.FindFile)                                  // 查询当前文件成功的切片
		FileRouter.POST("breakpointContinueFinish", v1.BreakpointContinueFinish) // 查询当前文件成功的切片
		FileRouter.POST("removeChunk", v1.RemoveChunk)                           // 查询当前文件成功的切片
	}
}
