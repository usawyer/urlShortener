package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/usawyer/urlShortener/internal/service"
	"github.com/usawyer/urlShortener/models"
)

type Handler struct {
	s *service.Service
}

func New(s *service.Service) *Handler {
	return &Handler{s: s}
}

func (h *Handler) ShortenUrlHandler(c *fiber.Ctx) error {
	longUrl := c.Query("url", "")

	if longUrl == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": "no URL found"})
	}

	alias, err := h.s.ShortenUrl(c.Context(), longUrl)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(models.Response{Alias: alias})
}

func (h *Handler) ResolveUrlHandler(c *fiber.Ctx) error {
	alias := c.Params("url")

	longUrl, err := h.s.ResolveUrl(c.Context(), alias)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"Error": err.Error()})
	}

	return c.Redirect(longUrl, fiber.StatusFound)
}
