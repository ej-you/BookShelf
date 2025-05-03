package http

import (
	goerrors "errors"
	"net/http"
	"strconv"
	"strings"

	fiber "github.com/gofiber/fiber/v2"

	"BookShelf/internal/app/errors"
)

func CustomErrorHandler(ctx *fiber.Ctx, err error) error {
	// render 404 page if NotFound error
	var fiberErr *fiber.Error
	if goerrors.As(err, &fiberErr) && strings.HasPrefix(fiberErr.Message, "Cannot GET") {
		return ctx.Status(http.StatusNotFound).Render("404", fiber.Map{})
	}

	errorStatusCode := errors.CodeByError(err)
	// if unknown error
	if errorStatusCode == http.StatusInternalServerError {
		return ctx.Status(http.StatusInternalServerError).
			Render("500", fiber.Map{"message": err.Error()})
	}

	// show popup error message
	// url and query params before error occurs
	url := ctx.OriginalURL()
	queryParams := ctx.Queries()
	// add query params for redirect bask with error message
	queryParams["statusCode"] = strconv.Itoa(errorStatusCode)
	queryParams["message"] = err.Error()

	// redirect back with error message query-params
	return ctx.RedirectToRoute(url, fiber.Map{"queries": queryParams}, http.StatusSeeOther)
}
