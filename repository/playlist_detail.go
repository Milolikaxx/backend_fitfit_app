package repository

import (
	"backend_fitfit_app/model"
	// "log"

	"gorm.io/gorm"
)

type playlistDetailRepo struct {
	db *gorm.DB
}

func NewPlaylistDetailRepository() playlistDetailRepository {
	db, err := NewDatabaseConnection()
	if err != nil {
		return nil
	}
	return playlistDetailRepo{db: db}
}

type playlistDetailRepository interface {
	FindAll() ([]model.PlaylistDetail, error)
	FindListByPID(key int) ([]model.PlaylistDetail, error)
	// FindByID(key int) (*model.PlaylistDetail, error)
	// AddPlaylist(model.PlaylistDetail) int64
	// UpdatePlaylist(model.PlaylistDetail, int) int64
}

func (playlistDetailRepo) FindAll() ([]model.PlaylistDetail, error) {
	playlistDetail := []model.PlaylistDetail{}
	result := db.Find(&playlistDetail)
	if result.Error != nil {
		return nil, result.Error
	}
	return playlistDetail, nil
}

func (playlistDetailRepo) FindListByPID(id int) ([]model.PlaylistDetail, error) {
	playlistDetail := []model.PlaylistDetail{}
	result := db.Joins("playlist").Where("playlist_detail.pid = ?", id).Find(&playlistDetail)
	if result.Error != nil {
		return nil, result.Error
	}
	return playlistDetail, nil
}
