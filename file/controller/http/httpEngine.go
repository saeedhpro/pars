package http

import (
	"file/controller"
	"github.com/gin-gonic/gin"
)

func Run(port string) {
	engine := gin.Default()
	engine.Use(gin.Recovery())

	files := engine.Group("files")

	fileApi := NewFileApi(controller.UploadFile, controller.GetFile, controller.GetFiles)

	files.POST("/", fileApi.UploadFile)
	files.POST("/file", fileApi.GetFile)
	files.POST("/files", fileApi.GetFiles)
	_ = engine.Run(port)
}
