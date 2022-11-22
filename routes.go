package main

import "github.com/gofiber/fiber/v2"

func registerQueryRoutes(app *fiber.App) {
	// routes for paths with query
	// ex: https://njs.icu/wiki/Alan-Turing
	app.Get("/*/:query", func(c *fiber.Ctx) error {
		results := pathRegex.FindStringSubmatch(c.Path())
		if len(results) == 0 {
			return renderFallback(c)
		}
		path := results[1]
		search := config.Paths[path]
		if search == "" {
			return renderFallback(c)
		}
		return c.Redirect(config.Engine + c.Params("query") + " " + search)
	})
}

func registerRootRoutes(app *fiber.App) {
	// routes for paths without query
	// ex: https://njs.icu/wiki
	app.Get("/*", func(c *fiber.Ctx) error {
		results := pathRegex.FindStringSubmatch(c.Path())
		if len(results) == 0 {
			return renderFallback(c)
		}
		path := results[1]
		search := config.Paths[path]
		if search == "" {
			return renderFallback(c)
		}
		return c.Redirect("https://" + search[5:])
	})
}

func registerGoogleRoute(app *fiber.App) {
	app.Get("/g/:query", func(c *fiber.Ctx) error {
		return c.Redirect("https://google.com/search?q=" + c.Params("query"))
	})
}
