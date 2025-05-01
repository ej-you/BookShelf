package middleware

import (
	fiber "github.com/gofiber/fiber/v2"
	logger "github.com/gofiber/fiber/v2/middleware/logger"
)

// Middleware for logging all request-response chains.
func Logger() fiber.Handler {
	return logger.New(logger.Config{
		TimeFormat: "2006-01-02T15:04:05-0700",
		TimeZone:   "UTC",
		Format:     "${time} | ${status} | ${latency} | ${method} | ${path} | ${error}\n",
	})
}
