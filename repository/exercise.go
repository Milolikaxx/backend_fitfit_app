package repository

import (
	"backend_fitfit_app/model"
	"log"

	"gorm.io/gorm"
)

type exerciseRepo struct {
	db *gorm.DB
}

func NewExerciseRepository() exerciseRepository {
	db, err := NewDatabaseConnection()
	if err != nil {
		return nil
	}
	return exerciseRepo{db: db}
}

type exerciseRepository interface {
	FindAll() ([]model.Exercise, error)
	FindByID(id int) ([]model.Exercise, error)
	AddExercise(exercise model.Exercise) int64
	UpdateExercise(exercise model.Exercise, id int) int64
	FindExerciseByDay(keyword string) ([]model.Exercise, error)
	FindExerciseLast7Days() ([]model.Exercise, error)
	FindExercisesLast12Months() ([]model.Exercise, error)
}

func (exerciseRepo) FindAll() ([]model.Exercise, error) {
	exercise := []model.Exercise{}
	result := db.Find(&exercise)
	if result.Error != nil {
		return nil, result.Error
	}
	return exercise, nil
}

func (exerciseRepo) FindByID(id int) ([]model.Exercise, error) {
	exercise := []model.Exercise{}
	result := db.Where("uid = ?", id).Find(&exercise)
	if result.Error != nil {
		return nil, result.Error
	}
	return exercise, nil
}

func (exerciseRepo) AddExercise(exercise model.Exercise) int64 {
	result := db.Create(&exercise)
	if result.RowsAffected > 0 {
		log.Printf("Add Exercise complete\nAffected row : %v", result.RowsAffected)
		return int64(exercise.Eid)
	} else {
		log.Printf("Add Exercise failed %v", result.RowsAffected)
		return 0
	}
}

func (exerciseRepo) UpdateExercise(exercise model.Exercise, id int) int64 {
	result := db.Model(&model.Exercise{}).Where("eid = ?", id).Updates(&exercise)
	if result.RowsAffected > 0 {
		log.Printf("Update Exercise History complete\nAffected row : %v", result.RowsAffected)
	} else {
		log.Printf("Update Exercise History failed\nAffected row : %v", result.RowsAffected)
	}
	return result.RowsAffected
}

func (exerciseRepo) FindExerciseByDay(keyword string) ([]model.Exercise, error) {
	exercise := []model.Exercise{}
	result := db.Where("DATE(edate) = ?", keyword).Find(&exercise)
	if result.Error != nil {
		return nil, result.Error
	}
	return exercise, nil
}

func (exerciseRepo) FindExerciseLast7Days() ([]model.Exercise, error) {
	exercise := []model.Exercise{}
	result := db.Where("edate >= NOW() - INTERVAL 7 DAY").Find(&exercise)
	if result.Error != nil {
		return nil, result.Error
	}
	return exercise, nil
}

func (exerciseRepo) FindExercisesLast12Months() ([]model.Exercise, error) {
	exercises := []model.Exercise{}
	result := db.Where("edate >= DATE_SUB(CURDATE(), INTERVAL 12 MONTH)").Find(&exercises)
	if result.Error != nil {
		return nil, result.Error
	}
	return exercises, nil
}
