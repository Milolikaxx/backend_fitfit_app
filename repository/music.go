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
	FindAllMusicByMusictype(musicType int) ([]model.Music, error)
	RandomMusicByMusictype(id int) (*model.Music, error)
	FindAllMusicByLevel(bpm int, musicType []int) ([]model.Music, error)
}

func (musicRepo) FindAllMusicByMusictype(musicType int) ([]model.Music, error) {
	music := []model.Music{}
	result := db.Joins("MusicType").Where("music.mtid = ?", musicType).Find(&music)
	if result.Error != nil {
		return nil, result.Error
	}
	return music, nil
}

func (musicRepo) RandomMusicByMusictype(id int) (*model.Music, error) {
	music := model.Music{}
	result := db.Joins("MusicType").Where("music.mtid = ?", id).Order("RAND()").First(&music)
	if result.Error != nil {
		return nil, result.Error
	}
	return &music, nil
}

func (musicRepo) FindAllMusicByLevel(bpm int, musicType []int) ([]model.Music, error) {
	music := []model.Music{}
	result := db.Joins("MusicType").Where("music.bpm <= ?", bpm).Where("music.mtid in (?)", musicType).Find(&music)
	if result.Error != nil {
		return nil, result.Error
	}
	return music, nil
}
