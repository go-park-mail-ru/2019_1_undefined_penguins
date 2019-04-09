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
	//rows, err := database.Query("SELECT * FROM users WHERE LOWER(nickname)=LOWER($1) OR LOWER(email)=LOWER($2)", user.Nickname, user.Email)
	//if err != nil {
	//	helpers.ResponseCtx(ctx, err.Error(), fasthttp.StatusInternalServerError)
	//	return
	//}
	//defer rows.Close()
	//
	//matchUsers := helpers.RowsToUsers(rows)
	//
	//if len(matchUsers) != 0 {
	//	j, _ := json.Marshal(matchUsers)
	//	helpers.ResponseCtx(ctx, string(j), fasthttp.StatusConflict)
	//	return
	//}

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
SELECT login, name, score, email, hashpassword
FROM users
WHERE email = $1`

func GetUserByEmail(email string) *models.User {
	var user []models.User
	rows, err := Query(selectByEmail, email)
	if err != nil {
		helpers.LogMsg(err)
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
		helpers.LogMsg(err)
		return users
	}
	defer rows.Close()
	fmt.Println(users)

	users = RowsToUsers(rows)
	fmt.Println(users)
	return users
}
