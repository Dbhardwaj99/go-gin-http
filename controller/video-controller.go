package controller

import (
	"fmt"
	"go-gin-http/entity"
	"go-gin-http/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SongController interface {
	FindAll() []entity.Song
	Save(ctx *gin.Context) entity.Song
	Update(ctx *gin.Context) error
	Delete(ctx *gin.Context)
	ShowAll(ctx *gin.Context)
}

type controller struct {
	service service.SongService
}

func New(service service.SongService) SongController {
	return &controller{
		service: service,
	}
}

func (c controller) FindAll() []entity.Song {
	return c.service.FindAll()
}

func (c controller) Save(ctx *gin.Context) entity.Song {
	var song entity.Song
	err := ctx.ShouldBindJSON(&song)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return song
	}
	c.service.Save(song)
	ctx.JSON(http.StatusOK, song)
	return song
}

func (c controller) ShowAll(ctx *gin.Context) {
	fmt.Println("Executing ShowAll() function")
	songs := c.service.FindAll()
	fmt.Printf("Found %d songs\n", len(songs))
	data := gin.H{
		"title": "Song Page",
		"songs": songs,
	}
	ctx.HTML(http.StatusOK, "index.html", data)
}

func (c controller) Update(ctx *gin.Context) error {
	var song entity.Song
	err := ctx.ShouldBindJSON(&song)
	if err != nil {
		return err
	}

	id, err := strconv.ParseUint(ctx.Params("id"), 0, 0)
	if err != nil {
		return err
	}
	song.ID = id
	c.service.Update(song)
	return nil
}

func (c controller) Delete(ctx *gin.Context) error {
	var song entity.Song
	id, err := strconv.ParseUint(ctx.Params("id"), 0, 0)
	if err != nil {
		return err
	}
	song.ID = id
	c.service.Delete(song)
	return nil
}
