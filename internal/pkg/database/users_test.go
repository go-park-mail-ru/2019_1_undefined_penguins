package database

import (
	"2019_1_undefined_penguins/internal/pkg/models"
	"fmt"
	"testing"
	"time"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
)

func TestGetLeaders(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"login", "score", "email"}).
		AddRow("login", "one", "login@mail.ru").
		RowError(1, fmt.Errorf("error"))
	mock.ExpectBegin()
	mock.ExpectQuery("SELECT").WillReturnRows(rows)
	mock.ExpectCommit()
	SetMock(db)
	GetLeaders(1)
	mock.ExpectBegin()
	mock.ExpectQuery("SELECT").WillReturnRows(rows).WillReturnError(fmt.Errorf("some error"))
	mock.ExpectRollback()
	GetLeaders(-1)
	rows = sqlmock.NewRows([]string{"login", "score"}).
		AddRow("login", "one").
		RowError(1, fmt.Errorf("error"))
	mock.ExpectBegin()
	mock.ExpectQuery("SELECT").WillReturnRows(rows)
	mock.ExpectRollback()
	GetLeaders(1)
	connection = nil
	GetLeaders(1)

}

func TestCreateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "login", "score"}).
		AddRow(1, "login", 10).
		RowError(1, fmt.Errorf("error"))
	mock.ExpectQuery("INSERT").WillReturnRows(rows)
	var user models.User
	user.Login = time.Now().Format("20060102150405") + "login"
	user.Email = time.Now().Format("20060102150405") + "@mail.ru"
	user.HashPassword = time.Now().Format("20060102150405") + "password"
	connection = nil
	CreateUser(&user)
	connection = db

	CreateUser(&user)
	mock.ExpectQuery("INSERT").WillReturnRows(rows).WillReturnError(fmt.Errorf("some error"))
	CreateUser(&user)
	rows = sqlmock.NewRows([]string{"id", "login", "email", "hashpassword", "score", "name", "games"}).
		AddRow(1, "login", "login@mail.ru", "hdfbkfbdj", 10, "name", "5").
		RowError(1, fmt.Errorf("error"))
	mock.ExpectQuery("SELECT").WillReturnRows(rows)
	GetUserByID(1)
	rows = sqlmock.NewRows([]string{"id", "login", "email", "hashpassword", "score", "name", "games"}).
		AddRow(1, "login", "login@mail.ru", "hdfbkfbdj", 10, "name", "5").
		RowError(1, fmt.Errorf("error"))
	mock.ExpectQuery("SELECT").WillReturnRows(rows)
	GetUserByEmail("login@mail.ru")
	mock.ExpectQuery("SELECT").WillReturnRows(rows).WillReturnError(fmt.Errorf("some error"))
	GetUserByID(100005000)
	GetUserByEmail("notalogin@mail.ru")
	mock.ExpectBegin()
	mock.ExpectExec("UPDATE").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	err = AddGame("login@mail.ru")
	mock.ExpectBegin()
	mock.ExpectExec("UPDATE").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	err = NewRecord("login@mail.ru", 100500)

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE").WillReturnError(fmt.Errorf("some error"))
	mock.ExpectRollback()
	err = AddGame("login@mail.ru")

	mock.ExpectBegin()
	mock.ExpectExec("SELECT").WillReturnError(fmt.Errorf("some error"))
	mock.ExpectRollback()
	err = UpdateImage("login", "name")

	mock.ExpectBegin()
	mock.ExpectExec("SELECT").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	err = UpdateImage("login", "name")

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE").WillReturnError(fmt.Errorf("some error"))
	mock.ExpectRollback()
	err = NewRecord("login@mail.ru", 100500)

	rows = sqlmock.NewRows([]string{"games", "name", "score"}).
		AddRow(10, "name", 5).
		RowError(1, fmt.Errorf("error"))
	mock.ExpectQuery("UPDATE").WillReturnRows(rows)
	UpdateUserByID(user, 1)

	rows = sqlmock.NewRows([]string{"games", "name", "score"}).
		AddRow(10, "name", 5).
		RowError(1, fmt.Errorf("error"))
	mock.ExpectQuery("UPDATE").WillReturnError(fmt.Errorf("some error"))
	UpdateUserByID(user, 1)

	rows = sqlmock.NewRows([]string{"games", "name", "score"}).
		AddRow(10, "name", 5).
		RowError(1, fmt.Errorf("error"))
	mock.ExpectQuery("UPDATE").WillReturnRows(rows)
	UpdateUser(user, "login@mail.ru")

	rows = sqlmock.NewRows([]string{"games", "name", "score"}).
		AddRow(10, "name", 5).
		RowError(1, fmt.Errorf("error"))
	mock.ExpectQuery("UPDATE").WillReturnError(fmt.Errorf("some error"))
	UpdateUser(user, "login@mail.ru")
	Disconnect()
}

