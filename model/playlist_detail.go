package model

type PlaylistDetail struct {
	ID    int   `gorm:"column:id;primaryKey;autoIncrement"`
	Pid   int   `gorm:"column:pid;not null"`
	Mid   int   `gorm:"column:mid;not null"`
	Music Music `gorm:"foreignKey:Mid;references:Mid"`
}

// Join and Preload
// type PlaylistDetail struct {
// 	ID    int   `gorm:"column:id;primary_key;AUTO_INCREMENT"`
// 	Pid   int   `gorm:"column:pid;NOT NULL"`
// 	Mid   int   `gorm:"column:mid;NOT NULL"`
// 	Music Music `gorm:"foreignKey:Mid"`
// }

func (m *PlaylistDetail) TableName() string {
	return "playlist_detail"
}
