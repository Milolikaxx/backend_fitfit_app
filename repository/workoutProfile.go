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
	// UpdateWorkProfile(id int, wp model.WorkoutProfile) int64
}

func (u wpRepo) FindAll() ([]model.WorkoutProfile, error) {
	wps := []model.WorkoutProfile{}
	result := u.db.Find(&wps)
	if result.Error != nil {
		return nil, result.Error
	}
	return wps, nil
}
func (w wpRepo) FindByID(id int) (*model.WorkoutProfile, error) {
	wp := model.WorkoutProfile{}
	result := w.db.Where("wpid = ?", id).Find(&wp)
	if result.Error != nil {
		return nil, result.Error
	}
	return &wp, nil
}

func (w wpRepo) AddWorkProfile(wp model.WorkoutProfile) int64 {
	result := w.db.Create(&wp)
	if result.RowsAffected > 0 {
		log.Printf("Add workoutProfile complete\nAffected row : %v", result.RowsAffected)
	} else {
		log.Printf("Add workoutProfile failed %v", result.RowsAffected)
	}
	return result.RowsAffected
}

// func (w wpRepo) UpdateWorkProfile(id int, wp model.WorkoutProfile) int64 {
// 	wp = model.WorkoutProfile{}
// 	result := db.Find(&wp, id)
// 	if result.Error != nil {
// 		log.Println(result.Error)
// 	}
// 	result = db.Save(&wp)
// 	if result.RowsAffected > 0 {
// 		log.Println("Update completed")
// 	}
// 	return result.RowsAffected
// }
