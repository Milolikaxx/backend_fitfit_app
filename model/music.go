package model

type Music struct {
	Mid        int        `gorm:"column:mid;primary_key;AUTO_INCREMENT"`
	Mtid       int        `gorm:"column:mtid;NOT NULL"`
	MLink      string     `gorm:"column:m_link;NOT NULL"`
	Name       string     `gorm:"column:name;NOT NULL"`
	MusicImage string     `gorm:"column:music_image"`
	Artist     string     `gorm:"column:artist;NOT NULL"`
	Duration   float64    `gorm:"column:duration;NOT NULL"`
	Bpm        int        `gorm:"column:bpm;NOT NULL"`
}

func (m *Music) TableName() string {
	return "music"
}
