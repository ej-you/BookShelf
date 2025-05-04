// Package http contains HTTP router registrators for
// every entity with its usecase. Also this package
// contains HTTP handlers for handling router endpoints.
package http

import (
	fiber "github.com/gofiber/fiber/v2"

	"BookShelf/internal/app/middleware"
	"BookShelf/internal/app/usecase"
	"BookShelf/internal/pkg/cookie"
	"BookShelf/internal/pkg/validator"
)

func RegisterIndexEndpoints(router fiber.Router) {
	indexHandler := newIndexHandler()
	router.Get("/", indexHandler.indexHTML)
}

func RegisterUserEndpoints(
	router fiber.Router,
	userUC usecase.UserUsecase,
	valid validator.Validator,
	cookieBuilder cookie.Builder) {

	userHandler := newUserHandler(userUC, valid, cookieBuilder)
	userGroup := router.Group("/user")
	{
		userGroup.Get("/sign-up", middleware.ToSettingsIfCookie(), userHandler.signUpHTML)
		userGroup.Post("/sign-up", userHandler.signUp)

		userGroup.Get("/login", middleware.ToSettingsIfCookie(), userHandler.loginHTML)
		userGroup.Post("/login", userHandler.login)

		userGroup.Get("/settings",
			middleware.ToLoginIfNoCookie(), middleware.CookieParser(),
			userHandler.settingsHTML)
		userGroup.Post("/logout", userHandler.logout)
	}
}

func RegisterGenreEndpoints(
	router fiber.Router,
	genreUC usecase.GenreUsecase,
	valid validator.Validator) {

	genreHandler := newGenreHandler(genreUC, valid)
	genreGroup := router.Group("/genre", middleware.ToLoginIfNoCookie())
	{
		genreGroup.Get("/", genreHandler.listHTML)
		genreGroup.Post("/create", genreHandler.create)
		genreGroup.Post("/remove/:genreID", genreHandler.remove)
	}
}
