package controller

import (
	"backend_fitfit_app/model"
	"backend_fitfit_app/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var exerServ = service.NewExerService()

func NewExerciseController(router *gin.Engine) {
	ping := router.Group("/exercise")
	{
		ping.GET("", getAllExer)
		ping.GET(":id", getExerByID)
		ping.POST("/addexercise", SaveExercise)
		ping.PUT("/edithistory/:id", UpdateExerciseHis)
	}
}

func getAllExer(ctx *gin.Context) {
	exers, err := exerServ.GetAllExer()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error": err,
		})
	}
	ctx.JSON(http.StatusOK, exers)
}

func getExerByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	exer, err := exerServ.GetExerByID(id)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error": err,
		})
	}
	ctx.JSON(http.StatusOK, exer)
}

func SaveExercise(ctx *gin.Context) {
	history := model.Exercise{}
	ctx.ShouldBindJSON(&history)
	err := exerServ.Save(history)
	ctx.JSON(http.StatusCreated, err)
}

func UpdateExerciseHis(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	exercise := model.Exercise{}
	ctx.ShouldBindJSON(&exercise)
	err, _ := exerServ.Update(exercise, id)
	ctx.JSON(http.StatusOK, err)
}