package repository

import (
	"backend_fitfit_app/model"
	"log"

	"gorm.io/gorm"
)

type wpmusicTypeRepo struct {
	db *gorm.DB
}

func NewWpMusicTypeRepository() wpMusicTypeRepository {
	db, err := NewDatabaseConnection()
	if err != nil {
		return nil
	}
	return wpmusicTypeRepo{db: db}
}

type wpMusicTypeRepository interface {
	FindByWPID(key int) ([]model.WorkoutMusictype, error)
	AddWpMusicType(model.WorkoutMusictype) int64
	UpdateWpMusicType(model.WorkoutMusictype, int) int64
}

func (wpmusicTypeRepo) FindByWPID(id int) ([]model.WorkoutMusictype, error) {
    wpMusicType := []model.WorkoutMusictype{}
    result := db.Preload("MusicType").Where("wpid = ?", id).Find(&wpMusicType)
    if result.Error != nil {
        return nil, result.Error
    }
    return wpMusicType, nil
}

// func (wpmusicTypeRepo) FindByWPID(id int) ([]model.WorkoutMusictype, error) {
// 	wpMusicType := []model.WorkoutMusictype{}
// 	result := db.Where("wpid = ?", id).Find(&wpMusicType)
// 	if result.Error != nil {
// 		return nil, result.Error
// 	}
// 	return wpMusicType, nil
// }

func (wpmusicTypeRepo) AddWpMusicType(wpmt model.WorkoutMusictype) int64 {
	result := db.Create(&wpmt)
	if result.RowsAffected > 0 {
		log.Printf("Add workoutProfile complete\nAffected row : %v", result.RowsAffected)
	} else {
		log.Printf("Add workoutProfile failed %v", result.RowsAffected)
	}
	return result.RowsAffected
}

func (wpmusicTypeRepo) UpdateWpMusicType(wpmt model.WorkoutMusictype, id int) int64 {
	result := db.Model(&model.WorkoutMusictype{}).Where("wpid = ?", id).Updates(&wpmt)
	if result.RowsAffected > 0 {
		log.Printf("Update workoutProfile complete\nAffected row : %v", result.RowsAffected)
	} else {
		log.Printf("Update workoutProfile failed\nAffected row : %v", result.RowsAffected)
	}
	return result.RowsAffected
}
