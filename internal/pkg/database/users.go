package database

import (
	"2019_1_undefined_penguins/internal/pkg/helpers"
	"2019_1_undefined_penguins/internal/pkg/models"
	"fmt"
)


const selectUser = `
INSERT INTO users (email, hashpassword)
VALUES ($1, $2)
RETURNING login, name, score`

const insertUser = `
INSERT INTO users (email, hashpassword)
VALUES ($1, $2)
RETURNING login, name, score`

func CreateUser(newUser *models.User) error {
	if _, err := Exec(insertUser, newUser.Email, newUser.HashPassword); err != nil {
		helpers.LogMsg(err)
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
		helpers.LogMsg(err)
		return err
	}

	return nil
}

const selectByEmail = `
SELECT login, name, email, hashpassword
FROM users
WHERE email = $1`

func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	rows, err := Query(selectByEmail, email)
	if err != nil {
		helpers.LogMsg(err)
		return nil, err
	}
	defer rows.Close()

	err = connection.QueryRow(selectByEmail, email).Scan(&user.Login, &user.Name, &user.Email, &user.HashPassword)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

const GetLeadersPage = `
SELECT login, name, score, email
FROM users
ORDER BY score DESC
LIMIT 3 OFFSET $1`

func GetLeaders(id int) ([]models.User, error) {
	var users []models.User
	rows, err := Query(GetLeadersPage, (id-1)*3)
	if err != nil {
		helpers.LogMsg(err)
		return users, err
	}
	defer rows.Close()
	fmt.Println(users)

	users = RowsToUsers(rows)
	fmt.Println(users)
	return users, nil
}
