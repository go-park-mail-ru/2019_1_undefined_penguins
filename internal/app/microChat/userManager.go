package microChat

import (
	"context"
	"2019_1_undefined_penguins/internal/pkg/database"
)

type UserManager struct {
	Login string
}

func NewUserManager() *UserManager {
	return &UserManager{
		Login: "",
	}
}

func (u *UserManager) Check(ctx context.Context, in *User) (*User, error) {
	user, _ := database.GetUserByLogin(in.Login)
	in.Login = user.Login
	return in, nil
}
