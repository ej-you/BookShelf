package middleware

import (
	"fmt"
	"net/http"
	"net/url"

	fiber "github.com/gofiber/fiber/v2"

	"BookShelf/internal/app/constants"
)

// Parsing access token and email from cookies to context
func CookieParser() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// parse cookies from request
		authToken := ctx.Cookies(constants.CookieAuth)
		login := ctx.Cookies(constants.CookieLogin)

		// set vars to context
		ctx.Locals(constants.LocalsKeyAuthToken, authToken)
		ctx.Locals(constants.LocalsKeyLogin, login)
		return ctx.Next()
	}
}

// Redirect to login if cookies are not specified.
func ToLoginIfNoCookie() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// parse cookies from request
		accessToken := ctx.Cookies(constants.CookieAuth)

		// if cookies is not specified
		if accessToken == "" {
			// send next param (current url before redirect to login)
			redirectURL := fmt.Sprintf(
				"%s?%s=%s",
				constants.LoginPath,
				constants.NextQueryParam,
				url.QueryEscape(ctx.OriginalURL()),
			)
			return ctx.Redirect(redirectURL, http.StatusSeeOther)
		}
		return ctx.Next()
	}
}

// Redirect to settings if cookies are specified.
func ToSettingsIfCookie() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// parse cookies from request
		authToken := ctx.Cookies(constants.CookieAuth)

		// if cookies is specified
		if authToken != "" {
			return ctx.Redirect(constants.SettingsPath, http.StatusSeeOther)
		}
		return ctx.Next()
	}
}
