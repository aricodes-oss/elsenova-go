package web

import (
	"net/http/httputil"
	"net/url"

	"elsenova/config"
	"elsenova/web/controllers"

	"github.com/gin-gonic/gin"
)

var (
	vore    = &controllers.VoreController{}
	discord = &controllers.DiscordController{}
)

func NewRouter() *gin.Engine {
	conf := config.Load()
	router := gin.Default()
	api := router.Group("api")

	// Mount controllers
	vore.Mount(api)
	discord.Mount(api)

	// Proxy all other routes to the frontend
	// TODO: add production mode
	frontendUrl, _ := url.Parse(conf.Web.Frontend)
	proxy := httputil.NewSingleHostReverseProxy(frontendUrl)
	router.NoRoute(gin.WrapH(proxy))

	return router
}