// func GetUserFromJSON(fileName string) (*models.User, error) {
// 	dir, err := os.Getwd()
// 	if err != nil {
// 		helpers.LogMsg("Getting directory error: ", err)
// 		return nil, err

// 	}
// 	dir = strings.Replace(dir, "/internal/pkg/controllers", "", -1)
// 	file, err := os.Open(dir + "/configs/" + fileName)
// 	if err != nil {
// 		helpers.LogMsg("Open directory error: ", err)
// 		return nil, err
// 	}

// 	body, _ := ioutil.ReadAll(file)
// 	var user models.User
// 	err = json.Unmarshal(body, &user)
// 	if err != nil {
// 		helpers.LogMsg("Init parse user error: ", err)
// 		return nil, err
// 	}
// 	return &user, nil
// }

// func TestCreate(t *testing.T) {
// 	err := Connect()
// 	if err != nil {
// 		helpers.LogMsg("Connection error: ", err)
// 		t.Error(err)
// 	}
// 	defer Disconnect()

// 	var user models.User
// 	user.Login = time.Now().Format("20060102150405") + "login"
// 	user.Email = time.Now().Format("20060102150405") + "@mail.ru"
// 	user.Password = time.Now().Format("20060102150405") + "password"
// 	err = CreateUser(&user)

// 	if err != nil {
// 		t.Error(err)
// 	}

// 	found, err := GetUserByEmail(user.Email)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	if found.Login != user.Login {
// 		t.Error(found.Login + "!=" + user.Login)
// 	}
// }

// func TestGetUserByID(t *testing.T) {
// 	err := Connect()
// 	if err != nil {
// 		helpers.LogMsg("Connection error: ", err)
// 		t.Error(err)
// 	}
// 	defer Disconnect()

// 	_, err = GetUserByID(1)
// 	if err != nil {
// 		t.Error(err)
// 	}

// }

// func TestUpdateImage(t *testing.T) {
// 	err := Connect()
// 	if err != nil {
// 		helpers.LogMsg("Connection error: ", err)
// 		t.Error(err)
// 	}
// 	defer Disconnect()

// 	var user models.User
// 	user.Login = time.Now().Format("20060102150405") + "log"
// 	user.Email = time.Now().Format("20060102150405") + "@corp.mail.ru"
// 	user.Password = time.Now().Format("20060102150405") + "password"
// 	err = CreateUser(&user)

// 	if err != nil {
// 		t.Error(err)
// 	}

// 	err = UpdateImage(user.Login, time.Now().Format("20060102150405")+"pic")
// 	if err != nil {
// 		t.Error(err)
// 	}
// }

// func TestLeaders(t *testing.T) {
// 	err := Connect()
// 	if err != nil {
// 		helpers.LogMsg("Connection error: ", err)
// 		t.Error(err)
// 	}
// 	defer Disconnect()
// 	users, err := GetLeaders(1)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	if len(users) == 0 {
// 		t.Error(err)
// 	}
// }

