package web

import (
	"net/http/httputil"
	"net/url"

	"elsenova/config"

	"github.com/gin-gonic/gin"
)

func mountFrontend(router *gin.Engine) {
	router.NoRoute(noRoute())
}

func noRoute() gin.HandlerFunc {
	conf := config.Load()
	frontendUrl, _ := url.Parse(conf.Web.Frontend)
	proxy := httputil.NewSingleHostReverseProxy(frontendUrl)

	return gin.WrapH(proxy)
}
