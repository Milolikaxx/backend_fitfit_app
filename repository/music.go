package repository

import (
	"backend_fitfit_app/model"

	"gorm.io/gorm"
)

type musicRepo struct {
	db *gorm.DB
}

func NewMusicRepository() musicRepository {
	db, err := NewDatabaseConnection()
	if err != nil {
		return nil
	}
	return musicRepo{db: db}
}

type musicRepository interface {
	FindAllMusicByMusictype(key int) ([]model.Music, error)
}

func (musicRepo) FindAllMusicByMusictype(id int) ([]model.Music, error) {
	music := []model.Music{}
	result := db.Joins("MusicType").Where("music.mtid = ?", id).Find(&music)
	if result.Error != nil {
		return nil, result.Error
	}
	return music, nil
}
