package model

import (
	"time"
)

type Playlist struct {
	Pid              int              `gorm:"column:pid;primaryKey;autoIncrement"`
	Wpid             int              `gorm:"column:wpid;not null"`
	WorkoutProfile   WorkoutProfile   `gorm:"foreignKey:Wpid;references:Wpid"`
	PlaylistName     string           `gorm:"column:playlist_name;not null"`
	DurationPlaylist float64          `gorm:"column:duration_playlist;not null"`
	ImagePlaylist    string           `gorm:"column:image_playlist;not null"`
	CreatedAt        time.Time        `gorm:"column:created_at;default:CURRENT_TIMESTAMP;not null"`
	UpdatedAt        time.Time        `gorm:"column:updated_at;default:CURRENT_TIMESTAMP;not null"`
	PlaylistDetail   []PlaylistDetail `gorm:"foreignKey:Pid;references:Pid"`
	// TotalTime        float64          `json:"TotalTime"`
}

// JOin
// type Playlist struct {
// 	Pid              int            `gorm:"column:pid;primary_key;AUTO_INCREMENT"`
// 	Wpid             int            `gorm:"column:wpid;NOT NULL"`
// 	PlaylistName     string         `gorm:"column:playlist_name;NOT NULL"`
// 	DurationPlaylist float64        `gorm:"column:duration_playlist;NOT NULL"`
// 	ImagePlaylist    string         `gorm:"column:image_playlist;NOT NULL"`
// 	CreatedAt        time.Time      `gorm:"column:created_at;default:CURRENT_TIMESTAMP;NOT NULL"`
// 	UpdatedAt        time.Time      `gorm:"column:updated_at;default:CURRENT_TIMESTAMP;NOT NULL"`
// 	PlaylistDetail   PlaylistDetail `gorm:"foreignKey:Pid"`
// }

//Preload
// type Playlist struct {
// 	Pid              int       `gorm:"column:pid;primary_key;AUTO_INCREMENT"`
// 	Wpid             int       `gorm:"column:wpid;NOT NULL"`
// 	PlaylistName     string    `gorm:"column:playlist_name;NOT NULL"`
// 	DurationPlaylist float64   `gorm:"column:duration_playlist;NOT NULL"`
// 	ImagePlaylist    string    `gorm:"column:image_playlist;NOT NULL"`
// 	CreatedAt        time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP;NOT NULL"`
// 	UpdatedAt        time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP;NOT NULL"`
// 	Musics           []Music   `gorm:"many2many:playlist_detail;foreignKey:Pid;joinForeignKey:Pid;References:Mid;joinReferences:Mid"`
// }

func (m *Playlist) TableName() string {
	return "playlist"
}
