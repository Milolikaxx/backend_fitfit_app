package controller

import (
	"backend_fitfit_app/model"
	"backend_fitfit_app/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var musichistoryServ = service.NewMusicHistoryService()

func NewMusicHistoryController(router *gin.Engine) {
	ping := router.Group("/musichistory")
	{
		ping.GET("", getAllMusicHistory)
		ping.GET("/:id", getAllMusicHistoryByEid)
		ping.POST("/addmusichistory", InsertMusicHistory)
	}
}

func getAllMusicHistory(ctx *gin.Context) {
	musichis, err := musichistoryServ.GetAllMusicHistory()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error": err,
		})
	}
	ctx.JSON(http.StatusOK, musichis)
}

func getAllMusicHistoryByEid(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	musichis, err := musichistoryServ.GetAllMusicHistoryByEid(id)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error": err,
		})
	}
	ctx.JSON(http.StatusOK, musichis)
}

func InsertMusicHistory(ctx *gin.Context) {
	musichis := model.MusicHistory{}
	ctx.ShouldBindJSON(&musichis)
	err := musichistoryServ.AddMusicHistory(musichis)
	ctx.JSON(http.StatusCreated, err)
}
