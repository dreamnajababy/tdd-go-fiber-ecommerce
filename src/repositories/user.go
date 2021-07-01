package repositories

import models "github.com/dreamnajababy/go-ecom/src/models"

type UserRepository interface {
	Login(username, password string) (models.User, error)
}
