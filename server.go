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
	router.Run()
}
