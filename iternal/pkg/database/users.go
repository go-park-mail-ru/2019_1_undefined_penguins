package database

import (
	"2019_1_undefined_penguins/iternal/pkg/helpers"
	"2019_1_undefined_penguins/iternal/pkg/models"
	"fmt"
)

const insertUser = `
INSERT INTO users (email, hashpassword)
VALUES ($1, $2)
RETURNING login, name, score`

func CreateUser(newUser *models.User) error {
	if _, err := Exec(insertUser, newUser.Email, newUser.HashPassword); err != nil {
		helpers.LogMsg(err.Error())
		return err
	}

	return nil
}

const updateUserByEmail = `
UPDATE users
SET lastVisit = now(),
	login = $2,
	email = $3,
	name = $4
WHERE email = $1`

func UpdateUser(user *models.User, oldEmail string) error {
	user.Password = "" //Лови коммент
	if _, err := Exec(updateUserByEmail, oldEmail, user.Login, user.Email, user.Name); err != nil {
		helpers.LogMsg(err.Error())
		return err
	}

	return nil
}

const selectByEmail = `
SELECT login, name, score, email, hashpassword
FROM users
WHERE email = $1`

func GetUserByEmail(email string) *models.User {
	var user []models.User
	rows, err := Query(selectByEmail, email)
	if err != nil {
		helpers.LogMsg(err.Error())
		return nil
	}
	defer rows.Close()

	user = RowsToUsers(rows)
	if len(user) != 0 {
		user[0].Password = ""
		return &user[0]
	}
	return nil
}

const GetLeadersPage = `
SELECT login, email, score
FROM users
ORDER BY score DESC
LIMIT 3 OFFSET $1`

func GetLeaders(id int) []models.User {
	var users []models.User
	rows, err := Query(GetLeadersPage, (id-1)*3)
	if err != nil {
		helpers.LogMsg(err.Error())
		return users
	}
	defer rows.Close()
	fmt.Println(users)

	users = RowsToUsers(rows)
	fmt.Println(users)
	return users
}
