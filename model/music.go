package model

type Music struct {
	Mid        int       `gorm:"column:mid;primaryKey;autoIncrement"`
	Mtid       int       `gorm:"column:mtid;not null"`
	MusicType  MusicType `gorm:"foreignKey:Mtid;references:Mtid"`
	MLink      string    `gorm:"column:m_link;not null"`
	Name       string    `gorm:"column:name;not null"`
	MusicImage string    `gorm:"column:music_image"`
	Artist     string    `gorm:"column:artist;not null"`
	Duration   float64   `gorm:"column:duration;not null"`
	Bpm        int       `gorm:"column:bpm;not null"`
}

// JOin
// type Music struct {
// 	Mid            int            `gorm:"column:mid;primary_key;AUTO_INCREMENT"`
// 	Mtid           int            `gorm:"column:mtid;NOT NULL"`
// 	MusicType      MusicType      `gorm:"foreignKey:Mtid"`
// 	MLink          string         `gorm:"column:m_link;NOT NULL"`
// 	Name           string         `gorm:"column:name;NOT NULL"`
// 	MusicImage     string         `gorm:"column:music_image"`
// 	Artist         string         `gorm:"column:artist;NOT NULL"`
// 	Duration       float64        `gorm:"column:duration;NOT NULL"`
// 	Bpm            int            `gorm:"column:bpm;NOT NULL"`
// }

// type Music struct {
// 	Mid        int        `gorm:"column:mid;primary_key;AUTO_INCREMENT"`
// 	Mtid       int        `gorm:"column:mtid;NOT NULL"`
// 	MusicType  MusicType  `gorm:"foreignKey:Mtid"`
// 	MLink      string     `gorm:"column:m_link;NOT NULL"`
// 	Name       string     `gorm:"column:name;NOT NULL"`
// 	MusicImage string     `gorm:"column:music_image"`
// 	Artist     string     `gorm:"column:artist;NOT NULL"`
// 	Duration   float64    `gorm:"column:duration;NOT NULL"`
// 	Bpm        int        `gorm:"column:bpm;NOT NULL"`
// 	Playlists  []Playlist `gorm:"many2many:playlist_detail;foreignKey:Mid;joinForeignKey:Mid;References:Pid;joinReferences:Pid"`
// }

func (m *Music) TableName() string {
	return "music"
}
