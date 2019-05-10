package controllers

import (
	"2019_1_undefined_penguins/internal/pkg/models"
	"encoding/json"
	"fmt"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"

	"io/ioutil"
	"net/http"

	"2019_1_undefined_penguins/internal/pkg/helpers"
	//"2019_1_undefined_penguins/internal/pkg/models"
)

func SignIn(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer r.Body.Close()

	//var user models.User
	var user *models.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	grcpConn, err := grpc.Dial(
		authAddress,
		grpc.WithInsecure(),
	)
	if err != nil {
		helpers.LogMsg("Can`t connect to grpc")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer grcpConn.Close()

	authManager := models.NewAuthCheckerClient(grcpConn)
	ctx := context.Background()
	token, err := authManager.LoginUser(ctx, user)

	fmt.Println(err)
	if err != nil {
		switch errGRPC, _ := status.FromError(err); errGRPC.Code() {
		// case 2:
		// 	w.WriteHeader(http.StatusUnauthorized)
		// 	return
		// case 5:
		// 	w.WriteHeader(http.StatusNotFound)
		// 	return
		// case 7:
		// 	w.WriteHeader(http.StatusForbidden)
		// 	return
		default:
			helpers.LogMsg("Unknown gprc error")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

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

	//var user models.User
	var user *models.User
	err = json.Unmarshal(body, &user)

	if err != nil {
		helpers.LogMsg(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	grcpConn, err := grpc.Dial(
		authAddress,
		grpc.WithInsecure(),
	)
	if err != nil {
		helpers.LogMsg("Can`t connect to grpc")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer grcpConn.Close()

	authManager := models.NewAuthCheckerClient(grcpConn)
	ctx := context.Background()
	token, err := authManager.RegisterUser(ctx, user)

	fmt.Println(err)
	if err != nil {
		switch errGRPC, _ := status.FromError(err); errGRPC.Code() {
		// case 2:
		// 	w.WriteHeader(http.StatusUnauthorized)
		// 	return
		// case 5:
		// 	w.WriteHeader(http.StatusNotFound)
		// 	return
		case 6:
			w.WriteHeader(http.StatusConflict)
			return
		// case 7:
		// 	w.WriteHeader(http.StatusForbidden)
		// 	return
		// case 13:
		// 	w.WriteHeader(http.StatusInternalServerError)
		// 	return
		default:
			helpers.LogMsg("Unknown gprc error")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	cookie := &http.Cookie{
		Name:     "sessionid",
		Value:    token.Token,
		Expires:  time.Now().Add(time.Hour),
		HttpOnly: true,
	}

	user.Password = ""
	user.Picture = defaultPictureAddress
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

var authAddress string
var defaultPictureAddress string

func SetAuthAddress(address string) {
	authAddress = address
}

func SetDefaultPictureAddress(address string) {
	defaultPictureAddress = address
}
