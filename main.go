package main

import (
	"encoding/json"
	"log"
	"os"
	"regexp"

	"github.com/gofiber/fiber/v2"
)

type Config struct {
	Engine    string `json:"engine"`
	Fallback  string `json:"fallback"`
	PortSSL   string `json:"portSSL"`
	Port      string `json:"port"`
	Fullchain string `json:"fullchain"`
	Privkey   string `json:"privkey"`
	Paths     map[string]string
}

func main() {
	configFile, err := os.ReadFile("config.json")
	if os.IsNotExist(err) {
		log.Fatalln("config.json not found")
	} else if err != nil {
		log.Fatalln(err)
	}

	var config Config
	if err := json.Unmarshal(configFile, &config); err != nil {
		log.Fatalln(err)
	}
	pathRegex := regexp.MustCompile(`\/(.*)\/`)
	if err != nil {
		log.Fatalln(err)
	}
	app := fiber.New()
	app.Get("/*/:query", func(c *fiber.Ctx) error {
		results := pathRegex.FindStringSubmatch(c.Path())
		if len(results) == 0 {
			return c.Redirect(config.Fallback)
		}
		path := results[1]
		search := config.Paths[path]
		if search == "" {
			return c.Redirect(config.Fallback)
		}
		return c.Redirect(config.Engine + c.Params("query") + " " + search)
	})
	app.Get("/*", func(c *fiber.Ctx) error {
		results := pathRegex.FindStringSubmatch(c.Path())
		if len(results) == 0 {
			return c.Redirect(config.Fallback)
		}
		path := results[1]
		search := config.Paths[path]
		if search == "" {
			return c.Redirect(config.Fallback)
		}
		return c.Redirect("https://" + search[5:])
	})
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect(config.Fallback)
	})
	go forwardToHttps(config.Port)
	app.ListenTLS(":"+config.PortSSL, config.Fullchain, config.Privkey)
}

func forwardToHttps(port string) {
	app := fiber.New()
	app.Get("/*", func(c *fiber.Ctx) error {
		return c.Redirect("https://" + c.Hostname() + c.Path())
	})
	app.Listen(":" + port)
}
