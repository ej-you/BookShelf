// Package http contains HTTP router registrators for
// every entity with its usecase. Also this package
// contains HTTP handlers for handling router endpoints.
package http

import (
	fiber "github.com/gofiber/fiber/v2"

	"BookShelf/internal/app/usecase"
)

func RegisterUserEndpoints(
	router fiber.Router,
	userUC usecase.UserUsecase,
	mwToSettingsIfCookie fiber.Handler) {

	userHandler := newUserHandler(userUC)
	userGroup := router.Group("/user")
	{
		userGroup.Get("/login", mwToSettingsIfCookie, userHandler.loginHTML)
		userGroup.Get("/sign-up", mwToSettingsIfCookie, userHandler.signUpHTML)
		// userGroup.Post("/login", userHandler.login)
		// userGroup.Post("/sign-up", userHandler.signUp)
	}
}
