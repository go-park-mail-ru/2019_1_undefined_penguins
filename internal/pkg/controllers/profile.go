package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	//"2019_1_undefined_penguins/internal/pkg/helpers"
	"2019_1_undefined_penguins/internal/pkg/models"

	db "2019_1_undefined_penguins/internal/pkg/database"
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

	err = db.UpdateUser(&user, email)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	models.Sessions[cookie.Value] = user.Email
	w.Write(body)
}
