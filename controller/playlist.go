package controller

import (
	"backend_fitfit_app/model"
	"backend_fitfit_app/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var playlistServ = service.NewPlaylistService()

func NewPlaylistController(router *gin.Engine) {
	ping := router.Group("/playlist")
	{
		ping.GET("", getAllPlaylist)
		ping.GET("/wp/:id", getAllPlaylistByWpid)
		ping.GET(":id", getPlaylistByID)
		ping.POST("/save", SavePlaylist)
		ping.PUT("/update/:id", UpdatePlaylist)
	}
}

func getAllPlaylist(ctx *gin.Context) {
	wps, err := playlistServ.GetAllPlaylist()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error": err,
		})
	}
	ctx.JSON(http.StatusOK, wps)
}
func getAllPlaylistByWpid(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	wps, err := playlistServ.GetAllPlaylistByWpid(id)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error": err,
		})
	}
	ctx.JSON(http.StatusOK, wps)
}
func getPlaylistByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	user, err := playlistServ.GetByID(id)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error": err,
		})
	}
	ctx.JSON(http.StatusOK, user)
}

func SavePlaylist(ctx *gin.Context) {
	playlist := model.Playlist{}
	ctx.ShouldBindJSON(&playlist)
	err := playlistServ.Save(playlist)
	ctx.JSON(http.StatusCreated, err)
}

func UpdatePlaylist(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	playlist := model.Playlist{}
	ctx.ShouldBindJSON(&playlist)
	err := playlistServ.Update(playlist, id)
	ctx.JSON(http.StatusOK, err)
}
