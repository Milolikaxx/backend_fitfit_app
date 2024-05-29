package service

import (
	"backend_fitfit_app/model"
	"backend_fitfit_app/repository"
)

type musicServ struct{}

func NewMusicService() MusicService {
	return musicServ{}
}

var musicRepo = repository.NewMusicRepository()

type MusicService interface {
	GetMusicByMtid(key int) ([]model.Music, error)
}

func (musicServ) GetMusicByMtid(id int) ([]model.Music, error) {
	music, err := musicRepo.FindAllMusicByMusictype(id)
	if err != nil {
		return nil, err
	}
	return music, nil
}
