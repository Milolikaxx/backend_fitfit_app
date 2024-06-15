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
	GetAllPlaylistByWpid(id int) ([]model.Playlist, error)
	GetPlaylistMusicByID(key int) (*model.Playlist, error)
	GetPlaylistWithOutMusicByID(id int) (*model.Playlist, error)
	Save(model.Playlist) int64
	Update(playlist model.Playlist, id int) int64
	Delete(id int) (int64, error)
}

func (playlistServ) GetAllPlaylist() ([]model.Playlist, error) {
	playlist, err := playlistRepo.FindAll()
	if err != nil {
		return nil, err
	}
	return playlist, nil
}
func (playlistServ) GetAllPlaylistByWpid(id int) ([]model.Playlist, error) {
	playlist, err := playlistRepo.FindAllByWpid(id)
	if err != nil {
		return nil, err
	}
	return playlist, nil
}

func (playlistServ) GetPlaylistMusicByID(id int) (*model.Playlist, error) {
	playlist, err := playlistRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return playlist, nil
}

func (playlistServ) GetPlaylistWithOutMusicByID(id int) (*model.Playlist, error) {
	playlist, err := playlistRepo.FindWithoutMusicByID(id)
	if err != nil {
		return nil, err
	}
	return playlist, nil
}

func (playlistServ) Save(playlist model.Playlist) int64 {
	pid := playlistRepo.AddPlaylist(playlist)
	if pid > 0 {
		return pid
	} else if pid == 0 {
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

func (playlistServ) Delete(id int) (int64, error) {
	row, error := playlistRepo.DeletePlaylist(id)
	if row > 0 {
		return 1, error
	} else if row == 0 {
		return 0, error
	} else {
		return -1, error
	}
}
