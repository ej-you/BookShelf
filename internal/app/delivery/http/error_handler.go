package http

import (
	goerrors "errors"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	fiber "github.com/gofiber/fiber/v2"

	"BookShelf/internal/app/constants"
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
	// query params before error occurs
	urlQuery := url.Values{}
	for k, v := range ctx.Queries() {
		urlQuery.Add(k, v)
	}
	urlQuery.Set("statusCode", strconv.Itoa(errorStatusCode))
	urlQuery.Set("message", err.Error())

	// parse last GET-request URL
	redirectURL := ctx.Path()
	// set "/genre" path for all "/genre" subroutes
	if strings.HasPrefix(redirectURL, constants.GenrePath) {
		redirectURL = constants.GenrePath
	}
	// add query params to URL
	redirectURL = redirectURL + "?" + urlQuery.Encode()

	// redirect back (last GET-request URL) with error message query-params
	return ctx.Redirect(redirectURL, http.StatusSeeOther)
}
