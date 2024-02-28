package repository

import (
	"backend_fitfit_app/model"

	"gorm.io/gorm"
)

type musicTypeRepo struct {
	db *gorm.DB
}

func NewMusicTypeRepository() musicTypeRepository {
	db, err := NewDatabaseConnection()
	if err != nil {
		return nil
	}
	return musicTypeRepo{db: db}
}

type musicTypeRepository interface {
	FindAll() ([]model.MusicType, error)
	FindByID(key int) (*model.MusicType, error)
}

func (musicTypeRepo) FindAll() ([]model.MusicType, error) {
	music_types := []model.MusicType{}
	result := db.Find(&music_types)
	if result.Error != nil {
		return nil, result.Error
	}
	return music_types, nil
}

func (musicTypeRepo) FindByID(id int) (*model.MusicType, error) {
	music_type := model.MusicType{}
	result := db.Where("mtid = ?", id).Find(&music_type)
	if result.Error != nil {
		return nil, result.Error
	}
	return &music_type, nil
}
