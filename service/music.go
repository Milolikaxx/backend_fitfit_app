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
	GetMusicByLevel(bpm int, musicType []int) ([]model.Music, error)
	SearchMusic(searchMusic model.SearchMusic) ([]model.Music, error)
}

func (musicServ) GetMusicByMtid(musicType int) ([]model.Music, error) {
	music, err := musicRepo.FindAllMusicByMusictype(musicType)
	if err != nil {
		return nil, err
	}
	return music, nil
}

func (musicServ) GetMusicByLevel(bpm int, musicType []int) ([]model.Music, error) {
	music, err := musicRepo.FindAllMusicByLevel(bpm, musicType)
	if err != nil {
		return nil, err
	}
	return music, nil
}

func (musicServ) SearchMusic(searchMusic model.SearchMusic) ([]model.Music, error) {
	music, err := musicRepo.SearchMusic(searchMusic)
	if err != nil {
		return nil, err
	}
	return music, nil
}
