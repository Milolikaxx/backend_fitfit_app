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
	FindByID(key int) (*model.WorkoutProfile, error)
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

func (wpRepo) FindByID(id int) (*model.WorkoutProfile, error) {
	wp := model.WorkoutProfile{}
	result := db.Where("wpid = ?", id).Find(&wp)
	if result.Error != nil {
		return nil, result.Error
	}
	return &wp, nil
}

func (wpRepo) AddWorkProfile(wp model.WorkoutProfile) int64 {
	result := db.Create(&wp)
	if result.RowsAffected > 0 {
		log.Printf("Add workoutProfile complete\nAffected row : %v", result.RowsAffected)
	} else {
		log.Printf("Add workoutProfile failed %v", result.RowsAffected)
	}
	return result.RowsAffected
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
