package controller

import (
	"go-gin-http/entity"
	"go-gin-http/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) entity.Video
}

type controller struct {
	service service.VideoService
}

func New(service service.VideoService) VideoController {
	return controller{
		service: service,
	}
}

func (c controller) FindAll() []entity.Video {
	return c.service.FindAll()
}

func (c controller) Save(ctx *gin.Context) entity.Video {
	var video entity.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return video
	}
	c.service.Save(video)
	ctx.JSON(http.StatusOK, video)
	return video
}
