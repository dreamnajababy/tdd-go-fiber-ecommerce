package repositories

import (
	"errors"

	models "github.com/dreamnajababy/go-ecom/src/models"
)

type UserInlineRepository struct {
	Users []models.User
}

func (this *UserInlineRepository) InitUsers() {
	this.Users = []models.User{
		{Username: "dreamnajababy", Password: "1234"},
		{Username: "anijjung", Password: "12345678"},
		{Username: "dhukhung", Password: "87654321"},
		{Username: "anutta", Password: "7231894438"},
	}
}

func (this UserInlineRepository) Login(username, password string) (models.User, error) {
	for _, user := range this.Users {
		if user.Username == username && user.Password == password {
			return user, nil
		}
	}
	return models.User{}, errors.New("login unsuccessfully.")
}
