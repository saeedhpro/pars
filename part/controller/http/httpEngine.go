package http

import (
	"github.com/gin-gonic/gin"
	"part/controller"
)

func Run(port string) {
	engine := gin.Default()
	engine.Use(gin.Recovery())

	parts := engine.Group("parts")

	partApi := NewPartApi(controller.GetAutomobileParts, controller.AddFile, controller.GetPartFiles, controller.GetAutomobileFiles, controller.AddAutomobileParts)

	parts.GET("/automobiles/:id/parts", partApi.GetAutomobileParts)
	parts.GET("/automobiles/:id/files", partApi.GetAutomobileFiles)
	parts.POST("/automobile/:id/parts", partApi.AddAutomobileParts)
	parts.POST("/:id/file", partApi.AddFile)
	parts.GET("/:id/files", partApi.GetPartFiles)
	_ = engine.Run(port)
}
