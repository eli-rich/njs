package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

type Config struct {
	Engine   string `json:"engine"`
	Search   string `json:"search"`
	Fallback string `json:"fallback"`
	Port     string `json:"port"`
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
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect(config.Fallback)
	})
	app.Get("/:query", func(c *fiber.Ctx) error {
		return c.Redirect(config.Engine + c.Params("query") + " " + config.Search)
	})
	app.Listen(":" + config.Port)
}
