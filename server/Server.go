package server

import (
	"GoSvelteBoilerplate/server/helpers"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func Serve(hostAndPort string) {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()

	engine.Use(static.Serve("/", static.LocalFile("./dist", false)))

	// TODO: Add custom 404 page
	engine.NoRoute(func(c *gin.Context) {
		log.WithField("path", c.Request.URL.Path).Debug("calling unknown path")
		c.File("./dist/index.html")
	})

	apiGroup := engine.Group("/api")
	registerRoutes(apiGroup)

	log.WithFields(log.Fields{"address": hostAndPort}).Info("Starting server")
	engine.Run(hostAndPort)
}

func registerRoutes(apiGroup *gin.RouterGroup) {
	apiGroup.GET("/", func(c *gin.Context) {
		c.Writer.WriteHeader(http.StatusOK)
		c.Writer.WriteString("Hello World")
	})

	apiGroup.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"uptime": helpers.Uptime().Seconds(),
			"status": "ok",
		})
	})

}
