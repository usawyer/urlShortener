package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/usawyer/urlShortener/internal/handlers"
)

func InitRoutes(app *fiber.App, handler *handlers.Handler) {
	app.Post("/a/", handler.ShortenUrlHandler)
	app.Get("/s/:url", handler.ResolveUrlHandler)
}
