package database

import (
	"fmt"
	"2019_1_undefined_penguins/internal/pkg/models"
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
	UpdateUserByID(&user, 1)

	rows = sqlmock.NewRows([]string{"games", "name", "score"}).
		AddRow(10, "name", 5).
		RowError(1, fmt.Errorf("error"))
	mock.ExpectQuery("UPDATE").WillReturnError(fmt.Errorf("some error"))
	UpdateUserByID(&user, 1)

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
	rows = sqlmock.NewRows([]string{"id", "login", "email", "hashpassword", "score", "name", "games"}).
		AddRow(1, "login", "login@mail.ru", "hdfbkfbdj", 10, "name", "5").
		RowError(1, fmt.Errorf("error"))
	mock.ExpectQuery("SELECT").WillReturnRows(rows)

	GetUserByLogin("login")
	rows = sqlmock.NewRows([]string{"count(*)"}).
		AddRow(1).
		RowError(1, fmt.Errorf("error"))
	mock.ExpectQuery("SELECT").WillReturnRows(rows)
	UsersCount()
	
	Disconnect()
	initConfig()
}
