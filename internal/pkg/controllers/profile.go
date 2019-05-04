package controllers

import (
	"2019_1_undefined_penguins/internal/pkg/helpers"
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"time"

	"io/ioutil"
	"net/http"

	"2019_1_undefined_penguins/internal/pkg/database"
	db "2019_1_undefined_penguins/internal/pkg/database"
	"2019_1_undefined_penguins/internal/pkg/models"

	"github.com/dgrijalva/jwt-go"
)

func Me(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("sessionid")

	token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
		return SECRET, nil
	})

	claims, _ := token.Claims.(jwt.MapClaims)

	temp := claims["userID"]
	mytemp := uint(temp.(float64))

	user, err := db.GetUserByID(mytemp)
	if user == nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	bytes, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(bytes)
	return
}

func ChangeProfile(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("sessionid")

	token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
		return SECRET, nil
	})

	claims, _ := token.Claims.(jwt.MapClaims)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	var user models.User
	err = json.Unmarshal(body, &user)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	temp := claims["userID"]
	mytemp := uint(temp.(float64))
	user, err = db.UpdateUserByID(user, mytemp)
	if err != nil {
		helpers.LogMsg(err)
		w.WriteHeader(http.StatusConflict)
		return
	}

	bytes, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(bytes)
	return
}

func UploadImage(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("sessionid")

	token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
		return SECRET, nil
	})

	claims, _ := token.Claims.(jwt.MapClaims)
	user, err := db.GetUserByEmail(claims["userEmail"].(string))
	if user == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = r.ParseMultipartForm(5 * 1024 * 1025)
	if err != nil {
		helpers.LogMsg("Ошибка при парсинге тела запроса")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	file, handler, err := r.FormFile("avatar")
	if err != nil {
		helpers.LogMsg("Ошибка при получении файла из тела запроса")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer file.Close()
	extension := filepath.Ext(handler.Filename)
	if extension == "" {
		helpers.LogMsg("Файл не имеет расширения")

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	t := time.Now()

	fileName := user.Login + t.Format("20060102150405") + extension
	fileAndPath := "static/" + fileName
	saveFile, err := os.Create(fileAndPath)
	if err != nil {
		helpers.LogMsg("Create", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer saveFile.Close()

	_, err = io.Copy(saveFile, file)
	if err != nil {
		helpers.LogMsg("Copy", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = database.UpdateImage(user.Login, fileName)
	if err != nil {
		helpers.LogMsg("Ошибка при обновлении картинки в базе данных")

		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	return
}
