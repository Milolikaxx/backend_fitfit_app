package model

type MusicType struct {
	Mtid int    `gorm:"column:mtid;primary_key;AUTO_INCREMENT"`
	Name string `gorm:"column:name;NOT NULL"`
}

func (m *MusicType) TableName() string {
	return "music_type"
}