// func TestAddGame(t *testing.T) {
// 	err := Connect()
// 	if err != nil {
// 		helpers.LogMsg("Connection error: ", err)
// 		t.Error(err)
// 	}
// 	defer Disconnect()

// 	var user models.User
// 	user.Login = time.Now().Format("20060102150405") + "l"
// 	user.Email = time.Now().Format("20060102150405") + "@my.com"
// 	user.Password = time.Now().Format("20060102150405") + "password"
// 	err = CreateUser(&user)

// 	if err != nil {
// 		t.Error(err)
// 	}

// 	err = AddGame(user.Email)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	found, err := GetUserByEmail(user.Email)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	if found.Login != user.Login {
// 		t.Error(found.Login + "!=" + user.Login)
// 	}
// 	if found.Games != 1 {
// 		t.Error("Неверное количество игр")
// 	}

// }

// func TestNewRecord(t *testing.T) {
// 	err := Connect()
// 	if err != nil {
// 		helpers.LogMsg("Connection error: ", err)
// 		t.Error(err)
// 	}
// 	defer Disconnect()

// 	var user models.User
// 	user.Login = time.Now().Format("20060102150405") + "lo"
// 	user.Email = time.Now().Format("20060102150405") + "@bk.ru"
// 	user.Password = time.Now().Format("20060102150405") + "password"
// 	err = CreateUser(&user)

// 	if err != nil {
// 		t.Error(err)
// 	}
// 	record := 200
// 	err = NewRecord(user.Email, record)
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	found, err := GetUserByEmail(user.Email)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	if found.Login != user.Login {
// 		t.Error(found.Login + "!=" + user.Login)
// 	}
// 	if found.Games != 1 {
// 		t.Error("Неверное количество игр")
// 	}
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	if record != int(found.Score) {
// 		t.Error("Неверно записан рекорд")
// 	}
// }

// func TestUpdateByID(t *testing.T) {
// 	err := Connect()
// 	if err != nil {
// 		helpers.LogMsg("Connection error: ", err)
// 		t.Error(err)
// 	}
// 	defer Disconnect()

// 	var user models.User
// 	user.Login = time.Now().Format("20060102150405") + "loginn"
// 	user.Email = time.Now().Format("20060102150405") + "@pochta.ru"
// 	user.Password = time.Now().Format("20060102150405") + "password"
// 	err = CreateUser(&user)
// 	found, err := GetUserByEmail(user.Email)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	user.Login = time.Now().Format("20060102150405") + user.Login
// 	user.Email = time.Now().Format("20060102150405") + user.Email
// 	user.Password = time.Now().Format("20060102150405") + user.Password
// 	UpdateUserByID(user, found.ID)
// 	found, err = GetUserByEmail(user.Email)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	if found.Login != user.Login {
// 		t.Error(err)
// 	}

// }

// func TestUpdateByEmail(t *testing.T) {
// 	err := Connect()
// 	if err != nil {
// 		helpers.LogMsg("Connection error: ", err)
// 		t.Error(err)
// 	}
// 	defer Disconnect()

// 	var user models.User
// 	user.Login = time.Now().Format("20060102150405") + "loginn"
// 	user.Email = time.Now().Format("20060102150405") + "@pochta.ru"
// 	user.Password = time.Now().Format("20060102150405") + "password"
// 	err = CreateUser(&user)
// 	found, err := GetUserByEmail(user.Email)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	user.Login = time.Now().Format("20060102150405") + user.Login
// 	user.Email = time.Now().Format("20060102150405") + user.Email
// 	user.Password = time.Now().Format("20060102150405") + user.Password
// 	UpdateUser(user, found.Email)
// 	found, err = GetUserByEmail(user.Email)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	if found.Login != user.Login {
// 		t.Error(err)
// 	}

// }
