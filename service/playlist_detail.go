package service

import (
	"backend_fitfit_app/model"
	"backend_fitfit_app/repository"
)

type playlistDetailServ struct{}

func NewPlaylistDetailService() PlaylistDetailService {
	return playlistDetailServ{}
}

var playlistDetailRepo = repository.NewPlaylistDetailRepository()

type PlaylistDetailService interface {
	GetAllPlaylistDetail() ([]model.PlaylistDetail, error)
	GetListWpByPID(key int) ([]model.PlaylistDetail, error)
	// GetByID(key int) (*model.Playlist, error)
	Save(model.PlaylistDetail) int64
	// Update(playlist model.Playlist, id int) int64
}

func (playlistDetailServ) GetAllPlaylistDetail() ([]model.PlaylistDetail, error) {
	playlistDetail, err := playlistDetailRepo.FindAll()
	if err != nil {
		return nil, err
	}
	return playlistDetail, nil
}
func (playlistDetailServ) GetListWpByPID(id int) ([]model.PlaylistDetail, error) {
	playlistDetail, err := playlistDetailRepo.FindListByPID(id)
	if err != nil {
		return nil, err
	}
	return playlistDetail, nil
}
func (playlistDetailServ) Save(PlaylistDetail model.PlaylistDetail) int64 {
	pldid := playlistDetailRepo.AddMusicToPlaylist(PlaylistDetail)
	if pldid > 0 {
		return pldid
	} else if pldid == 0 {
		return 0
	} else {
		return -1
	}
}
