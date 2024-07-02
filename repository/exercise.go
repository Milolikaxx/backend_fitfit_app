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
