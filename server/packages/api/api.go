package api

import (
	"minisearch/server/packages/app"

	"github.com/gofiber/fiber/v2"
)

type Settings struct{}

type API struct {
	search   app.SearchI
	settings Settings
}

func (a *API) ServiceStart() error {
	app := fiber.New()

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
