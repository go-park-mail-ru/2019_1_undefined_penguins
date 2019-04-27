package microChat

import (
	"context"
	"2019_1_undefined_penguins/internal/pkg/database"
)

type UserManager struct {
	Login string
	ID uint64
}

func NewUserManager() *UserManager {
	return &UserManager{
		Login: "",
		ID: 0,
	}
}

func (u *UserManager) Check(ctx context.Context, in *User) (*User, error) {
	user, _ := database.GetUserByLogin(in.Login)
	in.Login = user.Login
	in.ID = uint64(user.ID)
	return in, nil
}
