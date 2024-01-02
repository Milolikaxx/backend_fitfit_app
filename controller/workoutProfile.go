package controller

import (
	"backend_fitfit_app/model"
	"backend_fitfit_app/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var wpServ = service.NewWpService()

func NewWpController(router *gin.Engine) {
	ping := router.Group("/workprofile")
	{
		ping.GET("", getAllWp)
		ping.GET(":id", getWpByID)
		ping.POST("/save", Save)
		// ping.PUT("/update/:id", Update)
	}
}

func getAllWp(ctx *gin.Context) {
	wps, err := wpServ.GetAllWps()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error": err,
		})
	}
	ctx.JSON(http.StatusOK, wps)
}
func getWpByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	user, err := wpServ.GetWpByID(id)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error": err,
		})
	}
	ctx.JSON(http.StatusOK, user)
}
func Save(ctx *gin.Context) {
	wp := model.WorkoutProfile{}
	ctx.ShouldBindJSON(&wp)
	err := wpServ.Save(wp)
	ctx.JSON(http.StatusCreated, err)
}

// func Update(ctx *gin.Context) {
// 	id, _ := strconv.Atoi(ctx.Param("id"))
// 	wp := model.WorkoutProfile{}
// 	ctx.ShouldBindJSON(&wp)
// 	err := wpServ.Update(id, wp)
// 	ctx.JSON(http.StatusOK, err)
// }
