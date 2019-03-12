package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	uuid "github.com/satori/uuid"
)

// later remove hardcode
type User struct {
	ID       uint   `json:"id"`
	Login    string `json:"login"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
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
		ID:       1,
		Login:    "Penguin1",
		Email:    "a.penguin1@corp.mail.ru",
		Name:     "Пингвин Северного Полюса",
		Password: "password",
		// LastVisit:  "25.02.2019",
		Score:      0,
		avatarName: "default1.png",
		avatarBlob: "./images/user.svg",
	},
	"b.penguin2@corp.mail.ru": User{
		ID:       2,
		Login:    "Penguin2",
		Email:    "b.penguin2@corp.mail.ru",
		Name:     "Пингвин Южного Полюса",
		Password: "password",
		// LastVisit:  "25.02.2019",
		Score:      100500,
		avatarName: "default2.png",
		avatarBlob: "./images/user.svg",
	},
	"c.penguin3@corp.mail.ru": User{
		ID:       3,
		Login:    "Penguin3",
		Email:    "c.penguin3@corp.mail.ru",
		Name:     "Залетный Пингвин",
		Password: "password",
		// LastVisit:  "25.02.2019",
		Score:      173,
		avatarName: "default3.png",
		avatarBlob: "./images/user.svg",
	},
	"d.penguin4@corp.mail.ru": User{
		ID:       4,
		Login:    "Penguin4",
		Email:    "d.penguin4@corp.mail.ru",
		Name:     "Рядовой Пингвин",
		Password: "password",
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
	email := r.FormValue("email")
	password := r.FormValue("password")
	user, found := users[email]
	if !found {
		//УТОЧНИТЬ У ФРОНТА КАКОЙ СТАТУС
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not found"))
		return
	}
	if user.Password != password {
		//УТОЧНИТЬ У ФРОНТА КАКОЙ СТАТУС
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Password is wrong"))
		return
	}
	sessionID, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}
	sessions[sessionID.String()] = users[email]
	http.SetCookie(w, &http.Cookie{
		Name:     "sessionid",
		Value:    sessionID.String(),
		Expires:  time.Now().Add(20 * time.Minute),
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
	})

	w.WriteHeader(http.StatusOK)
}

func SignOut(w http.ResponseWriter, r *http.Request) {
	if !isAuth(r) {
		models.SendMessage(w, http.StatusUnauthorized, "already signed out")
		return
	}

	jwtCookie, errNoCookie := r.Cookie(auth.CookieName)
	if errNoCookie != nil {
		models.SendMessage(w, http.StatusUnauthorized, "already signed out")
		return
	}

	jwtCookie.Expires = time.Unix(0, 0)
	http.SetCookie(w, jwtCookie)
	models.SendMessage(w, http.StatusOK, "signed out")
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
	if !found {
		users[userInfo.Email] = User{
			ID:       4,
			Login:    "Default",
			Email:    userInfo.Email,
			Name:     "Default",
			Password: userInfo.Password,
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

func IsAuth(w http.ResponseWriter, r *http.Request) {
}
