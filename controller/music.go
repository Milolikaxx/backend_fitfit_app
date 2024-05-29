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
	}
}

func getMusicByWtid(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	user, err := musicServ.GetMusicByMtid(id)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error": err,
		})
	}
	ctx.JSON(http.StatusOK, user)
}
