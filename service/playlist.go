package service

import (
	"backend_fitfit_app/model"
	"backend_fitfit_app/repository"
)

type playlistServ struct{}

func NewPlaylistService() PlaylistService {
	return playlistServ{}
}

var playlistRepo = repository.NewPlaylistRepository()

type PlaylistService interface {
	GetAllPlaylist() ([]model.Playlist, error)
	GetByID(key int) (*model.Playlist, error)
	Save(model.Playlist) int64
	Update(playlist model.Playlist, id int) int64
}

func (playlistServ) GetAllPlaylist() ([]model.Playlist, error) {
	playlist, err := playlistRepo.FindAll()
	if err != nil {
		return nil, err
	}
	return playlist, nil
}

func (playlistServ) GetByID(id int) (*model.Playlist, error) {
	playlist, err := playlistRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return playlist, nil
}

func (playlistServ) Save(playlist model.Playlist) int64 {
	rowsAff := playlistRepo.AddPlaylist(playlist)
	if rowsAff > 0 {
		return 1
	} else if rowsAff == 0 {
		return 0
	} else {
		return -1
	}
}

func (playlistServ) Update(playlist model.Playlist, id int) int64 {
	rowsAff := playlistRepo.UpdatePlaylist(playlist, id)
	if rowsAff > 0 {
		return 1
	} else if rowsAff == 0 {
		return 0
	} else {
		return -1
	}
}
