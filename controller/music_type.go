package controller

import (
	"backend_fitfit_app/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var mtServ = service.NewMtService()

func NewMusicTypeController(router *gin.Engine) {
	ping := router.Group("/mt")
	{
		ping.GET("", getAllMt)
		ping.GET(":id", getMtByID)

	}
}

func getAllMt(ctx *gin.Context) {
	mts, err := mtServ.GetAllMt()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error": err,
		})
	}
	ctx.JSON(http.StatusOK, mts)
}

func getMtByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	mt, err := mtServ.GetMtByID(id)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error": err,
		})
	}
	ctx.JSON(http.StatusOK, mt)
}
