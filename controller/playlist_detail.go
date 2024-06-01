package controller

import (
	"backend_fitfit_app/model"
	"backend_fitfit_app/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var playlistDetailServ = service.NewPlaylistDetailService()

func NewPlaylistDetailController(router *gin.Engine) {
	ping := router.Group("/playlist_detail")
	{
		ping.GET("", getAllPlaylistDetail)
		ping.GET("/:id", getListWpByPID)
		ping.POST("/addmusic", AddMusic)
		ping.DELETE("/delete/:id", DeleteMusic)
	}
}

func getAllPlaylistDetail(ctx *gin.Context) {
	pl_detail, err := playlistDetailServ.GetAllPlaylistDetail()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error": err,
		})
	}
	ctx.JSON(http.StatusOK, pl_detail)
}

func getListWpByPID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	pl_detail, err := playlistDetailServ.GetListWpByPID(id)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error": err,
		})
	}
	ctx.JSON(http.StatusOK, pl_detail)
}

func AddMusic(ctx *gin.Context) {
	pld := model.PlaylistDetail{}
	ctx.ShouldBindJSON(&pld)
	err := playlistDetailServ.Save(pld)
	ctx.JSON(http.StatusCreated, err)
}

func DeleteMusic(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	pl_detail, err := playlistDetailServ.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error": err,
		})
	}
	ctx.JSON(http.StatusOK, pl_detail)
}
