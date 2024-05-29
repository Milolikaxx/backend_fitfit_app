package model

type PlaylistDetail struct {
	ID  int `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	Pid int `gorm:"column:pid;NOT NULL"`
	Mid int `gorm:"column:mid;NOT NULL"`
}

func (m *PlaylistDetail) TableName() string {
	return "playlist_detail"
}
