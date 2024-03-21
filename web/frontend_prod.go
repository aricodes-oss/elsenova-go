//go:build prod

package web

import (
	"embed"
	"net/http"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

//go:embed dist
var slug embed.FS

// In prod we need to redirect to the embed folder route
func mountFrontend(router *gin.Engine) {
	staticFiles := static.Serve("/", static.EmbedFolder(slug, "dist"))
	router.Use(staticFiles)
	router.NoRoute(noRoute)
}

func noRoute(c *gin.Context) {
	data, _ := slug.ReadFile("dist/index.html")
	c.Data(http.StatusOK, "text/html", data)
	return
}
