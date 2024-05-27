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
	GetWpByWPID(key int) (*model.WorkoutProfile, error)
	GetListWpByWPID(key int) ([]model.WorkoutProfile, error)
	GetWpByUID(key int) ([]model.WorkoutProfile, error)
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

func (wpServ) GetWpByWPID(id int) (*model.WorkoutProfile, error) {
	wp, err := wpRepo.FindByWPID(id)
	if err != nil {
		return nil, err
	}
	return wp, nil
}
func (wpServ) GetListWpByWPID(id int) ([]model.WorkoutProfile, error) {
	wps, err := wpRepo.FindListByWPID(id)
	if err != nil {
		return nil, err
	}
	return wps, nil
}
func (wpServ) GetWpByUID(id int) ([]model.WorkoutProfile, error) {
	wps, err := wpRepo.FindByUID(id)
	if err != nil {
		return nil, err
	}
	return wps, nil
}

func (wpServ) Save(wp model.WorkoutProfile) int64 {
	wpid := wpRepo.AddWorkProfile(wp)
	if wpid > 0 {
		return wpid
	} else if wpid == 0 {
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
