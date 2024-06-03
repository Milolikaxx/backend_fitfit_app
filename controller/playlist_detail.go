package controller

import (
	"backend_fitfit_app/model"
	"backend_fitfit_app/repository"
	"backend_fitfit_app/service"
	"fmt"
	"math"
	"math/rand"
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
	music := CreatePlaylistWP(id)
	ctx.JSON(http.StatusOK, music)

}

var bpm = []int{114, 133, 152, 171, 190}
var levelProceed = [][]int{
	{100},                //Lv.1
	{10, 35},             //Lv.2
	{10, 15, 35},         //Lv.3
	{10, 10, 15, 35},     //Lv.4
	{10, 10, 10, 15, 35}, //Lv.5
}

type MusicGroup struct {
	level int
	music model.Music
}

func groupBy(maps []MusicGroup, level int) []model.Music {
	groups := make([]model.Music, 0)
	for _, m := range maps {
		if m.level == level {
			groups = append(groups, m.music)
		}
	}
	return groups
}

func CreatePlaylistWP(wpid int) []model.Music {
	//workout_profile
	var wpRepo = repository.NewWpRepository()
	wp, _ := wpRepo.FindByWpid(wpid)
	fmt.Printf("wp:%v  \n\n", wp)
	//infomation
	lvl := wp.LevelExercise
	duration := int(wp.Duration * 60)
	exeType := wp.ExerciseType
	var musicType []int
	for _, t := range wp.WorkoutMusictype {
		musicType = append(musicType, t.Mtid)
	}
	fmt.Printf("infomation  lvl:%d  duration:%d  type:%s  musicType:%v  bpm:%d  \n\n", lvl, duration, exeType, musicType, bpm[lvl-1])

	//music - level
	musicRepo := repository.NewMusicRepository()
	music, _ := musicRepo.FindAllMusicByLevel(bpm[lvl-1], musicType)
	fmt.Printf("result music:%d\n\n", len(music))

	//grouping
	fmt.Printf("music type %s\n\n", reflect.TypeOf(music))
	var groupMusic []MusicGroup
	for i := 0; i < lvl; i++ {
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

	fmt.Printf("level1 = %d\n", len(groupBy(groupMusic, 0)))
	fmt.Printf("level2 = %d\n", len(groupBy(groupMusic, 1)))
	fmt.Printf("level3 = %d\n", len(groupBy(groupMusic, 2)))
	fmt.Printf("level4 = %d\n", len(groupBy(groupMusic, 3)))
	fmt.Printf("level5 = %d\n\n", len(groupBy(groupMusic, 4)))

	//proceed
	// fmt.Println(levelProceed[2])
	timeRemain := duration                 //เวลาที่ยังเหลือ
	timeThisLevel := 0                     //เวลาเปอร์เซ็น ของเลเวลนี้
	levelProceed_used := levelProceed[lvl] //ระดับ Cardio ตามเลเวล
	levelProceed_curr := 0                 //เลเวลที่ตรงกับเวลาตอนนี้
	var musicList []model.Music
	fmt.Println("* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * ")
	for {
		if timeRemain <= 0 {
			break
		}
		idx := rand.Int() % len(groupBy(groupMusic, levelProceed_curr))
		m := groupBy(groupMusic, levelProceed_curr)[idx]
		musicList = append(musicList, m)
		t := int(int(m.Duration)*60) + int(math.Round((m.Duration-math.Trunc(m.Duration))*100))
		timeRemain -= t
		timeThisLevel += t
		percent := levelProceed_used[levelProceed_curr]
		timeChange := duration * percent / 100
		// fmt.Printf("timeThisLevel:%d  levelProceed_curr:%d  percent%d  timeChange:%d  duration:%d\n", timeThisLevel, levelProceed_curr, percent, timeChange, duration)
		if timeThisLevel > timeChange {
			timeThisLevel = 0   //รีเซ็ตเวลาเปอร์เซ็นของเลเวลนี้
			levelProceed_curr++ //ขยับเลือกเลเวลอื่น
			if levelProceed_curr >= lvl {
				levelProceed_curr = 0
			}
		}
	}
	fmt.Printf("musicList:%d\n\n", len(musicList))
	return musicList
}
