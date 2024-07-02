package service

import (
	"backend_fitfit_app/model"
	"backend_fitfit_app/repository"
)

type exerServ struct{}

func NewExerService() exerService {
	return exerServ{}
}

var exerciseRepo = repository.NewExerciseRepository()

type exerService interface {
	GetAllExer() ([]model.Exercise, error)
	GetExerByID(key int) ([]model.Exercise, error)
	Save(exercise model.Exercise) int64
}

func (exerServ) GetAllExer() ([]model.Exercise, error) {
	exercise, err := exerciseRepo.FindAll()
	if err != nil {
		return nil, err
	}
	return exercise, nil
}

func (exerServ) GetExerByID(id int) ([]model.Exercise, error) {
	exercise, err := exerciseRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return exercise, nil
}

func (exerServ) Save(exercise model.Exercise) int64 {
	eid := exerciseRepo.AddExercise(exercise)
	if eid > 0 {
		return eid
	} else if eid == 0 {
		return 0
	} else {
		return -1
	}
}
