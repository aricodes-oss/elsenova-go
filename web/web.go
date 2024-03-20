package web

import (
	"elsenova/web/controllers"

	"github.com/gin-gonic/gin"
)

var (
	vore = &controllers.VoreController{}
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	api := router.Group("api")

	// Mount controllers
	vore.Mount(api)

	return router
}
