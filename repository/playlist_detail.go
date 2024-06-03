package repository

import (
	"backend_fitfit_app/model"
	"log"

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
	AddMusicToPlaylist(model.PlaylistDetail) int64
	DeleteMusicInPlaylist(key int) (int, error)
}

func (playlistDetailRepo) FindAll() ([]model.PlaylistDetail, error) {
	playlistDetail := []model.PlaylistDetail{}
	result := db.Joins("Music").Find(&playlistDetail)
	if result.Error != nil {
		return nil, result.Error
	}
	return playlistDetail, nil
}

func (playlistDetailRepo) AddMusicToPlaylist(playlistDetail model.PlaylistDetail) int64 {
	result := db.Create(&playlistDetail)
	if result.RowsAffected > 0 {
		log.Printf("Add Music complete\nAffected row : %v", result.RowsAffected)
		return int64(playlistDetail.ID)
	} else {
		log.Printf("Add Music failed %v", result.RowsAffected)
		return 0
	}
}

func (playlistDetailRepo) DeleteMusicInPlaylist(id int) (int, error) {
	result := db.Delete(&model.PlaylistDetail{}, id)
	if result.Error != nil {
		return int(result.RowsAffected), result.Error
	}
	return int(result.RowsAffected), nil
}
