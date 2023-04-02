package api

import (
	"fmt"

	"minisearch/server/packages/app"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Settings struct {
	Client string
}

type API struct {
	search   app.SearchI
	settings Settings
}

func (a *API) ServiceStart() error {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: fmt.Sprintf("http://localhost:3000, %s", a.settings.Client),
		AllowHeaders: "Origin, Content-Type, Accept, GET",
	}))

	app.Static("/", "./static")

	app.Get("/api/search", func(c *fiber.Ctx) error {
		query := c.Query("query")
		if query == "" {
			return c.SendStatus(500)
		}

		results := a.search.Find(query)

		if len(results) == 0 {
			return c.SendStatus(404)
		}

		return c.Status(200).JSON(results)
	})

	return app.Listen(":7777")
}

func New(settings Settings, search app.SearchI) *API {
	return &API{
		settings: settings,
		search:   search,
	}
}
