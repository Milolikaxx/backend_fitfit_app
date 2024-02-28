package service

import (
	"backend_fitfit_app/model"
	"backend_fitfit_app/repository"
)

type mtServ struct{}

func NewMtService() MtService {
	return mtServ{}
}

var musicTypeRepo = repository.NewMusicTypeRepository()

type MtService interface {
	GetAllMt() ([]model.MusicType, error)
	GetMtByID(key int) (*model.MusicType, error)
}

func (mtServ) GetAllMt() ([]model.MusicType, error) {
	music_types, err := musicTypeRepo.FindAll()
	if err != nil {
		return nil, err
	}
	return music_types, nil
}

func (mtServ) GetMtByID(id int) (*model.MusicType, error) {
	music_type, err := musicTypeRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return music_type, nil
}
