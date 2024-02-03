package model

import (
	"time"
)

type User struct {
	Uid          int       `gorm:"column:uid;primary_key;AUTO_INCREMENT"`
	Name         string    `gorm:"column:name;NOT NULL"`
	Birthday     time.Time `gorm:"column:birthday"`
	Email        string    `gorm:"column:email;NOT NULL"`
	Password     string    `gorm:"column:password;NOT NULL"`
	ImageProfile string    `gorm:"column:image_profile"`
	GoogleID     string    `gorm:"column:google_id"`
	CreatedAt    time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP;NOT NULL"`
	UpdatedAt    time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP;NOT NULL"`
}

func (m *User) TableName() string {
	return "user"
}
