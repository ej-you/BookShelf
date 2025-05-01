package middleware

import (
	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
)

// Middleware for compressing app responses.
func Compression() fiber.Handler {
	return compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	})
}
