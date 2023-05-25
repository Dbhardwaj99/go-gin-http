package repository

import (
	"github.com/jinzhu/gorm"
	// "github.com/jinzhu/gorm/dialects/sqlite"
	"go-gin-http/entity"
)

type SongRepository interface {
	Save(song entity.Song)
	Update(song entity.Song)
	Delete(song entity.Song)
	FindAll() []entity.Song
	CloseDB()
}

type database struct {
	connection *gorm.DB
}

func NewSongRepository(db *gorm.DB) SongRepository {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&entity.Song{}, &entity.Person{})
	return &database{
		connection: db,
	}
}

func (db *database) CloseDB() {
	err := db.connection.Close()
	if err != nil {
		panic("failed to connect database")
	}
}

func (db *database) Save(song entity.Song) {
	db.connection.Create(&song)
}
func (db *database) Update(song entity.Song) {
	db.connection.Save(&song)
}
func (db *database) Delete(song entity.Song) {
	db.connection.Delete(&song)
}

func (db *database) FindAll() []entity.Song {
	var songs []entity.Song
	db.connection.Set("gorm:auto_preload", true).Find(&songs)
	return songs
}
