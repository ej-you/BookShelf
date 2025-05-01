package middleware

import (
	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// Middleware for panic recovery for continuous work.
func Recover() fiber.Handler {
	return recover.New(recover.Config{
		EnableStackTrace: true,
	})
}
