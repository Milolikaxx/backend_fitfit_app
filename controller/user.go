package controller

import (
	"backend_fitfit_app/model"
	"backend_fitfit_app/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var userServ = service.NewUserService()

func NewUserController(router *gin.Engine) {
	ping := router.Group("/user")
	{
		ping.GET("", getAllUser)
		ping.GET("/ByEmail", getByEmail)
		ping.GET(":id", getByID)
		ping.POST("/register", register)
		ping.POST("/login", login)
		ping.PUT("/update/:id", UpdateUser)
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

func getByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	user, err := userServ.GetUserByID(id)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error": err,
		})
	}
	ctx.JSON(http.StatusOK, user)
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

func login(ctx *gin.Context) {
	user := model.User{}
	ctx.ShouldBindJSON(&user)
	acc := userServ.Login(user)
	ctx.JSON(http.StatusOK, acc)
}

func register(ctx *gin.Context) {
	user := model.User{}
	ctx.ShouldBindJSON(&user)
	err := userServ.Register(user)
	ctx.JSON(http.StatusCreated, err)
}

func UpdateUser(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	user := model.User{}
	ctx.ShouldBindJSON(&user)
	err := userServ.Update(user, id)
	ctx.JSON(http.StatusOK, err)
}
