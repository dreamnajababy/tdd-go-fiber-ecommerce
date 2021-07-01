package handlers

import (
	helper "github.com/dreamnajababy/go-ecom/src/helper"
	repo "github.com/dreamnajababy/go-ecom/src/repositories"
	"github.com/gofiber/fiber/v2"
)

type Credential struct {
	Username string
	Password string
}

type AuthHandler struct {
	UserRepository repo.UserRepository
}

func (a *AuthHandler) InitHandler(ur repo.UserRepository) {
	a.UserRepository = ur
}

func (a *AuthHandler) Login(c *fiber.Ctx) error {
	var credential Credential
	c.BodyParser(&credential)

	user, err := a.UserRepository.Login(credential.Username, credential.Password)

	if err != nil {
		return c.Status(401).JSON(fiber.Map{"msg": err.Error(), "statusCode": 401})
	}

	token, err := helper.GetUserToken(user)

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"msg": "login successfully.", "statusCode": 200, "token": token})
}
