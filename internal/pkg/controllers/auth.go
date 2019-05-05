package controllers

import (
	"2019_1_undefined_penguins/internal/app/auth"
	"encoding/json"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
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

	//var user models.User
	var user *auth.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	grcpConn, err := grpc.Dial(
		"127.0.0.1:8083",
		grpc.WithInsecure(),
	)
	if err != nil {
		helpers.LogMsg("Can`t connect to grpc")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer grcpConn.Close()

	authManager := auth.NewAuthCheckerClient(grcpConn)
	ctx := context.Background()
	token, err := authManager.CreateUser(ctx, user)

	fmt.Println(err)

	cookie := &http.Cookie{
		Name:     "sessionid",
		Value:    token.Token,
		Expires:  time.Now().Add(time.Hour),
		HttpOnly: true,
	}

	user, _ = authManager.GetUser(ctx, token)
	bytes, err := json.Marshal(user)

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
	foundByEmail, _ := db.GetUserByEmail(user.Email)
	foundByLogin, _ := db.GetUserByLogin(user.Login)

	if foundByEmail != nil || foundByLogin != nil{
		w.WriteHeader(http.StatusConflict)
		return
	}

	user.HashPassword = helpers.HashPassword(user.Password)

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
	user.Picture = "http://localhost:8081/data/Default.png"
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
		w.WriteHeader(http.StatusUnauthorized)
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
