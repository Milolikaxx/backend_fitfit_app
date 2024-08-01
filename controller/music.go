package controller

import (
	"backend_fitfit_app/model"
	"backend_fitfit_app/repository"
	"backend_fitfit_app/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var musicServ = service.NewMusicService()

func NewMusicController(router *gin.Engine) {
	ping := router.Group("/music")
	{
		ping.GET(":id", getMusicByWtid)
		ping.GET("/findbywp/:id", getMusic)
		ping.GET("/search", getSearchMusic)
		ping.GET("/findmusicaddsong", findMusicForAddSongPL)
	}
}

func getMusicByWtid(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	music, err := musicServ.GetMusicByMtid(id)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error": err,
		})
	}
	ctx.JSON(http.StatusOK, music)
}
func getMusic(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	music, err := musicPlaylistWP(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, music)

}
func getSearchMusic(ctx *gin.Context) {
	musicKey := model.SearchMusic{}
	ctx.ShouldBindJSON(&musicKey)
	music, err := musicServ.SearchMusic(musicKey)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, music)
}

var Bpm = []int{100, 114, 133, 152, 171, 190}

func musicPlaylistWP(wpid int) ([]model.Music, error) {
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

	music, _ := musicServ.GetMusicByLevel(Bpm[lvl], musicType)
	fmt.Printf("result music:%d\n\n", len(music))
	return music, nil
}

func findMusicForAddSongPL(ctx *gin.Context) {
	data := model.RandMusicOfPlaylist{}
	ctx.ShouldBindJSON(&data)
	music, err := findMusicForAddSongOfPlaylist(data)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, music)
}

func findMusicForAddSongOfPlaylist(data model.RandMusicOfPlaylist) ([]model.Music, error) {
	//workout_profile
	var wpRepo = repository.NewWpRepository()
	wp, _ := wpRepo.FindByWpid(data.Wpid)
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

	// Find a new song
	originalSong := data.PlaylistDetail[data.Index]
	minBPM := int(float64(originalSong.Music.Bpm) * 0.95)
	maxBPM := int(float64(originalSong.Music.Bpm) * 1.05)

	var filteredMusic []model.Music
	for _, m := range music {
		if m.Bpm >= minBPM && m.Bpm <= maxBPM {
			filteredMusic = append(filteredMusic, m)
		}
	}
	fmt.Printf("filtered music:%d\n\n", len(filteredMusic))

	return filteredMusic, nil
}
