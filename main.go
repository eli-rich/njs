package main

import (
	"regexp"

	"github.com/gofiber/fiber/v2"
)

type Config struct {
	Engine    string `json:"engine"`
	Port      string `json:"port"`
	Fullchain string `json:"fullchain"`
	Privkey   string `json:"privkey"`
	Paths     map[string]string
}

var pathRegex = regexp.MustCompile(`\/(.*)\/`)
var config Config = loadConfig()

func main() {
	app := fiber.New()
	registerQueryRoutes(app)
	registerRootRoutes(app)
	app.Listen(":" + config.Port)
}

func renderFallback(c *fiber.Ctx) error {
	return c.Redirect("https://github.com/eli-rich/njs.icu")
}
