package controller

import (
	"backend_fitfit_app/model"
	"backend_fitfit_app/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var wpmtServ = service.NewWpMusicTypeService()

func NewWpMusicTypeController(router *gin.Engine) {
	ping := router.Group("/wpmt")
	{

		ping.GET(":id", getByWPID)
		ping.POST("/save", SaveWpMusicType)
		ping.PUT("/update/:id", UpdateWpMusicType)
	}
}

func getByWPID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	user, err := wpmtServ.GetByWPID(id)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error": err,
		})
	}
	ctx.JSON(http.StatusOK, user)
}

func SaveWpMusicType(ctx *gin.Context) {
	wpmt := model.WorkoutMusictype{}
	ctx.ShouldBindJSON(&wpmt)
	err := wpmtServ.Save(wpmt)
	ctx.JSON(http.StatusCreated, err)
}

func UpdateWpMusicType(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	wpmt := model.WorkoutMusictype{}
	ctx.ShouldBindJSON(&wpmt)
	err := wpmtServ.Update(wpmt, id)
	ctx.JSON(http.StatusOK, err)
}
