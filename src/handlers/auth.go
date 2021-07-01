package handlers

import (
	"time"

	repo "github.com/dreamnajababy/go-ecom/src/repositories"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
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
	err := c.BodyParser(&credential)

	user, err := a.UserRepository.Login(credential.Username, credential.Password)

	if err != nil {
		return fiber.NewError(401, "failed login.")
	}

	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = user.Name
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Minute * 10).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.JSON(fiber.Map{"msg": "login successfully.", "statusCode": 200, "token": t})
}
