package service

import (
	"backend_fitfit_app/model"
	"backend_fitfit_app/repository"
)

type wpServ struct{}

func NewWpService() WpService {
	return wpServ{}
}

var wpRepo = repository.NewWpRepository()

type WpService interface {
	GetAllWps() ([]model.WorkoutProfile, error)
	GetWpByID(key int) (*model.WorkoutProfile, error)
	Save(model.WorkoutProfile) int64
	Update(wp model.WorkoutProfile, id int) int64
}

func (wpServ) GetAllWps() ([]model.WorkoutProfile, error) {
	wps, err := wpRepo.FindAll()
	if err != nil {
		return nil, err
	}
	return wps, nil
}

func (wpServ) GetWpByID(id int) (*model.WorkoutProfile, error) {
	wp, err := wpRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return wp, nil
}

func (wpServ) Save(wp model.WorkoutProfile) int64 {
	rowsAff := wpRepo.AddWorkProfile(wp)
	if rowsAff > 0 {
		return 1
	} else if rowsAff == 0 {
		return 0
	} else {
		return -1
	}
}

func (wpServ) Update(wp model.WorkoutProfile, id int) int64 {
	rowsAff := wpRepo.UpdateWorkProfile(wp, id)
	if rowsAff > 0 {
		return 1
	} else if rowsAff == 0 {
		return 0
	} else {
		return -1
	}
}
