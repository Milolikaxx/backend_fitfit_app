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
		ping.GET(":id", getWpByWPID)
		ping.POST("/save", Save)
		ping.PUT("/update/:id", UpdateWorkProfile)
		ping.GET("/user/:id", getListWpByUid)
		ping.DELETE("/delprofile/:id", DeleteWorkProfile)
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

func getWpByWPID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	user, err := wpServ.GetWpByWpid(id)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error": err,
		})
	}
	ctx.JSON(http.StatusOK, user)
}

func getListWpByUid(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	user, err := wpServ.GetListWpByUid(id)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error": err,
		})
	}
	ctx.JSON(http.StatusOK, user)
}

//	func getWpByUID(ctx *gin.Context) {
//		id, _ := strconv.Atoi(ctx.Param("id"))
//		user, err := wpServ.GetWpByUID(id)
//		if err != nil {
//			ctx.JSON(http.StatusOK, gin.H{
//				"error": err,
//			})
//		}
//		ctx.JSON(http.StatusOK, user)
//	}
func Save(ctx *gin.Context) {
	wp := model.WorkoutProfile{}
	ctx.ShouldBindJSON(&wp)
	err := wpServ.Save(wp)
	ctx.JSON(http.StatusCreated, err)
}

func DeleteWorkProfile(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	wp, err := wpServ.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error": err,
		})
	}
	ctx.JSON(http.StatusOK, wp)
}

func UpdateWorkProfile(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	wp := model.WorkoutProfile{}
	ctx.ShouldBindJSON(&wp)
	err := wpServ.Update(wp, id)
	ctx.JSON(http.StatusOK, err)
}
