package controller

import (
	"backend_fitfit_app/model"
	"backend_fitfit_app/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

var userServ = service.NewUserService()

func NewUserController(router *gin.Engine) {
	ping := router.Group("/user")
	{
		ping.GET("", getAllUser)
		ping.GET("/ByEmail", getByEmail)
		ping.POST("/register", registerAccount)
	}
}

func getAllUser(ctx *gin.Context) {
	// users := []model.User{}
	users, err := userServ.GetAllUsers()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error": err,
		})
	}
	ctx.JSON(http.StatusOK, users)
}

func getByEmail(ctx *gin.Context) {
	user := &model.User{}
	ctx.ShouldBindJSON(&user)
	user, err := userServ.GetUserByEmail(user.Email)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error": err,
		})
	}
	ctx.JSON(http.StatusOK, user)
}

func registerAccount(ctx *gin.Context) {
	user := model.User{}
	ctx.ShouldBindJSON(&user)
	err := userServ.Register(user)
	ctx.JSON(http.StatusOK, err)
}
