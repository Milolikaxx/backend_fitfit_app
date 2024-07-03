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
	GetWpByWpid(key int) (*model.WorkoutProfile, error)
	GetListWpByUid(key int) ([]model.WorkoutProfile, error)
	// GetWpByUID(key int) ([]model.WorkoutProfile, error)
	Save(model.WorkoutProfile) int64
	Delete(id int) (int64, error)
	Update(wp model.WorkoutProfile, id int) int64
	GetListWpByKey(key string) ([]model.WorkoutProfile, error)
}

func (wpServ) GetAllWps() ([]model.WorkoutProfile, error) {
	wps, err := wpRepo.FindAll()
	if err != nil {
		return nil, err
	}
	return wps, nil
}

func (wpServ) GetWpByWpid(id int) (*model.WorkoutProfile, error) {
	wp, err := wpRepo.FindByWpid(id)
	if err != nil {
		return nil, err
	}
	return wp, nil
}

func (wpServ) GetListWpByUid(id int) ([]model.WorkoutProfile, error) {
	wps, err := wpRepo.FindListByUid(id)
	if err != nil {
		return nil, err
	}
	return wps, nil
}

// func (wpServ) GetWpByUID(id int) ([]model.WorkoutProfile, error) {
// 	wps, err := wpRepo.FindByUID(id)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return wps, nil
// }

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

func (wpServ) Delete(id int) (int64, error) {
	row, error := wpRepo.DeleteProfile(id)
	if row > 0 {
		return 1, error
	} else if row == 0 {
		return 0, error
	} else {
		return -1, error
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

func (wpServ) GetListWpByKey(key string) ([]model.WorkoutProfile, error) {
	wps, err := wpRepo.FindByWord(key)
	if err != nil {
		return nil, err
	}
	return wps, nil
}
