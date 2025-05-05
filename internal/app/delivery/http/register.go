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
	cookieBuilder cookie.Builder,
	valid validator.Validator) {

	userHandler := newUserHandler(userUC, cookieBuilder, valid)
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

func RegisterBookEndpoints(
	router fiber.Router,
	bookUC usecase.BookUsecase,
	genreUC usecase.GenreUsecase,
	tokenSigningKey []byte,
	mediaPath string,
	valid validator.Validator) {

	bookHandler := newBookHandler(bookUC, genreUC, tokenSigningKey, mediaPath, valid)
	router.Get("/library",
		middleware.CookieParser(), middleware.ToLoginIfNoCookie(),
		bookHandler.listHTML)
	bookGroup := router.Group("/book", middleware.CookieParser(), middleware.ToLoginIfNoCookie())
	{
		bookGroup.Get("/create", bookHandler.createHTML)
		bookGroup.Post("/create", bookHandler.create)

		bookGroup.Get("/edit/:genreID", bookHandler.editHTML)
		bookGroup.Post("/edit/:genreID", bookHandler.edit)

		bookGroup.Post("/export/excel", bookHandler.exportExcel)

		bookGroup.Post("/remove/:genreID", bookHandler.remove)
	}
}
