package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	db "github.com/go-park-mail-ru/2019_1_undefined_penguins/iternal/pkg/database"
	"github.com/go-park-mail-ru/2019_1_undefined_penguins/iternal/pkg/helpers"
	"github.com/go-park-mail-ru/2019_1_undefined_penguins/iternal/pkg/models"

	"github.com/satori/uuid"
)

//add concret error + body (w.Write())
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
	found := db.GetUserByEmail(user.Email)
	if found == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if !helpers.CheckPasswordHash(user.Password, found.HashPassword) {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	sessionID := uuid.NewV4()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	bytes, err := json.Marshal(found)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	helpers.CreateCookie(&w, sessionID.String())
	models.Sessions[sessionID.String()] = user.Email
	w.Write(bytes)
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	// var user models.User
	var user =  models.User{}  //где User - это таблица
	err = json.Unmarshal(body, &user)
	if err != nil {
		helpers.LogMsg(err)
		return
	}
	found := db.GetUserByEmail(user.Email)
	if found != nil {
		w.WriteHeader(409)
		return
	}

	user.HashPassword, err = helpers.HashPassword(user.Password)

	err = db.CreateUser(&user)
	if err != nil {

		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	sessionID := uuid.NewV4()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Println(sessionID)
	helpers.CreateCookie(&w, sessionID.String())
	//сюда как то подрубить мемкэш, вместо user.Email будет id
	models.Sessions[sessionID.String()] = user.Email
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

//add w.Write() everywhere