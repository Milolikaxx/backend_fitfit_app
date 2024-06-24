// package model

// import (
// 	"time"
// )

// type Exercise struct {
// 	Eid           int          `gorm:"column:eid;type:int(11);primary_key;AUTO_INCREMENT"`
// 	Uid           int          `gorm:"column:uid;type:int(11);NOT NULL"`
// 	MHistoryid    int          `gorm:"column:m_historyid;type:int(11);NOT NULL"`
// 	Edate         time.Time    `gorm:"column:edate;type:date;NOT NULL"`
// 	Estart        time.Time    `gorm:"column:estart;type:time;NOT NULL"`
// 	Estop         time.Time    `gorm:"column:estop;type:time;NOT NULL"`
// 	PlaylistName  string       `gorm:"column:playlist_name;type:varchar(100);NOT NULL"`
// 	ExerciseType  ExerciseType `gorm:"column:exercise_type;NOT NULL"`
// 	LevelExercise int          `gorm:"column:level_exercise;type:int(11);NOT NULL"`
// 	ImagePlaylist string       `gorm:"column:image_playlist;type:varchar(1000);NOT NULL"`
// 	Duration      int          `gorm:"column:duration;type:int(11);NOT NULL"`
// }

// func (m *Exercise) TableName() string {
// 	return "exercise"
// }

package model

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type MyTime struct {
	time.Time
}

const timeLayout = "15:04:05"

// Scan implements the sql.Scanner interface.
func (t *MyTime) Scan(value interface{}) error {
	if value == nil {
		*t = MyTime{Time: time.Time{}}
		return nil
	}

	switch v := value.(type) {
	case time.Time:
		*t = MyTime{Time: v}
		return nil
	case []byte:
		parsedTime, err := time.Parse(timeLayout, string(v))
		if err != nil {
			return err
		}
		*t = MyTime{Time: parsedTime}
		return nil
	case string:
		parsedTime, err := time.Parse(timeLayout, v)
		if err != nil {
			return err
		}
		*t = MyTime{Time: parsedTime}
		return nil
	default:
		return fmt.Errorf("cannot scan type %T into MyTime", value)
	}
}

// Value implements the driver.Valuer interface.
func (t MyTime) Value() (driver.Value, error) {
	return t.Format(timeLayout), nil
}

// Your Exercise struct with MyTime fields
type Exercise struct {
	Eid           int          `gorm:"column:eid;type:int(11);primary_key;AUTO_INCREMENT"`
	Uid           int          `gorm:"column:uid;type:int(11);NOT NULL"`
	MHistoryid    int          `gorm:"column:m_historyid;type:int(11);NOT NULL"`
	Edate         time.Time    `gorm:"column:edate;type:date;NOT NULL"`
	Estart        MyTime       `gorm:"column:estart;type:time;NOT NULL"`
	Estop         MyTime       `gorm:"column:estop;type:time;NOT NULL"`
	PlaylistName  string       `gorm:"column:playlist_name;type:varchar(100);NOT NULL"`
	ExerciseType  ExerciseType `gorm:"column:exercise_type;NOT NULL"`
	LevelExercise int          `gorm:"column:level_exercise;type:int(11);NOT NULL"`
	ImagePlaylist string       `gorm:"column:image_playlist;type:varchar(1000);NOT NULL"`
	Duration      int          `gorm:"column:duration;type:int(11);NOT NULL"`
}

func (m *Exercise) TableName() string {
	return "exercise"
}
