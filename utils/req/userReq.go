package req

import (
	"kasir/model"
	"kasir/model/web"
)

func PassBody(users web.UserRequest) *model.User {
	return &model.User{
		Name: users.Name,
		Email: users.Email,
		Password: users.Password,
	}
}
