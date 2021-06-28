package handlers

import "github.com/gofiber/fiber/v2"

type AuthHandler struct {
}

func (a *AuthHandler) InitHandler() {
}

func (a *AuthHandler) Login(c *fiber.Ctx) error {

	return c.JSON(struct {
		msg    string
		status int
	}{
		msg:    "login successfully.",
		status: 200,
	})
}
