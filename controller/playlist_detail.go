package controller

import (
	"backend_fitfit_app/model"
	"backend_fitfit_app/repository"
	"backend_fitfit_app/service"
	"math"
	"math/rand"

	"fmt"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
)

var playlistDetailServ = service.NewPlaylistDetailService()

func NewPlaylistDetailController(router *gin.Engine) {
	ping := router.Group("/playlist_detail")
	{
		ping.GET("", getAllPlaylistDetail)
		ping.POST("/addmusic", AddMusic)
		ping.DELETE("/delete/:id", DeleteMusic)
		ping.GET("/musiclist/:id", getMusicList)
	}
}

func getAllPlaylistDetail(ctx *gin.Context) {
	pl_detail, err := playlistDetailServ.GetAllPlaylistDetail()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error": err,
		})
	}
	ctx.JSON(http.StatusOK, pl_detail)
}

func AddMusic(ctx *gin.Context) {
	pld := model.PlaylistDetail{}
	ctx.ShouldBindJSON(&pld)
	err := playlistDetailServ.Save(pld)
	ctx.JSON(http.StatusCreated, err)
}

func DeleteMusic(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	pl_detail, err := playlistDetailServ.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error": err,
		})
	}
	ctx.JSON(http.StatusOK, pl_detail)
}

func getMusicList(ctx *gin.Context) {

	id, _ := strconv.Atoi(ctx.Param("id"))
	music, err := CreatePlaylistWP(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, music)

}

var bpm = []int{100, 114, 133, 152, 171, 190}

var levelProceed = [][]int{
	{100},          //Lv.0
	{10, 35},       //Lv.1
	{10, 12, 22},   //Lv.2
	{5, 5, 10, 25}, //Lv.3
	// {10, 20, 30, 40},        //Lv.4
	// {5, 10, 15, 10, 15, 30}, //Lv.5
}

type MusicGroup struct {
	level int
	music model.Music
}

func groupBy(music []MusicGroup, level int) []model.Music {
	groups := make([]model.Music, 0)
	for _, m := range music {
		if m.level == level {
			groups = append(groups, m.music)
		}
	}
	return groups
}

