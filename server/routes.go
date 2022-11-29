package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

type Search struct {
	Engine string `uri:"engine" binding:"required"`
	Query  string `uri:"query" binding:"required"`
}

func registerQueryRoutes(app *gin.Engine) {
	// routes for paths with query
	// ex: https://njs.icu/wiki/Alan-Turing
	app.GET("/:engine/:query", func(c *gin.Context) {
		path := c.Param("engine")
		search := config.Paths[path]
		if search == "" {
			c.Redirect(302, "/")
			return
		}
		c.Redirect(302, config.Engine+c.Param("query")+" "+search)
	})
}

func registerRootRoutes(app *gin.Engine) {
	// routes for paths without query
	// ex: https://njs.icu/wiki
	app.GET("/:engine", func(c *gin.Context) {
		path := c.Param("engine")
		search := config.Paths[path]
		if search == "" {
			// check if standalone url
			standalone := config.Standalone[path][5:]
			log.Println(standalone)
			if standalone == "" {
				c.Redirect(302, "/")
				return
			}
			c.Redirect(302, "https://"+standalone)
		}
		c.Redirect(302, "https://"+search[5:])
	})
}

func registerGoogleRoute(app *gin.Engine) {
	app.GET("/g/:query", func(c *gin.Context) {
		c.Redirect(302, "https://google.com/search?q="+c.Param("query"))
	})
}
