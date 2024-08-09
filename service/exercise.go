package service

import (
	"backend_fitfit_app/model"
	"backend_fitfit_app/repository"
	"sort"
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
	ExerciseLast7Day() ([]Day, error)
	GetLast12Months() ([]MonthlyExerciseData, error)
	SearchByMonth(numMonth string) ([]model.Exercise, error)
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

type Day struct {
	Date      string           `json:"date"`
	Exercises []model.Exercise `json:"exercises"`
	Count     int              `json:"count"`
}

func (exerServ) ExerciseLast7Day() ([]Day, error) {
	exercises, err := exerciseRepo.FindExerciseLast7Days()
	if err != nil {
		return nil, err
	}

	// Initialize the result slice with 7 Day structs for the last 7 days
	var result []Day
	for i := 0; i < 7; i++ {
		day := time.Now().AddDate(0, 0, -i).Format("2006-01-02")
		result = append(result, Day{
			Date:      day,
			Exercises: []model.Exercise{},
			Count:     0,
		})
	}

	// Group exercises by date and count them
	for _, exercise := range exercises {
		day := exercise.Edate.Format("2006-01-02")
		for i := range result {
			if result[i].Date == day {
				result[i].Exercises = append(result[i].Exercises, exercise)
				result[i].Count++
			}
		}
	}

	return result, nil
}

type MonthlyExerciseData struct {
	MonthNumber   int              `json:"month_number"`
	MonthName     string           `json:"month_name"`
	ExerciseCount int              `json:"exercise_count"`
	Exercises     []model.Exercise `json:"exercises"`
}

func (exerServ) GetLast12Months() ([]MonthlyExerciseData, error) {
	exercises, err := exerciseRepo.FindExercisesLast12Months()
	if err != nil {
		return nil, err
	}

	monthlyData := make(map[string]*MonthlyExerciseData)
	monthNames := []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}

	// Initialize data for each of the last 12 months
	for i := 0; i < 12; i++ {
		date := time.Now().AddDate(0, -i, 0)
		yearMonth := date.Format("2006-01")
		monthNumber := date.Month()
		monthName := monthNames[monthNumber-1]

		monthlyData[yearMonth] = &MonthlyExerciseData{
			MonthNumber:   int(monthNumber),
			MonthName:     monthName,
			ExerciseCount: 0,
			Exercises:     []model.Exercise{},
		}
	}

	// Populate the data with exercises
	for _, exercise := range exercises {
		yearMonth := exercise.Edate.Format("2006-01")
		if data, exists := monthlyData[yearMonth]; exists {
			data.ExerciseCount++
			data.Exercises = append(data.Exercises, exercise)
		}
	}

	var sortedMonthlyData []MonthlyExerciseData
	for _, data := range monthlyData {
		sortedMonthlyData = append(sortedMonthlyData, *data)
	}

	// Sort by year and month
	sort.Slice(sortedMonthlyData, func(i, j int) bool {
		yearMonthI, _ := time.Parse("January", sortedMonthlyData[i].MonthName)
		yearMonthJ, _ := time.Parse("January", sortedMonthlyData[j].MonthName)
		return yearMonthI.Before(yearMonthJ)
	})

	return sortedMonthlyData, nil
}

func (exerServ) SearchByMonth(numMonth string) ([]model.Exercise, error) {
	exercise, err := exerciseRepo.FindExerciseByMonth(numMonth)
	if err != nil {
		return nil, err
	}
	return exercise, nil
}
