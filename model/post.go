package model

import (
	"time"
)

type Post struct {
	Postid int  `gorm:"column:postid;primary_key;AUTO_INCREMENT"`
	Uid    int  `gorm:"column:uid;NOT NULL"`
	User   User `gorm:"foreignKey:Uid;references:Uid"`
	Pid    int  `gorm:"column:pid;NOT NULL"`
	// Playlist     Playlist  `gorm:"foreignKey:Pid;references:Pid"`
	PlaylistName string    `gorm:"column:playlist_name;NOT NULL"`
	Description  string    `gorm:"column:description;NOT NULL"`
	PDatetime    time.Time `gorm:"column:p_datetime;default:CURRENT_TIMESTAMP;NOT NULL"`
	CreatedAt    time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP;NOT NULL"`
	UpdatedAt    time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP;NOT NULL"`
}

func (m *Post) TableName() string {
	return "post"
}
