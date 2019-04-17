package database

import (
	"2019_1_undefined_penguins/internal/pkg/helpers"
	"2019_1_undefined_penguins/internal/pkg/models"
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
	"testing"
	"time"
)

func GetUserFromJSON(fileName string) (*models.User, error) {
	dir, err := os.Getwd()
	if err != nil {
		helpers.LogMsg("Getting directory error: ", err)
		return nil, err

	}
	dir = strings.Replace(dir, "/internal/pkg/controllers", "", -1)
	file, err := os.Open(dir + "/configs/" + fileName)
	if err != nil {
		helpers.LogMsg("Open directory error: ", err)
		return nil, err
	}

	body, _ := ioutil.ReadAll(file)
	var user models.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		helpers.LogMsg("Init parse user error: ", err)
		return nil, err
	}
	return &user, nil
}

func TestCreate(t *testing.T) {
	err := Connect()
	if err != nil {
		helpers.LogMsg("Connection error: ", err)
		t.Fatal(err)
	}
	defer Disconnect()

	var user models.User
	user.Login = time.Now().Format("20060102150405") + "login"
	user.Email = time.Now().Format("20060102150405") + "@mail.ru"
	user.Password = time.Now().Format("20060102150405") + "password"
	err = CreateUser(&user)

	if err != nil {
		t.Fatal(err)
	}

	found, err := GetUserByEmail(user.Email)
	if err != nil {
		t.Fatal(err)
	}
	if found.Login != user.Login {
		t.Fatal(found.Login + "!=" + user.Login)
	}
}

func TestGetUserByID(t *testing.T) {
	err := Connect()
	if err != nil {
		helpers.LogMsg("Connection error: ", err)
		t.Fatal(err)
	}
	defer Disconnect()

	_, err = GetUserByID(1)
	if err != nil {
		t.Fatal(err)
	}

}

func TestUpdateImage(t *testing.T) {
	err := Connect()
	if err != nil {
		helpers.LogMsg("Connection error: ", err)
		t.Fatal(err)
	}
	defer Disconnect()

	var user models.User
	user.Login = time.Now().Format("20060102150405") + "log"
	user.Email = time.Now().Format("20060102150405") + "@corp.mail.ru"
	user.Password = time.Now().Format("20060102150405") + "password"
	err = CreateUser(&user)

	if err != nil {
		t.Fatal(err)
	}

	err = UpdateImage(user.Login, time.Now().Format("20060102150405")+"pic")
	if err != nil {
		t.Fatal(err)
	}
}

func TestLeaders(t *testing.T) {
	err := Connect()
	if err != nil {
		helpers.LogMsg("Connection error: ", err)
		t.Fatal(err)
	}
	defer Disconnect()
	users, err := GetLeaders(1)
	if err != nil {
		t.Fatal(err)
	}
	if len(users) == 0 {
		t.Fatal(err)
	}
}

func TestAddGame(t *testing.T) {
	err := Connect()
	if err != nil {
		helpers.LogMsg("Connection error: ", err)
		t.Fatal(err)
	}
	defer Disconnect()

	var user models.User
	user.Login = time.Now().Format("20060102150405") + "l"
	user.Email = time.Now().Format("20060102150405") + "@my.com"
	user.Password = time.Now().Format("20060102150405") + "password"
	err = CreateUser(&user)

	if err != nil {
		t.Fatal(err)
	}

	err = AddGame(user.Email)
	if err != nil {
		t.Fatal(err)
	}
	found, err := GetUserByEmail(user.Email)
	if err != nil {
		t.Fatal(err)
	}
	if found.Login != user.Login {
		t.Fatal(found.Login + "!=" + user.Login)
	}
	if found.Games != 1 {
		t.Fatal("Неверное количество игр")
	}

}

func TestNewRecord(t *testing.T) {
	err := Connect()
	if err != nil {
		helpers.LogMsg("Connection error: ", err)
		t.Fatal(err)
	}
	defer Disconnect()

	var user models.User
	user.Login = time.Now().Format("20060102150405") + "lo"
	user.Email = time.Now().Format("20060102150405") + "@bk.ru"
	user.Password = time.Now().Format("20060102150405") + "password"
	err = CreateUser(&user)

	if err != nil {
		t.Fatal(err)
	}
	record := 200
	err = NewRecord(user.Email, record)
	if err != nil {
		t.Fatal(err)
	}

	found, err := GetUserByEmail(user.Email)
	if err != nil {
		t.Fatal(err)
	}
	if found.Login != user.Login {
		t.Fatal(found.Login + "!=" + user.Login)
	}
	if found.Games != 1 {
		t.Fatal("Неверное количество игр")
	}
	if err != nil {
		t.Fatal(err)
	}
	if record != int(found.Score) {
		t.Fatal("Неверно записан рекорд")
	}
}
