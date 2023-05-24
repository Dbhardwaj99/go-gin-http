package service

import "go-gin-http/entity"

type SongService interface {
	Save(entity.Song) entity.Song
	FindAll() []entity.Song
}

type songService struct {
	songs []entity.Song
}

func New() SongService {
	return &songService{}
}

func (service *songService) Save(song entity.Song) entity.Song {
	service.songs = append(service.songs, song)
	return song
}

func (service *songService) FindAll() []entity.Song {
	return service.songs
}
