package service

import (
	"backend_fitfit_app/model"
	"backend_fitfit_app/repository"
)

type MusicHistoryServ struct{}

func NewMusicHistoryService() MusicHistoryServ {
	return MusicHistoryServ{}
}

var musichistoryRepo = repository.NewMusicHistoryRepository()

type MusicHistoryService interface {
	GetAllMusicHistory() ([]model.MusicHistory, error)
	GetAllMusicHistoryByEid(id int) ([]model.MusicHistory, error)
	AddMusicHistory(musichistory model.MusicHistory) int64
}

func (MusicHistoryServ) GetAllMusicHistory() ([]model.MusicHistory, error) {
	musichis, err := musichistoryRepo.FindAllMusicHistory()
	if err != nil {
		return nil, err
	}
	return musichis, nil
}

func (MusicHistoryServ) GetAllMusicHistoryByEid(id int) ([]model.MusicHistory, error) {
	musichis, err := musichistoryRepo.FindAllMusicHistoryByEid(id)
	if err != nil {
		return nil, err
	}
	return musichis, nil
}

func (MusicHistoryServ) AddMusicHistory(musichistory model.MusicHistory) int64 {
	eid := musichistoryRepo.AddMusicHistory(musichistory)
	if eid > 0 {
		return eid
	} else if eid == 0 {
		return 0
	} else {
		return -1
	}
}
