package repositories

import models "github.com/dreamnajababy/go-ecom/src/models"

type UserInlineRepository struct {
	Users []models.User
}

func (this UserInlineRepository) Login(username, password string) (models.User, error) {
	return models.User{"dreamnajababy"}, nil
}
