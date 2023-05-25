package service

import (
	"go-gin-http/entity"
	"go-gin-http/repository"
)

type SongService interface {
	Save(entity.Song) entity.Song
	FindAll() []entity.Song
	Update(song entity.Song)
	Delete(song entity.Song)
}

type songService struct {
	SongRepository repository.SongRepository
}

func New(repo repository.SongRepository) SongService {
	return &songService{
		SongRepository: repo,
	}
}

func (service *songService) Save(song entity.Song) entity.Song {
	service.SongRepository.Save(song)
	return song
}

func (service *songService) FindAll() []entity.Song {
	return service.SongRepository.FindAll()
}

func (service *songService) Update(song entity.Song) {
	service.SongRepository.Update(song)
}

func (service *songService) Delete(song entity.Song) {
	service.SongRepository.Delete(song)
}
