package main

import (
	"regexp"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
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
var DisplayPaths = make(map[string]string)

func main() {
	for k, v := range config.Paths {
		DisplayPaths["/"+k+"/"] = v[5:]
	}
	engine := html.New("./client", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Static("/style.css", "./client/style.css", fiber.Static{
		Compress:      false,
		CacheDuration: 1 * time.Millisecond,
	})
	registerQueryRoutes(app)
	registerRootRoutes(app)
	app.Listen(":" + config.Port)
}

func renderFallback(c *fiber.Ctx) error {

	return c.Render("index", fiber.Map{
		"Title": "njs.icu",
		"Sites": DisplayPaths,
	})
}
