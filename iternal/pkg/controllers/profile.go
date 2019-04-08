package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	//"github.com/go-park-mail-ru/2019_1_undefined_penguins/iternal/pkg/helpers"
	"github.com/go-park-mail-ru/2019_1_undefined_penguins/iternal/pkg/models"
	db "github.com/go-park-mail-ru/2019_1_undefined_penguins/iternal/pkg/database"
)

func Me(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		return
	}
	cookie, err := r.Cookie("sessionid")
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	email, found := models.Sessions[cookie.Value]
	if !found {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	user := db.GetUserByEmail(email)

	if user == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	bytes, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(bytes)
}

func ChangeProfile(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		return
	}
	cookie, err := r.Cookie("sessionid")
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	email, found := models.Sessions[cookie.Value]
	if !found {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	fmt.Println(email)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	var user models.User
	err = json.Unmarshal(body, &user)
	fmt.Println(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = db.UpdateUser(&user, email)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	models.Sessions[cookie.Value] = user.Email
	w.Write(body)
}
