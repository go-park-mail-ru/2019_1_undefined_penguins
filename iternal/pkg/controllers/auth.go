package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	uuid "github.com/satori/uuid"
	"golang.org/x/crypto/bcrypt"
)

func setupResponse(w *http.ResponseWriter, r *http.Request) {
	(*w).Header().Set("Access-Controle-Allow-Origin", "*")
	(*w).Header().Set("Access-Controle-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	fmt.Print("hi")
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
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
		HashPassword: "$2a$14$GrGEHcSfqNBiB/rcM.um8edRSbsgaW/e5kgejC5stKh9oZK5LcksK",
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
		HashPassword: "$2a$14$GrGEHcSfqNBiB/rcM.um8edRSbsgaW/e5kgejC5stKh9oZK5LcksK",
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
		HashPassword: "$2a$14$GrGEHcSfqNBiB/rcM.um8edRSbsgaW/e5kgejC5stKh9oZK5LcksK",
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
		HashPassword: "$2a$14$GrGEHcSfqNBiB/rcM.um8edRSbsgaW/e5kgejC5stKh9oZK5LcksK",
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
	//временно для тестов, перепилить на body
	// email := r.FormValue("email")
	// password := r.FormValue("password")
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
		//УТОЧНИТЬ У ФРОНТА КАКОЙ СТАТУС
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not found"))
		return
	}
	if CheckPasswordHash(userInfo.Password, user.HashPassword) {
		//УТОЧНИТЬ У ФРОНТА КАКОЙ СТАТУС
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("HashPassword is wrong"))
		return
	}
	sessionID, err := uuid.NewV4()
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
		Secure:   true,
	})
	setupResponse(&w, r)
	w.WriteHeader(http.StatusOK)
}

func SignUp(w http.ResponseWriter, r *http.Request) {
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
	hash, err := HashPassword(userInfo.Password)
	if !found && err != nil {
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
		sessionId, err := uuid.NewV4()
		if err != nil {
			panic(err)
		}
		sessions[sessionId.String()] = users[userInfo.Email]
		http.SetCookie(w, &http.Cookie{
			Name:     "sessionid",
			Value:    sessionId.String(),
			Expires:  time.Now().Add(60 * time.Hour),
			Path:     "/",
			HttpOnly: true,
			Secure:   true,
		})
		w.WriteHeader(http.StatusOK)
	} else {
		//УТОЧНИТЬ У ФРОНТА КАКОЙ СТАТУС
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("error"))
		return
	}
}

func SignOut(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("sessionid")
	if err != nil {
		//УТОЧНИТЬ У ФРОНТА КАКОЙ СТАТУС
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
		Secure:   true,
	})
	delete(sessions, cookie.Value)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error"))
	}
	w.WriteHeader(http.StatusOK)

}
