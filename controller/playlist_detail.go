package controller

import (
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
