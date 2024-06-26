package controller

import (
	"backend_fitfit_app/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var musicServ = service.NewMusicService()

func NewMusicController(router *gin.Engine) {
	ping := router.Group("/music")
	{
		ping.GET(":id", getMusicByWtid)
		ping.GET("/random/:id", getRandomMusicByWtid)
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
	// mtid := []model.MusictypeId{}
	// ctx.ShouldBindJSON(&mtid)
	// for _, id := range mtid {
	// 	music, err := musicServ.GetMusicByMtid(id)
	// 	if err != nil {
	// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 		return
	// 	}
	// 	ctx.JSON(http.StatusOK, music)
	// }

}

func getRandomMusicByWtid(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	music, err := musicServ.GetRandomMusicByMtid(id)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error": err,
		})
	}
	ctx.JSON(http.StatusOK, music)
}
