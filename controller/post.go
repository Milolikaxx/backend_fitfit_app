package controller

import (
	"backend_fitfit_app/model"
	"backend_fitfit_app/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var postServ = service.NewPostService()

func NewPostController(router *gin.Engine) {
	ping := router.Group("/post")
	{
		ping.GET("", getAllPosts)
		ping.GET(":id", getPostByID)
		ping.POST("/save", SavePost)
		ping.PUT("/update/:id", UpdatePost)
	}
}

func getAllPosts(ctx *gin.Context) {
	wps, err := postServ.GetAllPosts()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error": err,
		})
	}
	ctx.JSON(http.StatusOK, wps)
}

func getPostByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	user, err := postServ.GetPostByID(id)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error": err,
		})
	}
	ctx.JSON(http.StatusOK, user)
}

func SavePost(ctx *gin.Context) {
	post := model.Post{}
	ctx.ShouldBindJSON(&post)
	err := postServ.Save(post)
	ctx.JSON(http.StatusCreated, err)
}

func UpdatePost(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	post := model.Post{}
	ctx.ShouldBindJSON(&post)
	err := postServ.Update(post, id)
	ctx.JSON(http.StatusOK, err)
}
