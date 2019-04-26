package database

import (
	"2019_1_undefined_penguins/internal/pkg/helpers"
	"2019_1_undefined_penguins/internal/pkg/models"

	"github.com/jackc/pgx"
)

const insertUser = `
INSERT INTO users (email, login, hashpassword)
VALUES ($1, $2, $3)
RETURNING id, login, score`

func CreateUser(newUser *models.User) error {

	if connection == nil {
		return pgx.ErrDeadConn
	}

	if err := connection.QueryRow(insertUser, newUser.Email, newUser.Login, newUser.HashPassword).Scan(&newUser.ID, &newUser.Login, &newUser.Score); err != nil {
		helpers.LogMsg(err)
		return err
	}

	return nil
}

const updateUserByEmail = `
UPDATE users AS u
SET lastVisit = now(),
	login = $2,
	email = $3
FROM pictures AS p
WHERE u.email = $1 
AND u.picture = p.id
RETURNING games, name, score`

func UpdateUser(user models.User, oldEmail string) (models.User, error) {
	user.Password = ""
	err := connection.QueryRow(updateUserByEmail, oldEmail, user.Login, user.Email).Scan(&user.Score, &user.Picture, &user.Games)
	if err != nil {
		helpers.LogMsg(err)
		return user, err
	}
	user.Picture = "http://localhost:8081/data/" + user.Picture
	return user, nil
}

const updateUserByID = `
UPDATE users AS u
SET lastVisit = now(),
	login = $2,
	email = $3
FROM pictures AS p
WHERE u.id = $1 
AND u.picture = p.id
RETURNING games, name, score`

func UpdateUserByID(user models.User, id uint) (models.User, error) {
	user.Password = ""
	err := connection.QueryRow(updateUserByID, id, user.Login, user.Email).Scan(&user.Score, &user.Picture, &user.Games)
	if err != nil {
		helpers.LogMsg(err)
		return user, err
	}
	user.Picture = "http://localhost:8081/data/" + user.Picture
	return user, nil
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

func GetUserByID(id uint) (*models.User, error) {
	user := models.User{}
	row := connection.QueryRow(selectByID, id)

	err = row.Scan(&user.ID, &user.Login, &user.Email, &user.HashPassword, &user.Score, &user.Picture, &user.Games)

	if err != nil {
		helpers.LogMsg(err)
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
	_, err := Exec(newPersonalRecord, email, record)
	if err != nil {
		helpers.LogMsg(err)
		return err
	}
	return nil
}
