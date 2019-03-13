package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	uuid "github.com/satori/uuid"
	"golang.org/x/crypto/bcrypt"
)

func setupResponse(w *http.ResponseWriter, r *http.Request) {
	origin := r.Header.Get("Origin")

	responseHeader := (*w).Header()
	responseHeader.Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	responseHeader.Set("Access-Control-Allow-Credentials", "true")

	responseHeader.Set("Access-Control-Allow-Headers", "Content-Type")

	responseHeader.Set("Access-Control-Allow-Origin", origin)
	(*w).WriteHeader(200)
}

func getPwd(pwd string) []byte {
	return []byte(pwd)
}

func HashAndSalt(pwd []byte) string {

	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}

	return string(hash)
}

func ComparePasswords(hashedPwd string, plainPwd []byte) bool {

	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

// later remove hardcode
type User struct {
	ID           uint   `json:"id"`
	Login        string `json:"login"`
	Email        string `json:"email"`
	Name         string `json:"name"`
	HashPassword string `json:"password"`
	// LastVisit  Time   `json:"lastVisit"`
	Score      uint   `json:"score"`
	avatarName string `json:"avatarName"`
	avatarBlob string `json:"avatarBlob"`
}

type SignUpStruct struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type MeStruct struct {
	Login string `json:"login"`
	Email string `json:"email"`
	Name  string `json:"name"`
	Score uint   `json:"score"`
}

var sessions map[string]User

func init() {
	sessions = make(map[string]User)

}

var users = map[string]User{
	"a.penguin1@corp.mail.ru": User{
		ID:           1,
		Login:        "Penguin1",
		Email:        "a.penguin1@corp.mail.ru",
		Name:         "Пингвин Северного Полюса",
		HashPassword: "$2a$14$9s00w8l7VKS2gRr2mtmg..1hvANedLWgmux3yOjkS80dTZlXLnKs2",
		// LastVisit:  "25.02.2019",
		Score:      0,
		avatarName: "default1.png",
		avatarBlob: "./images/user.svg",
	},
	"b.penguin2@corp.mail.ru": User{
		ID:           2,
		Login:        "Penguin2",
		Email:        "b.penguin2@corp.mail.ru",
		Name:         "Пингвин Южного Полюса",
		HashPassword: "$2a$14$9s00w8l7VKS2gRr2mtmg..1hvANedLWgmux3yOjkS80dTZlXLnKs2",
		// LastVisit:  "25.02.2019",
		Score:      100500,
		avatarName: "default2.png",
		avatarBlob: "./images/user.svg",
	},
	"c.penguin3@corp.mail.ru": User{
		ID:           3,
		Login:        "Penguin3",
		Email:        "c.penguin3@corp.mail.ru",
		Name:         "Залетный Пингвин",
		HashPassword: "$2a$14$9s00w8l7VKS2gRr2mtmg..1hvANedLWgmux3yOjkS80dTZlXLnKs2",
		// LastVisit:  "25.02.2019",
		Score:      173,
		avatarName: "default3.png",
		avatarBlob: "./images/user.svg",
	},
	"d.penguin4@corp.mail.ru": User{
		ID:           4,
		Login:        "Penguin4",
		Email:        "d.penguin4@corp.mail.ru",
		Name:         "Рядовой Пингвин",
		HashPassword: "$2a$04$U2BYDHAfGa2cqJwlhSA2D.XyWD8kq1sAvh2s8nRlV5huDEJLF8pDu",
		// LastVisit:  "25.02.2019",
		Score:      72,
		avatarName: "default4.png",
		avatarBlob: "./images/user.svg",
	},
}

// end later remove hardcode

func RootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello penguins"))
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	fmt.Println(r.RequestURI)
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {

			panic(err)
		}
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
			w.Write([]byte("HashPassword is wrong"))
			return
		}
		sessionID, err := uuid.NewV4()
		fmt.Println(sessionID.String())
		if err != nil {
			panic(err)
		}
		sessions[sessionID.String()] = users[userInfo.Email]
		http.SetCookie(w, &http.Cookie{
			Name:     "sessionid",
			Value:    sessionID.String(),
			Expires:  time.Now().Add(20 * time.Minute),
			Path:     "/",
			HttpOnly: true,
			// Secure:   true,
		})
		setupResponse(&w, r)

		w.WriteHeader(http.StatusOK)
	} else {
		setupResponse(&w, r)
		w.WriteHeader(200)
		return
	}

}

func SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		origin := r.Header.Get("Origin")

		responseHeader := w.Header()
		responseHeader.Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		responseHeader.Set("Access-Control-Allow-Credentials", "true")

		if accessControlRequestHeaders := r.Header.Get("Access-Control-Request-Headers"); accessControlRequestHeaders != "" {
			responseHeader.Set("Access-Control-Allow-Headers", accessControlRequestHeaders)
		}
		responseHeader.Set("Access-Control-Allow-Origin", origin)
		w.WriteHeader(200)
		return
	} else {
		fmt.Println(r.Method)
		fmt.Println(r.RequestURI)
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
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
			sessions[sessionID.String()] = users[userInfo.Email]
			http.SetCookie(w, &http.Cookie{
				Name:     "sessionid",
				Value:    sessionID.String(),
				Expires:  time.Now().Add(60 * time.Hour),
				HttpOnly: true,
				// Secure:   true,
			})
			setupResponse(&w, r)
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("error"))
			return
		}
	}

}

func SignOut(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	fmt.Println(r.RequestURI)
	cookie, err := r.Cookie("sessionid")
	if err != nil {

		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("You are not authorized"))
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "sessionid",
		Value:    "",
		Expires:  time.Unix(0, 0),
		Path:     "/",
		HttpOnly: true,
		// Secure:   true,
	})
	delete(sessions, cookie.Value)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error"))
	}
	w.WriteHeader(http.StatusOK)

}
