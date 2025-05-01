package http

import (
	fiber "github.com/gofiber/fiber/v2"

	"BookShelf/internal/app/usecase"
)

// Manager for user subroutes handlers
type userHandler struct {
	userUC usecase.UserUsecase
}

func newUserHandler(userUC usecase.UserUsecase) *userHandler {
	return &userHandler{
		userUC: userUC,
	}
}

// Render login HTML.
func (u userHandler) loginHTML(ctx *fiber.Ctx) error {
	return ctx.Render("login", fiber.Map{})
}

// Render sign up HTML.
func (u userHandler) signUpHTML(ctx *fiber.Ctx) error {
	return ctx.Render("signup", fiber.Map{})
}
