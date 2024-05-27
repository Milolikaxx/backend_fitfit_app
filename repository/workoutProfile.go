package repository

import (
	"backend_fitfit_app/model"
	"log"

	"gorm.io/gorm"
)

type wpRepo struct {
	db *gorm.DB
}

func NewWpRepository() wpRepository {
	db, err := NewDatabaseConnection()
	if err != nil {
		return nil
	}
	return wpRepo{db: db}
}

type wpRepository interface {
	FindAll() ([]model.WorkoutProfile, error)
	FindByWPID(key int) (*model.WorkoutProfile, error)
	FindListByWPID(key int) ([]model.WorkoutProfile, error)
	FindByUID(key int) ([]model.WorkoutProfile, error)
	AddWorkProfile(model.WorkoutProfile) int64
	UpdateWorkProfile(wp model.WorkoutProfile, id int) int64
}

func (u wpRepo) FindAll() ([]model.WorkoutProfile, error) {
	wps := []model.WorkoutProfile{}
	result := u.db.Find(&wps)
	if result.Error != nil {
		return nil, result.Error
	}
	return wps, nil
}

func (wpRepo) FindByWPID(id int) (*model.WorkoutProfile, error) {
	wp := model.WorkoutProfile{}
	result := db.Where("wpid = ?", id).Find(&wp)
	if result.Error != nil {
		return nil, result.Error
	}
	return &wp, nil
}
func (wpRepo) FindListByWPID(id int) ([]model.WorkoutProfile, error) {
	wps := []model.WorkoutProfile{}
	result := db.Joins("WorkoutMusictype").Joins("WorkoutMusictype.MusicType").Where("workout_profile.wpid = ?", id).Find(&wps)
	if result.Error != nil {
		return nil, result.Error
	}
	return wps, nil
}

func (wpRepo) FindByUID(uid int) ([]model.WorkoutProfile, error) {
	wps := []model.WorkoutProfile{}
	result := db.Where("uid = ?", uid).Find(&wps)
	if result.Error != nil {
		return nil, result.Error
	}
	return wps, nil
}

func (wpRepo) AddWorkProfile(wp model.WorkoutProfile) int64 {
	result := db.Create(&wp)
	if result.RowsAffected > 0 {
		log.Printf("Add workoutProfile complete\nAffected row : %v", result.RowsAffected)
		return int64(wp.Wpid)
	} else {
		log.Printf("Add workoutProfile failed %v", result.RowsAffected)
		return 0
	}
}

func (wpRepo) UpdateWorkProfile(wp model.WorkoutProfile, id int) int64 {
	result := db.Model(&model.WorkoutProfile{}).Where("wpid = ?", id).Updates(&wp)
	if result.RowsAffected > 0 {
		log.Printf("Update workoutProfile complete\nAffected row : %v", result.RowsAffected)
	} else {
		log.Printf("Update workoutProfile failed\nAffected row : %v", result.RowsAffected)
	}
	return result.RowsAffected
}
