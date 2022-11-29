package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nanmu42/gzip"
)

type Config struct {
	Engine     string            `json:"engine"`
	Port       string            `json:"port"`
	Paths      map[string]string `json:"paths"`
	Standalone map[string]string `json:"standalone"`
}

var config Config = loadConfig()
var DisplayPaths = make(map[string]string)

func main() {
	for k, v := range config.Paths {
		DisplayPaths["/"+k+"/"] = v[5:]
	}
	for k, v := range config.Standalone {
		DisplayPaths["/"+k+"/"] = v[5:]
	}
	gin.SetMode(gin.ReleaseMode)
	app := gin.Default()
	app.Use(gzip.DefaultHandler().Gin)
	app.LoadHTMLGlob("client/*.html")
	app.StaticFile("/style.css", "./client/style.css")
	app.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"Title": "njs.icu",
			"Sites": DisplayPaths,
		})
	})
	registerGoogleRoute(app)
	registerQueryRoutes(app)
	registerRootRoutes(app)
	app.Run(":" + config.Port)
}
