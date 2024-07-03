package repository

import (
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
	// FindById(id int) ([]model.MusicHistory, error)
}
