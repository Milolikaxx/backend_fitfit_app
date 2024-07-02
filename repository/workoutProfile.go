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
	FindByWpid(key int) (*model.WorkoutProfile, error)
	FindListByUid(key int) ([]model.WorkoutProfile, error)
	// FindByUID(key int) ([]model.WorkoutProfile, error)
	AddWorkProfile(model.WorkoutProfile) int64
	DeleteProfile(id int) (int, error)
	UpdateWorkProfile(wp model.WorkoutProfile, id int) int64
	FindByWord(key string) ([]model.WorkoutProfile, error)
}

func (u wpRepo) FindAll() ([]model.WorkoutProfile, error) {
	wps := []model.WorkoutProfile{}
	result := u.db.Find(&wps)
	if result.Error != nil {
		return nil, result.Error
	}
	return wps, nil
}

func (wpRepo) FindByWpid(id int) (*model.WorkoutProfile, error) {
	wp := model.WorkoutProfile{}
	result := db.Preload("WorkoutMusictype.MusicType").Where("workout_profile.wpid = ?", id).Find(&wp)
	if result.Error != nil {
		return nil, result.Error
	}
	return &wp, nil
}

func (wpRepo) FindListByUid(id int) ([]model.WorkoutProfile, error) {
	wps := []model.WorkoutProfile{}
	// result := db.Joins("WorkoutMusictype").Joins("WorkoutMusictype.MusicType").Where("workout_profile.uid = ?", id).Find(&wps)
	result := db.Preload("WorkoutMusictype.MusicType").Where("workout_profile.uid = ?", id).Find(&wps)
	if result.Error != nil {
		return nil, result.Error
	}
	return wps, nil
}

// func (wpRepo) FindByUID(uid int) ([]model.WorkoutProfile, error) {
// 	wps := []model.WorkoutProfile{}
// 	result := db.Where("uid = ?", uid).Find(&wps)
// 	if result.Error != nil {
// 		return nil, result.Error
// 	}
// 	return wps, nil
// }

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

func (wpRepo) DeleteProfile(id int) (int, error) {
	result := db.Delete(&model.WorkoutProfile{}, id)
	if result.Error != nil {
		return int(result.RowsAffected), result.Error
	}
	return int(result.RowsAffected), nil
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

func (wpRepo) FindByWord(key string) ([]model.WorkoutProfile, error) {
	wps := []model.WorkoutProfile{}
	result := db.Preload("WorkoutMusictype.MusicType").Where("exercise_type like ? ", "%"+key+"%").Find(&wps)
	if result.Error != nil {
		return nil, result.Error
	}
	return wps, nil
}
