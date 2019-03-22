package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Me(w http.ResponseWriter, r *http.Request) {
	logMethodAndURL(r)
	if r.Method == "OPTIONS" {
		SetupCORS(&w, r)
		w.WriteHeader(200)
		return
	}
	cookie, err := r.Cookie("sessionid")
	if err != nil {
		fmt.Println("Session was not found")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("You are not authorized"))
		return
	}
	user, found := sessions[cookie.Value]
	if !found {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("You are not authorized"))
		return
	}
	SetupCORS(&w, r)
	bytes, err := json.Marshal(user)

	w.Write(bytes)
	w.WriteHeader(http.StatusOK)

}

func ChangeProfile(w http.ResponseWriter, r *http.Request) {
	logMethodAndURL(r)

	if r.Method == "OPTIONS" {
		SetupCORS(&w, r)
		w.WriteHeader(200)
		return
	}

	cookie, err := r.Cookie("sessionid")
	if err != nil {
		fmt.Println("Session was not found")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("You are not authorized"))
		return
	}
	_, found := sessions[cookie.Value]
	if !found {

		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("You are not authorized"))
		return
	} else {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		defer r.Body.Close()

		var userUpdates User
		err = json.Unmarshal(body, &userUpdates)
		sessions[cookie.Value] = userUpdates
		SetupCORS(&w, r)

		strVar, err := json.Marshal(userUpdates)
		if err != nil {

			panic(err)
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(strVar)
	}

}
