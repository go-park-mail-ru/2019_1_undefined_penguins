package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	uuid "github.com/satori/uuid"
)

func SignIn(w http.ResponseWriter, r *http.Request) {
	logMethodAndURL(r)
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {

			panic(err)
		}
		defer r.Body.Close()
		var userInfo SignUpStruct
		err = json.Unmarshal(body, &userInfo)
		if err != nil {
			panic(err)
		}

		user, found := users[userInfo.Email]
		if !found {

			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Not found"))
			return
		}

		if ComparePasswords(userInfo.Password, getPwd(user.HashPassword)) {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Password is wrong"))
			return
		}
		sessionID, err := uuid.NewV4()
		fmt.Println(sessionID.String())
		if err != nil {
			panic(err)
		}
		sessions[sessionID.String()] = users[userInfo.Email]
		CreateCookie(&w, sessionID.String())
		SetupCORS(&w, r)

		w.WriteHeader(http.StatusOK)
	} else {
		SetupCORS(&w, r)
		w.WriteHeader(200)
		return
	}

}

func SignUp(w http.ResponseWriter, r *http.Request) {
	logMethodAndURL(r)
	if r.Method == "OPTIONS" {
		SetupCORS(&w, r)
		w.WriteHeader(200)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	var userInfo SignUpStruct
	err = json.Unmarshal(body, &userInfo)

	if err != nil {
		panic(err)
	}

	_, found := users[userInfo.Email]
	hash := HashAndSalt(getPwd(userInfo.Password))

	if !found {

		users[userInfo.Email] = User{
			ID:           4,
			Login:        "Default",
			Email:        userInfo.Email,
			Name:         "Default",
			HashPassword: hash,
			// LastVisit:  "25.02.2019",
			Score:      0,
			avatarName: "default4.png",
			avatarBlob: "./images/user.svg",
		}
		sessionID, err := uuid.NewV4()
		if err != nil {
			panic(err)
		}
		CreateCookie(&w, sessionID.String())
		sessions[sessionID.String()] = users[userInfo.Email]
		SetupCORS(&w, r)
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("error"))
		return
	}

}

func SignOut(w http.ResponseWriter, r *http.Request) {
	logMethodAndURL(r)
	if r.Method == "OPTIONS" {
		SetupCORS(&w, r)
		w.WriteHeader(200)
		return
	}
	cookie, err := r.Cookie("sessionid")
	if err != nil {

		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("You are not authorized"))
		return
	}
	DeleteCookie(&w, cookie)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error"))
	}
	w.WriteHeader(http.StatusOK)

}
