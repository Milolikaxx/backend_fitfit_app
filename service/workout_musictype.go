package service

import (
	"backend_fitfit_app/model"
	"backend_fitfit_app/repository"
)

type wpmtServ struct{}

func NewWpMusicTypeService() WpMusicTypeService {
	return wpmtServ{}
}

var wpmtRepo = repository.NewWpMusicTypeRepository()

type WpMusicTypeService interface {
	GetByWPID(key int) ([]model.WorkoutMusictype, error)
	Save(model.WorkoutMusictype) int64
	Update(wp model.WorkoutMusictype, id int) int64
}

func (wpmtServ) GetByWPID(id int) ([]model.WorkoutMusictype, error) {
	wpmt, err := wpmtRepo.FindByWPID(id)
	if err != nil {
		return nil, err
	}
	return wpmt, nil
}
func (wpmtServ) Save(wpmt model.WorkoutMusictype) int64 {
	rowsAff := wpmtRepo.AddWpMusicType(wpmt)
	if rowsAff > 0 {
		return 1
	} else if rowsAff == 0 {
		return 0
	} else {
		return -1
	}
}

func (wpmtServ) Update(wpmt model.WorkoutMusictype, id int) int64 {
	rowsAff := wpmtRepo.UpdateWpMusicType(wpmt, id)
	if rowsAff > 0 {
		return 1
	} else if rowsAff == 0 {
		return 0
	} else {
		return -1
	}
}
