package web

import (
	"elsenova/web/controllers"

	"github.com/gin-gonic/gin"
)

var (
	vore    = &controllers.VoreController{}
	discord = &controllers.DiscordController{}
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("api")

	// Mount controllers
	vore.Mount(api)
	discord.Mount(api)

	// Proxy all other routes to the frontend
	mountFrontend(router)

	return router
}
