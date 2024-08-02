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
		ping.GET("/searchbyday", FindByDay)
		ping.GET("/last7day", GetLast7Day)
		ping.GET("/getmonth", Get12Month)
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

func FindByDay(ctx *gin.Context) {
	// สร้างตัวแปรเพื่อเก็บข้อมูลที่รับเข้ามา
	var requestData struct {
		Keyword string `json:"keyword"`
	}

	// ดึงข้อมูล JSON จาก request body และแปลงเป็น struct
	if err := ctx.ShouldBindJSON(&requestData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// เรียกใช้งาน service เพื่อค้นหาข้อมูล
	exercises, err := exerServ.SearchByDay(requestData.Keyword)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// ส่งข้อมูลกลับในรูป JSON
	ctx.JSON(http.StatusOK, exercises)
}

// func GetLast7Day(ctx *gin.Context) {
// 	exers, err := exerServ.ExerciseLast7Day()
// 	if err != nil {
// 		ctx.JSON(http.StatusOK, gin.H{
// 			"error": err,
// 		})
// 	}
// 	ctx.JSON(http.StatusOK, exers)
// }

// func GetLast7Day(ctx *gin.Context) {
// 	exers, err := exerServ.ExerciseLast7Day()
// 	if err != nil {
// 		ctx.JSON(http.StatusOK, gin.H{
// 			"error": err,
// 		})
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, exers)
// }

func GetLast7Day(ctx *gin.Context) {
	exers, err := exerServ.ExerciseLast7Day()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error": err,
		})
		return
	}
	ctx.JSON(http.StatusOK, exers)
}

func Get12Month(ctx *gin.Context) {
	exers, err := exerServ.GetLast12Months()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, exers)
}
