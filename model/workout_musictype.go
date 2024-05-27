package model

type WorkoutMusictype struct {
	ID             int            `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	Wpid           int            `gorm:"column:wpid;NOT NULL"`
	WorkoutProfile WorkoutProfile `gorm:"foreignKey:Wpid"`
	Mtid           int            `gorm:"column:mtid;NOT NULL"`
	MusicType      MusicType      `gorm:"foreignKey:Mtid"`
}

func (m *WorkoutMusictype) TableName() string {
	return "workout_musictype"
}
