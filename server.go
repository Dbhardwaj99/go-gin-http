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
	songService    service.SongService       = service.New()
	songController controller.SongController = controller.New(songService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	setupLogOutput()
	server := gin.New()

	server.Static("/css", "./templates/css")
	server.LoadHTMLGlob("templates/*.html")

	server.Use(gin.Recovery(), middlewares.Logger())

	apiRoutes := server.Group("/api")
	apiRoutes.Use(middlewares.Auth(), gindump.Dump())
	{
		apiRoutes.GET("/songs-api", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, songController.FindAll())
		})

		apiRoutes.POST("/songs-api", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, songController.Save(ctx))
		})
	}

	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/songs", songController.ShowAll)
	}

	server.Run(":8080")
}
