package helpers

import (
	"fmt"

	"github.com/go-park-mail-ru/2019_1_undefined_penguins/iternal/pkg/database"
	"github.com/go-park-mail-ru/2019_1_undefined_penguins/iternal/pkg/models"
)

const insertUser = `
INSERT INTO users (email, hashpassword) 
VALUES ($1, $2)
RETURNING login, name, score`

func CreateUser(newUser *models.User) error {
	transaction := database.StartTransaction()
	defer transaction.Rollback()

	if _, err := transaction.Exec(insertUser, newUser.Email, newUser.HashPassword); err != nil {

		transaction.Rollback()
		return err
	}

	if err := database.CommitTransaction(transaction); err != nil {
		transaction.Rollback()
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
	transaction := database.StartTransaction()
	defer transaction.Rollback()
	user.Password = "" //Лови коммент
	if _, err := transaction.Exec(updateUserByEmail, oldEmail, user.Login, user.Email, user.Name); err != nil {
		fmt.Println(err)
		transaction.Rollback()
		return err
	}

	if err := database.CommitTransaction(transaction); err != nil {
		transaction.Rollback()
		return err
	}

	return nil
}

const selectByEmail = `
SELECT login, name, score, email, hashpassword
FROM users 
WHERE email = $1`

func GetUserByEmail(email string) *models.User {

	transaction := database.StartTransaction()

	defer transaction.Rollback()

	var user models.User
	if err := transaction.QueryRow(selectByEmail, email).Scan(&user.Login, &user.Name, &user.Score, &user.Email, &user.HashPassword); err != nil {
		return nil

	}

	if err := database.CommitTransaction(transaction); err != nil {
		transaction.Rollback()
		return nil
	}
	user.Password = ""
	return &user

}

const GetLeadersPage = `
SELECT login, email, score
FROM users
ORDER BY score DESC
LIMIT 3 OFFSET $1`

func GetLeaders(id int) []*models.User {
	var users []*models.User
	transaction := database.StartTransaction()
	if elements, err := transaction.Query(GetLeadersPage, (id-1)*3); err != nil {
		fmt.Println(err)
		return users
	} else {
		for elements.Next() {
			var user models.User
			if err := elements.Scan(
				&user.Login,
				&user.Email,
				&user.Score); err != nil {
				return users
			}
			user.Password = ""
			users = append(users, &user)
		}
		//close\
		elements.Close()
	}
	return users
}
