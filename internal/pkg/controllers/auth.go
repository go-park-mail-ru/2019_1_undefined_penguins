package controllers

import (
	"2019_1_undefined_penguins/internal/pkg/models"
	"fmt"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc/status"

	"io/ioutil"
	"net/http"

	"2019_1_undefined_penguins/internal/pkg/helpers"
)

func SignIn(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer r.Body.Close()

	var user *models.User
	easyJsonUser := models.ToEasyJsonUser(user)
	//err = json.Unmarshal(body, &user)
	err = easyJsonUser.UnmarshalJSON(body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	user = models.ToModelUser(easyJsonUser)
	ctx := context.Background()
	token, err := models.AuthManager.LoginUser(ctx, user)

	fmt.Println(err)
	if err != nil {
		switch errGRPC, _ := status.FromError(err); errGRPC.Code() {
		case 2:
			w.WriteHeader(http.StatusUnauthorized)
			return
		case 5:
			w.WriteHeader(http.StatusNotFound)
			return
		case 7:
			w.WriteHeader(http.StatusForbidden)
			return
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

	user, _ = models.AuthManager.GetUser(ctx, token)
	easyJsonUser = models.ToEasyJsonUser(user)
	//bytes, err := json.Marshal(user)
	bytes, err := easyJsonUser.MarshalJSON()

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

	var user *models.User
	easyJsonUser := models.ToEasyJsonUser(user)
	//err = json.Unmarshal(body, &user)
	err = easyJsonUser.UnmarshalJSON(body)

	if err != nil {
		helpers.LogMsg(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user = models.ToModelUser(easyJsonUser)
	ctx := context.Background()
	token, err := models.AuthManager.RegisterUser(ctx, user)

	fmt.Println(err)
	if err != nil {
		switch errGRPC, _ := status.FromError(err); errGRPC.Code() {
		case 2:
			w.WriteHeader(http.StatusUnauthorized)
			return
		case 5:
			w.WriteHeader(http.StatusNotFound)
			return
		case 6:
			w.WriteHeader(http.StatusConflict)
			return
		case 7:
			w.WriteHeader(http.StatusForbidden)
			return
		case 13:
			w.WriteHeader(http.StatusInternalServerError)
			return
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
	easyJsonUser = models.ToEasyJsonUser(user)
	//bytes, err := json.Marshal(user)
	bytes, err := easyJsonUser.MarshalJSON()

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

var defaultPictureAddress string

func SetDefaultPictureAddress(address string) {
	defaultPictureAddress = address
}
