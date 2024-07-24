package repository

import (
	"backend_fitfit_app/model"
	"log"

	"gorm.io/gorm"
)

type musicHistoryRepo struct {
	db *gorm.DB
}

func NewMusicHistoryRepository() musicHistoryRepository {
	db, err := NewDatabaseConnection()
	if err != nil {
		return nil
	}
	return musicHistoryRepo{db: db}
}

type musicHistoryRepository interface {
	FindAllMusicHistory() ([]model.MusicHistory, error)
	FindAllMusicHistoryByEid(id int) ([]model.MusicHistory, error)
	AddMusicHistory(musichistory model.MusicHistory) int64
}

func (musicHistoryRepo) FindAllMusicHistory() ([]model.MusicHistory, error) {
	musicHistory := []model.MusicHistory{}
	result := db.Find(&musicHistory)
	if result.Error != nil {
		return nil, result.Error
	}
	return musicHistory, nil
}

func (musicHistoryRepo) FindAllMusicHistoryByEid(id int) ([]model.MusicHistory, error) {
	musichis := []model.MusicHistory{}
	result := db.Where("eid = ?", id).Find(&musichis)
	if result.Error != nil {
		return nil, result.Error
	}
	return musichis, nil
}

func (musicHistoryRepo) AddMusicHistory(musichistory model.MusicHistory) int64 {
	result := db.Create(&musichistory)
	if result.RowsAffected > 0 {
		log.Printf("Add MusicHistory complete\nAffected row : %v", result.RowsAffected)
		return int64(musichistory.ID)
	} else {
		log.Printf("Add MusicHistory failed %v", result.RowsAffected)
		return 0
	}
}
