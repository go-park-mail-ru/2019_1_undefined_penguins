package controllers

import (
	"encoding/json"
	"time"

	//"fmt"
	"io/ioutil"
	"net/http"

	db "2019_1_undefined_penguins/internal/pkg/database"
	"2019_1_undefined_penguins/internal/pkg/helpers"
	"2019_1_undefined_penguins/internal/pkg/models"
	jwt "github.com/dgrijalva/jwt-go"
	//"github.com/satori/uuid"
)


var SECRET = []byte("myawesomesecret")

func SignIn(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		return
	}
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
	//sessionID := uuid.NewV4()
	ttl := 15 * time.Second

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Email,
		"exp": time.Now().UTC().Add(ttl).Unix(),
	})
	//jwt.StandardClaims

	str, err := token.SignedString(SECRET)
	if err != nil {
		w.Write([]byte("=(" + err.Error()))
		return
	}

	//helpers.CreateCookie(&w, str)
	cookie := &http.Cookie{
		Name:  "sessionid",
		Value: str,
		Expires:  time.Now().Add(60 * time.Hour),
		HttpOnly: true,
	}

	http.SetCookie(w, cookie)
	w.Write([]byte(str))

	//bytes, err := json.Marshal(found)
	//
	//if err != nil {
	//	w.WriteHeader(http.StatusInternalServerError)
	//	return
	//}
	//
	//helpers.CreateCookie(&w, sessionID.String())
	//
	//models.Sessions[sessionID.String()] = user.Email
	//w.Write(bytes) // calls WriteHeader(http.StatusOK) by default
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		helpers.LogMsg(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	var user models.User
	//var user = models.User{} //где User - это таблица
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
		w.WriteHeader(http.StatusInternalServerError)
		return
	}


	//sessionID := uuid.NewV4()
	ttl := 15 * time.Second

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Email,
		"exp": time.Now().UTC().Add(ttl).Unix(),
	})

	str, err := token.SignedString(SECRET)
	if err != nil {
		w.Write([]byte("=(" + err.Error()))
		return
	}

	cookie := &http.Cookie{
		Name:  "sessionid",
		Value: str,
	}

	http.SetCookie(w, cookie)
	w.Write([]byte(str))


	//sessionID := uuid.NewV4()
	//
	//helpers.CreateCookie(&w, sessionID.String())
	//
	////сюда как то подрубить мемкэш, вместо user.Email будет id
	//models.Sessions[sessionID.String()] = user.Email

}

func SignOut(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		return
	}
	cookie, err := r.Cookie("sessionid")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("You are not authorized"))
		return
	}

	helpers.DeleteCookie(&w, cookie)

}

