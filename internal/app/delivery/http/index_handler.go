package http

import (
	fiber "github.com/gofiber/fiber/v2"
)

// Handlers for user subroutes.
type indexHandler struct{}

func newIndexHandler() *indexHandler {
	return &indexHandler{}
}

// Render index HTML.
func (i indexHandler) indexHTML(ctx *fiber.Ctx) error {
	return ctx.Render("index", fiber.Map{})
}
