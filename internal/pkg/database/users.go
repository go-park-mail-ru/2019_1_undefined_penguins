package database

import (
	"2019_1_undefined_penguins/internal/pkg/helpers"
	"2019_1_undefined_penguins/internal/pkg/models"
	"fmt"
	"github.com/jackc/pgx"
)

const insertUser = `
INSERT INTO users (email, login, hashpassword)
VALUES ($1, $2, $3)
RETURNING id, login, score`

func CreateUser(newUser *models.User) error {
	fmt.Println("user", newUser)
	fmt.Println("login", newUser.Login)
	fmt.Println("pass", newUser.Password)

	if connection == nil {
		return pgx.ErrDeadConn
	}

	//var row pgtype.Record
	//err := conn.QueryRow("insert ... returning (...)").Scan(&row)
	if err := connection.QueryRow(insertUser, newUser.Email, newUser.Login, newUser.HashPassword).Scan(&newUser.ID, &newUser.Login, &newUser.Score); err != nil {
		helpers.LogMsg(err)
		return err
	}

	return nil
}

const updateUserByEmail = `
UPDATE users
SET lastVisit = now(),
	login = $2,
	email = $3
WHERE email = $1`

func UpdateUser(user *models.User, oldEmail string) error {
	user.Password = "" //Лови коммент
	_, err := Exec(updateUserByEmail, oldEmail, user.Login, user.Email)
	if err != nil {
		helpers.LogMsg(err)
		return err
	}

	return nil
}

const updateUserByID = `
UPDATE users
SET lastVisit = now(),
	login = $2,
	email = $3
WHERE id = $1`

func UpdateUserByID(user *models.User, id uint) error {
	user.Password = "" //Лови коммент
	_, err := Exec(updateUserByID, id, user.Login, user.Email)
	if err != nil {
		helpers.LogMsg(err)
		return err
	}

	return nil
}


const updateImageByLogin = `
SELECT insertPicture($1, $2);
`

func UpdateImage(login string, avatar string) error {
	_, err := Exec(updateImageByLogin, login, avatar)
	if err != nil {
		helpers.LogMsg(err)
		return err
	}
	return nil
}

const selectByEmail = `
SELECT users.id, login, email, hashpassword, score, name, games
FROM users, pictures
WHERE users.email = $1
AND users.picture = pictures.id`

func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := connection.QueryRow(selectByEmail, email).Scan(&user.ID, &user.Login, &user.Email, &user.HashPassword, &user.Score, &user.Picture, &user.Games)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	user.Picture = "http://localhost:8081/data/" + user.Picture
	return &user, nil
}

const selectByID = `
SELECT users.id, login, email, hashpassword, score, name, games
FROM users, pictures
WHERE users.id = $1
AND users.picture = pictures.id`

func GetUserByID(id string) (*models.User, error) {
	var user models.User
	err := connection.QueryRow(selectByID, id).Scan(&user.ID, &user.Login, &user.Email, &user.HashPassword, &user.Score, &user.Picture, &user.Games)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	user.Picture = "http://localhost:8081/data/" + user.Picture
	return &user, nil
}

const getLeadersPage = `
SELECT login, score, email
FROM users
ORDER BY score DESC
LIMIT 3 OFFSET $1`

func GetLeaders(id int) ([]models.User, error) {
	var users []models.User
	rows, err := Query(getLeadersPage, (id-1)*3)
	if err != nil {
		helpers.LogMsg(err)
		return users, err
	}
	defer rows.Close()

	users = RowsToUsers(rows)
	fmt.Println(users)
	return users, nil
}

const iterateGame = `
UPDATE users
SET games = games + 1
WHERE email = $1`

const newPersonalRecord = `
UPDATE users
SET games = games + 1,
	score = $2
WHERE email = $1`

func AddGame(email string) error {
	_, err := Exec(iterateGame, email)
	if err != nil {
		helpers.LogMsg(err)
		return err
	}
	return nil
}

func NewRecord(email string, record int) error {
	_, err := Exec(iterateGame, email, record)
	if err != nil {
		helpers.LogMsg(err)
		return err
	}
	return nil
}
