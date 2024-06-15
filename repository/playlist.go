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
	FindAllByWpid(id int) ([]model.Playlist, error)
	FindByID(key int) (*model.Playlist, error)
	FindWithoutMusicByID(id int) (*model.Playlist, error)
	AddPlaylist(model.Playlist) int64
	UpdatePlaylist(model.Playlist, int) int64
	DeletePlaylist(key int) (int, error)
}

func (playlistRepo) FindAll() ([]model.Playlist, error) {
	playlist := []model.Playlist{}
	result := db.Find(&playlist)
	if result.Error != nil {
		return nil, result.Error
	}
	return playlist, nil
}
func (playlistRepo) FindAllByWpid(id int) ([]model.Playlist, error) {
	playlist := []model.Playlist{}
	result := db.Select("pid", "wpid", "playlist_name", "duration_playlist", "image_playlist").Where("wpid = ?", id).Find(&playlist)
	if result.Error != nil {
		return nil, result.Error
	}
	return playlist, nil
}
func (playlistRepo) FindByID(id int) (*model.Playlist, error) {
	playlist := model.Playlist{}
	result := db.Preload("PlaylistDetail.Music.MusicType").Where("pid = ?", id).Find(&playlist)
	if result.Error != nil {
		return nil, result.Error
	}
	return &playlist, nil
}

func (playlistRepo) FindWithoutMusicByID(id int) (*model.Playlist, error) {
	playlist := model.Playlist{}
	result := db.Where("pid = ?", id).Find(&playlist)
	if result.Error != nil {
		return nil, result.Error
	}
	return &playlist, nil
}

// Join
// func (playlistRepo) FindByID(id int) ([]model.Playlist, error) {
// 	playlist := []model.Playlist{}
// 	result := db.Joins("PlaylistDetail").Joins("PlaylistDetail.Music").Where("playlist.pid = ?", id).Find(&playlist)
// 	if result.Error != nil {
// 		return nil, result.Error
// 	}
// 	return playlist, nil
// }

func (playlistRepo) AddPlaylist(playlist model.Playlist) int64 {
	result := db.Create(&playlist)
	if result.RowsAffected > 0 {
		log.Printf("Add Playlist complete\nAffected row : %v", result.RowsAffected)
		return int64(playlist.Pid)
	} else {
		log.Printf("Add Playlist failed %v", result.RowsAffected)
		return 0
	}

}

func (playlistRepo) UpdatePlaylist(playlist model.Playlist, id int) int64 {
	result := db.Model(&model.Playlist{}).Where("pid = ?", id).Updates(&playlist)
	if result.RowsAffected > 0 {
		log.Printf("Update Playlist complete\nAffected row : %v", result.RowsAffected)
	} else {
		log.Printf("Update Playlist failed\nAffected row : %v", result.RowsAffected)
	}
	return result.RowsAffected
}

func (playlistRepo) DeletePlaylist(id int) (int, error) {
	result := db.Delete(&model.Playlist{}, id)
	if result.Error != nil {
		return int(result.RowsAffected), result.Error
	}
	return int(result.RowsAffected), nil
}
