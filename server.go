package main

import (
	"go-gin-http/controller"
	"go-gin-http/middlewares"
	"go-gin-http/service"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	setupLogOutput()
	server := gin.New()

	server.Use(gin.Recovery(), middlewares.Logger(), middlewares.Auth())

	protected := server.Group("")
	protected.Use(middlewares.Auth(), gindump.Dump())
	{
		protected.GET("/videos", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, videoController.FindAll())
		})

		protected.POST("/videos", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, videoController.Save(ctx))
		})
	}

	server.Run(":8080")
}
