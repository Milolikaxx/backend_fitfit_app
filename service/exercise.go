package service

import (
	"backend_fitfit_app/model"
	"backend_fitfit_app/repository"
	"time"
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
	Update(exercise model.Exercise, id int) ([]model.Exercise, int64)
	SearchByDay(keyword string) ([]model.Exercise, error)
	// ExerciseLast7Day() ([]model.Exercise, error)
	ExerciseLast7Day() (map[string]interface{}, error)
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

func (exerServ) Update(exercise model.Exercise, id int) ([]model.Exercise, int64) {
	rowsAff := exerciseRepo.UpdateExercise(exercise, id)
	if rowsAff > 0 {
		exercise, _ := exerciseRepo.FindByID(id)
		return exercise, 1

	} else if rowsAff == 0 {
		return nil, 0
	} else {
		return nil, -1
	}
}

func (exerServ) SearchByDay(keyword string) ([]model.Exercise, error) {
	exercise, err := exerciseRepo.FindExerciseByDay(keyword)
	if err != nil {
		return nil, err
	}
	return exercise, nil
}

// func (exerServ) ExerciseLast7Day() ([]model.Exercise, error) {
// 	exercise, err := exerciseRepo.FindExerciseLast7Days()
// 	if err != nil {
// 		return nil, err
// 	}
// 	return exercise, nil
// }

func (exerServ) ExerciseLast7Day() (map[string]interface{}, error) {
	exercises, err := exerciseRepo.FindExerciseLast7Days()
	if err != nil {
		return nil, err
	}

	// Initialize the result map with keys for the last 7 days
	result := make(map[string]interface{})
	for i := 0; i < 7; i++ {
		day := time.Now().AddDate(0, 0, -i).Format("2006-01-02")
		result[day] = map[string]interface{}{
			"exercises": []model.Exercise{},
			"count":     0,
		}
	}

	// Group exercises by date and count them
	for _, exercise := range exercises {
		day := exercise.Edate.Format("2006-01-02")
		dayData := result[day].(map[string]interface{})
		dayData["exercises"] = append(dayData["exercises"].([]model.Exercise), exercise)
		dayData["count"] = dayData["count"].(int) + 1
	}

	return result, nil
}
