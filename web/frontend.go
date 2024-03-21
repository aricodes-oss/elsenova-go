//go:build !prod

package web

import (
	"net/http/httputil"
	"net/url"

	"elsenova/config"

	"github.com/gin-gonic/gin"
)

// In dev we can just directly route all 404s to the frontend
func mountFrontend(router *gin.Engine) {
	router.NoRoute(noRoute())
}

func noRoute() gin.HandlerFunc {
	conf := config.Load()
	frontendUrl, _ := url.Parse(conf.Web.Frontend)
	proxy := httputil.NewSingleHostReverseProxy(frontendUrl)

	return gin.WrapH(proxy)
}