func groupMusic(music []model.Music, lvl int) []MusicGroup {
	var groupMusic []MusicGroup
	for i := 0; i <= lvl; i++ {
		for _, m := range music {
			if i == 0 {
				if m.Bpm <= bpm[i] {
					row := MusicGroup{}
					row.level = i
					row.music = m
					groupMusic = append(groupMusic, row)
				}
			} else {
				if m.Bpm >= bpm[i-1] && m.Bpm <= bpm[i] {
					row := MusicGroup{}
					row.level = i
					row.music = m
					groupMusic = append(groupMusic, row)
				}
			}
		}
	}
	return groupMusic
}
func CreatePlaylistWP(wpid int) ([]model.Music, error) {
	//workout_profile
	var wpRepo = repository.NewWpRepository()
	wp, _ := wpRepo.FindByWpid(wpid)
	fmt.Printf("wp:%v  \n\n", wp)
	//infomation
	lvl := wp.LevelExercise
	duration := int(wp.Duration * 60)
	// durationEx := wp.Duration
	exeType := wp.ExerciseType
	var musicType []int
	for _, t := range wp.WorkoutMusictype {
		musicType = append(musicType, t.Mtid)
	}
	fmt.Printf("infomation  lvl:%d  duration:%d  type:%s  musicType:%v  bpm:%d  \n\n", lvl, duration, exeType, musicType, bpm[lvl])

	//music - level
	musicRepo := repository.NewMusicRepository()
	music, _ := musicRepo.FindAllMusicByLevel(bpm[lvl], musicType)
	fmt.Printf("result music:%d\n\n", len(music))

	//grouping
	fmt.Printf("music type %s\n\n", reflect.TypeOf(music))
	groupMusic := groupMusic(music, lvl)
	//groupBy LV
	fmt.Printf("level0 = %d เพลง\n ", len(groupBy(groupMusic, 0)))
	fmt.Printf("level1 = %d เพลง\n", len(groupBy(groupMusic, 1)))
	fmt.Printf("level2 = %d เพลง\n", len(groupBy(groupMusic, 2)))
	fmt.Printf("level3 = %d เพลง\n", len(groupBy(groupMusic, 3)))
	fmt.Printf("level4 = %d เพลง\n", len(groupBy(groupMusic, 4)))
	fmt.Printf("level5 = %d เพลง\n\n", len(groupBy(groupMusic, 5)))

	//Process
	// var musicList []model.Music
	// if lvl == 2 {
	// 	if durationEx == 10 {
	// 		//ช่วง 1
	// 		groupedMusic1 := groupBy(groupMusic, 1)
	// 		length1 := len(groupedMusic1)
	// 		if length1 == 0 {
	// 			return musicList, fmt.Errorf("ไม่มีเพลงที่สามารถใช้ได้ในระดับปัจจุบัน")
	// 		}

	// 		idx1 := rand.Int() % length1
	// 		m1 := groupBy(groupMusic, 1)[idx1]

	// 		musicList = append(musicList, m1)

	// 		//ช่วง 2
	// 		groupedMusic2 := groupBy(groupMusic, 2)
	// 		length2 := len(groupedMusic2)
	// 		if length2 == 0 {
	// 			return musicList, fmt.Errorf("ไม่มีเพลงที่สามารถใช้ได้ในระดับปัจจุบัน")
	// 		}

	// 		idx2 := rand.Int() % length2
	// 		m2 := groupBy(groupMusic, 2)[idx2]
	// 		musicList = append(musicList, m2)
	// 		// ช่วง 3
	// 		groupedMusic3 := groupBy(groupMusic, 1)
	// 		length3 := len(groupedMusic3)
	// 		if length3 == 0 {
	// 			return musicList, fmt.Errorf("ไม่มีเพลงที่สามารถใช้ได้ในระดับปัจจุบัน")
	// 		}

	// 		idx3 := rand.Int() % length3
	// 		m3 := groupBy(groupMusic, 1)[idx3]
	// 		musicList = append(musicList, m3)
	// 	}
	// }
	//Process 1
	fmt.Println(levelProceed[lvl])
	timeRemain := duration                 //เวลาที่ยังเหลือ
	timeMusicThisLevelSec := 0             //เวลาเพลงของเลเวลนี้
	levelProceed_used := levelProceed[lvl] //ระดับ Cardio ตามเลเวล
	levelProceed_curr := 0                 //เลเวลปัจจุบันที่ขยับมา ขนาดออกกำลังกาย
	var musicList []model.Music
	// var usedSongs = make([]string, 0) // used songs
	found := false
	fmt.Println("* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * ")
	for {
		if timeRemain <= 0 {
			break
		}
		fmt.Print()
		groupedMusic := groupBy(groupMusic, levelProceed_curr)
		length := len(groupedMusic)
		if length == 0 {
			return musicList, fmt.Errorf("ไม่มีเพลงที่สามารถใช้ได้ในระดับปัจจุบัน")
		}

		idx := rand.Int() % length
		m := groupBy(groupMusic, levelProceed_curr)[idx]
		for _, song := range musicList {
			if song.Name == m.Name {
				found = true
				break
			}
		}
		if found {
			fmt.Printf("inList : %s\n", m.Name)
			continue //เพลงซ้ำ
		} else {
			musicList = append(musicList, m)
			fmt.Printf("add : %s , time : %f\n ", m.Name, m.Duration)
		}

		timeMusicSec := int(m.Duration)*60 + int(math.Round((m.Duration-math.Trunc(m.Duration))*60))
		timeRemain -= timeMusicSec
		timeMusicThisLevelSec += timeMusicSec
		percent := levelProceed_used[levelProceed_curr]
		timeExPercentSec := duration * percent / 100
		fmt.Printf("timeRemain : %d timeThisLevel:%d  levelProceed_curr:%d  percent%d  timeExPercentSec:%d  durationEx:%d\n", timeRemain, timeMusicThisLevelSec, levelProceed_curr, percent, timeExPercentSec, duration)
		if timeMusicThisLevelSec > timeExPercentSec {
			timeMusicThisLevelSec = 0 //รีเซ็ตเวลาเปอร์เซ็นของเลเวลนี้
			levelProceed_curr++       //ขยับไปเลเวลอื่น
			if levelProceed_curr > lvl {
				levelProceed_curr = 0
			}
			fmt.Printf("หลัง +- แล้ว levelProceed_curr:%d \n", levelProceed_curr)
		}
	}

	fmt.Printf("musicList:%d\n\n", len(musicList))
	return musicList, nil
}
