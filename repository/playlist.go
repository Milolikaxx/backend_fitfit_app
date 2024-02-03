package repository

import (
	"backend_fitfit_app/model"
	"log"

	"gorm.io/gorm"
)

type playlistRepo struct {
	db *gorm.DB
}

func NewPlaylistRepository() playlistRepository {
	db, err := NewDatabaseConnection()
	if err != nil {
		return nil
	}
	return playlistRepo{db: db}
}

type playlistRepository interface {
	FindAll() ([]model.Playlist, error)
	FindByID(key int) (*model.Playlist, error)
	AddPlaylist(model.Playlist) int64
	UpdatePlaylist(model.Playlist, int) int64
}

func (p playlistRepo) FindAll() ([]model.Playlist, error) {
	playlist := []model.Playlist{}
	result := p.db.Find(&playlist)
	if result.Error != nil {
		return nil, result.Error
	}
	return playlist, nil
}

func (p playlistRepo) FindByID(id int) (*model.Playlist, error) {
	playlist := model.Playlist{}
	result := p.db.Where("pid = ?", id).Find(&playlist)
	if result.Error != nil {
		return nil, result.Error
	}
	return &playlist, nil
}

func (p playlistRepo) AddPlaylist(playlist model.Playlist) int64 {
	result := p.db.Create(&playlist)
	if result.RowsAffected > 0 {
		log.Printf("Add workoutProfile complete\nAffected row : %v", result.RowsAffected)
	} else {
		log.Printf("Add workoutProfile failed %v", result.RowsAffected)
	}
	return result.RowsAffected
}

func (u playlistRepo) UpdatePlaylist(playlist model.Playlist, id int) int64 {
	result := u.db.Model(&model.Playlist{}).Where("pid = ?", id).Updates(&playlist)
	if result.RowsAffected > 0 {
		log.Printf("Update Playlist complete\nAffected row : %v", result.RowsAffected)
	} else {
		log.Printf("Update Playlist failed\nAffected row : %v", result.RowsAffected)
	}
	return result.RowsAffected
}
