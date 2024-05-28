package controller

import (
	"backend_fitfit_app/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

var playlistDetailServ = service.NewPlaylistDetailService()

func NewPlaylistDetailController(router *gin.Engine) {
	ping := router.Group("/playlist_detail")
	{
		ping.GET("", getAllPlaylistDetail)
		// ping.GET(":id", getPlaylistByID)
		// ping.POST("/save", SavePlaylist)
		// ping.PUT("/update/:id", UpdatePlaylist)
	}
}

func getAllPlaylistDetail(ctx *gin.Context) {
	wps, err := playlistDetailServ.GetAllPlaylistDetail()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error": err,
		})
	}
	ctx.JSON(http.StatusOK, wps)
}
