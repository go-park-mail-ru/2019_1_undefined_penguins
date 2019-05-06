package helpers

import (
	"2019_1_undefined_penguins/internal/pkg/models"
)

//TODO maybe one structure?
func ModelToProto(user *models.User) *models.UserProto {
	return &models.UserProto{
		ID: uint64(user.ID),
		Login: user.Login,
		Email: user.Email,
		Password: user.Password,
		HashPassword: user.HashPassword,
		Score: uint64(user.Score),
		Picture: user.Picture,
		Games: uint64(user.Games),
	}
}

func ProtoToModel(user *models.UserProto) *models.User {
	return &models.User{
		ID: uint(user.ID),
		Login: user.Login,
		Email: user.Email,
		Password: user.Password,
		HashPassword: user.HashPassword,
		Score: uint(user.Score),
		Picture: user.Picture,
		Games: uint(user.Games),
	}
}
