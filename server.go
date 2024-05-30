package main

import (
	"backend_fitfit_app/controller"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	router := gin.Default()
	controller.NewUserController(router)
	controller.NewWpController(router)
	controller.NewPlaylistController(router)
	controller.NewMusicTypeController(router)
	controller.NewWpMusicTypeController(router)
	controller.NewPostController(router)
	controller.NewPlaylistDetailController(router)
	// controller.NewMusicController(router)
	router.Run()
}
