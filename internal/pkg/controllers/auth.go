package controllers

import (
	"encoding/json"
	"time"

	"io/ioutil"
	"net/http"

	db "2019_1_undefined_penguins/internal/pkg/database"
	"2019_1_undefined_penguins/internal/pkg/helpers"
	"2019_1_undefined_penguins/internal/pkg/models"

	"github.com/dgrijalva/jwt-go"
)

var SECRET = []byte("myawesomesecret")

func SignIn(w http.ResponseWriter, r *http.Request) {

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

	found, _ := db.GetUserByEmail(user.Email)

	if found == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if !helpers.CheckPasswordHash(user.Password, found.HashPassword) {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	ttl := time.Hour

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":    found.ID,
		"userEmail": user.Email,
		"exp":       time.Now().UTC().Add(ttl).Unix(),
	})

	str, err := token.SignedString(SECRET)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("=(" + err.Error()))
		return
	}

	cookie := &http.Cookie{
		Name:     "sessionid",
		Value:    str,
		Expires:  time.Now().Add(time.Hour),
		HttpOnly: true,
	}

	bytes, err := json.Marshal(found)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, cookie)
	w.Write(bytes)
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		helpers.LogMsg(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	var user models.User
	err = json.Unmarshal(body, &user)

	if err != nil {
		helpers.LogMsg(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	found, _ := db.GetUserByEmail(user.Email)
	if found != nil {
		w.WriteHeader(http.StatusConflict)
		return
	}

	user.HashPassword, err = helpers.HashPassword(user.Password)

	err = db.CreateUser(&user)
	if err != nil {
		helpers.LogMsg(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	ttl := time.Hour

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":    user.ID,
		"userEmail": user.Email,
		"exp":       time.Now().UTC().Add(ttl).Unix(),
	})

	str, err := token.SignedString(SECRET)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("=(" + err.Error()))
		return
	}

	cookie := &http.Cookie{
		Name:     "sessionid",
		Value:    str,
		Expires:  time.Now().Add(time.Hour),
		HttpOnly: true,
	}
	user.Password = ""
	user.Picture = db.ImagesAddress + "Default.png"
	bytes, err := json.Marshal(user)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, cookie)
	w.Write(bytes)
}

func SignOut(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("sessionid")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("You are not authorized"))
		return
	}

	helpers.LogMsg("User " + cookie.Value + " has logged out")
	helpers.DeleteCookie(&w, cookie)
	w.WriteHeader(http.StatusOK)
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello penguins"))
}
