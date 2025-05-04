package http

import (
	"fmt"
	"net/http"
	"net/url"

	fiber "github.com/gofiber/fiber/v2"

	"BookShelf/internal/app/constants"
	"BookShelf/internal/app/entity"
	"BookShelf/internal/app/errors"
	"BookShelf/internal/app/usecase"
	"BookShelf/internal/pkg/cookie"
	"BookShelf/internal/pkg/validator"
)

// Handlers for user subroutes.
type userHandler struct {
	userUC        usecase.UserUsecase
	valid         validator.Validator
	cookieBuilder cookie.Builder
}

func newUserHandler(
	userUC usecase.UserUsecase,
	valid validator.Validator,
	cookieBuilder cookie.Builder) *userHandler {

	return &userHandler{
		userUC:        userUC,
		valid:         valid,
		cookieBuilder: cookieBuilder,
	}
}

// Render sign up HTML.
func (u *userHandler) signUpHTML(ctx *fiber.Ctx) error {
	return ctx.Render("signup", fiber.Map{})
}

// Render login HTML.
func (u *userHandler) loginHTML(ctx *fiber.Ctx) error {
	return ctx.Render("login", fiber.Map{})
}

// Render settings HTML.
func (u *userHandler) settingsHTML(ctx *fiber.Ctx) error {
	return ctx.Render("settings", fiber.Map{
		"login": ctx.Locals(constants.LocalsKeyLogin),
	})
}

// Sign up new user.
func (u *userHandler) signUp(ctx *fiber.Ctx) error {
	authInput := &AuthInput{}
	user := &entity.User{}

	// parse form data
	if err := ctx.BodyParser(authInput); err != nil {
		return err
	}
	// validate parsed data
	if err := u.valid.Validate(authInput); err != nil {
		return fmt.Errorf("%w: %s", errors.ErrValidateData, err.Error())
	}
	// check password and password confirm is equal
	if authInput.Password != authInput.PasswordConfirm {
		return fmt.Errorf("sign up: %w", errors.ErrConfirmPassword)
	}

	user.Login = authInput.Login
	user.Password = []byte(authInput.Password)
	// sign up new user
	userWithToken, err := u.userUC.SignUp(user)
	if err != nil {
		return fmt.Errorf("sign up: %w", err)
	}
	// set auth and login cookies
	ctx.Cookie(u.cookieBuilder.CreateCookie(constants.CookieAuth, userWithToken.AuthToken))
	ctx.Cookie(u.cookieBuilder.CreateCookie(constants.CookieLogin, user.Login))

	return ctx.Redirect(constants.SettingsPath, http.StatusSeeOther)
}

// Login existing user.
func (u *userHandler) login(ctx *fiber.Ctx) error {
	authInput := &AuthInput{}
	user := &entity.User{}

	// parse form data
	if err := ctx.BodyParser(authInput); err != nil {
		return err
	}
	// validate parsed data
	if err := u.valid.Validate(authInput); err != nil {
		return fmt.Errorf("%w: %s", errors.ErrValidateData, err.Error())
	}

	user.Login = authInput.Login
	user.Password = []byte(authInput.Password)
	// log in existing user
	userWithToken, err := u.userUC.Login(user)
	if err != nil {
		return fmt.Errorf("login: %w", err)
	}
	// set auth and login cookies
	ctx.Cookie(u.cookieBuilder.CreateCookie(constants.CookieAuth, userWithToken.AuthToken))
	ctx.Cookie(u.cookieBuilder.CreateCookie(constants.CookieLogin, user.Login))

	// redirect to the "next" (or to "library" if the "next" is not specified)
	redirectURL, err := url.QueryUnescape(ctx.Query(constants.NextQueryParam, constants.LibraryPath))
	if err != nil {
		return fmt.Errorf("login: %w", err)
	}
	return ctx.Redirect(redirectURL, http.StatusSeeOther)
}

func (u *userHandler) logout(ctx *fiber.Ctx) error {
	// clear auth and login cookies
	ctx.Cookie(u.cookieBuilder.ClearCookie(constants.CookieAuth))
	ctx.Cookie(u.cookieBuilder.ClearCookie(constants.CookieLogin))

	return ctx.Redirect("/", http.StatusSeeOther)
}
