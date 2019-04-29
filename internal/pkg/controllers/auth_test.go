package controllers

import (
	dab "2019_1_undefined_penguins/internal/pkg/database"
	"2019_1_undefined_penguins/internal/pkg/helpers"
	"2019_1_undefined_penguins/internal/pkg/models"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

func GetUserFromJSON(fileName string) (*models.User, error) {
	dir, err := os.Getwd()
	if err != nil {
		helpers.LogMsg("Getting directory error: ", err)
		return nil, err

	}
	dir = strings.Replace(dir, "/internal/pkg/controllers", "", -1)
	file, err := os.Open(dir + "/configs/" + fileName)
	if err != nil {
		helpers.LogMsg("Open directory error: ", err)
		return nil, err
	}

	body, _ := ioutil.ReadAll(file)
	var user models.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		helpers.LogMsg("Init parse user error: ", err)
		return nil, err
	}
	return &user, nil
}

func TestSignIn(t *testing.T) {
	var user models.User
	user.Login = time.Now().Format("20060102150405") + user.Login
	user.Email = time.Now().Format("20060102150405") + user.Email
	user.Password = "password"
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	dab.SetMock(db)
	rows := sqlmock.NewRows([]string{"id", "login", "email", "hashpassword", "score", "name", "games"}).
		AddRow(1, "login", "login@mail.ru", "$2a$08$Q1nN3cy96NhOW7jOx31atuzY.QuRXbnWRitfkwZDHbC3dY83bw53i", 10, "name", "5").
		RowError(1, fmt.Errorf("error"))
	mock.ExpectQuery("SELECT").WillReturnRows(rows)
	data, err := json.Marshal(user)
	if err != nil {
		t.Error(err)
	}
	buf := bytes.NewBuffer(data)
	req, err := http.NewRequest("POST", "/login", buf)
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(SignIn)
	handler.ServeHTTP(w, req)
	expectedStatus := http.StatusOK
	if w.Code != expectedStatus {
		t.Error(w.Code)
	}

	rows = sqlmock.NewRows([]string{"id", "login", "email", "hashpassword", "score", "name", "games"}).
		AddRow(1, "login", "login@mail.ru", "$2a$08$.QuRXbnWRitfkwZDHbC3dY83bw53i", 10, "name", "5").
		RowError(1, fmt.Errorf("error"))
	mock.ExpectQuery("SELECT").WillReturnRows(rows)
	data, err = json.Marshal(user)
	if err != nil {
		t.Error(err)
	}
	buf = bytes.NewBuffer(data)
	req, err = http.NewRequest("POST", "/login", buf)
	w = httptest.NewRecorder()
	handler = http.HandlerFunc(SignIn)
	handler.ServeHTTP(w, req)
	expectedStatus = http.StatusForbidden
	if w.Code != expectedStatus {
		t.Error(w.Code)
	}

	data, err = json.Marshal(user)
	if err != nil {
		t.Error(err)
	}
	buf = bytes.NewBuffer(data)
	req, err = http.NewRequest("POST", "/login", buf)
	w = httptest.NewRecorder()
	handler = http.HandlerFunc(SignIn)
	handler.ServeHTTP(w, req)
	expectedStatus = http.StatusNotFound
	if w.Code != expectedStatus {
		t.Error(w.Code)
	}

	data, err = json.Marshal("")
	if err != nil {
		t.Error(err)
	}
	buf = bytes.NewBuffer(data)
	req, err = http.NewRequest("POST", "/login", buf)
	w = httptest.NewRecorder()
	handler = http.HandlerFunc(SignIn)
	handler.ServeHTTP(w, req)
	expectedStatus = http.StatusInternalServerError
	if w.Code != expectedStatus {
		t.Error(w.Code)
	}
}

func TestSignUp(t *testing.T) {
	var user models.User
	user.Login = time.Now().Format("20060102150405") + user.Login
	user.Email = time.Now().Format("20060102150405") + user.Email
	user.Password = "password"
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	dab.SetMock(db)
	rows := sqlmock.NewRows([]string{"id", "login", "email", "hashpassword", "score", "name", "games"}).
		AddRow(1, "login", "login@mail.ru", "$2a$08$Q1nN3cy96NhOW7jOx31atuzY.QuRXbnWRitfkwZDHbC3dY83bw53i", 10, "name", "5").
		RowError(1, fmt.Errorf("error"))
	mock.ExpectQuery("SELECT").WillReturnRows(rows)
	data, err := json.Marshal(user)
	if err != nil {
		t.Error(err)
	}
	buf := bytes.NewBuffer(data)
	req, err := http.NewRequest("POST", "/signup", buf)
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(SignUp)
	handler.ServeHTTP(w, req)
	expectedStatus := http.StatusConflict
	if w.Code != expectedStatus {
		t.Error(w.Code)
	}

	data, err = json.Marshal("")
	if err != nil {
		t.Error(err)
	}
	buf = bytes.NewBuffer(data)
	req, err = http.NewRequest("POST", "/login", buf)
	w = httptest.NewRecorder()
	handler = http.HandlerFunc(SignUp)
	handler.ServeHTTP(w, req)
	expectedStatus = http.StatusInternalServerError
	if w.Code != expectedStatus {
		t.Error(w.Code)
	}

	rows = sqlmock.NewRows([]string{"id", "login", "score"}).
		AddRow(1, "login", 10).
		RowError(1, fmt.Errorf("error"))
	mock.ExpectQuery("SELECT").WillReturnError(fmt.Errorf("some error"))
	mock.ExpectQuery("INSERT").WillReturnRows(rows)
	data, err = json.Marshal(user)
	if err != nil {
		t.Error(err)
	}
	buf = bytes.NewBuffer(data)
	req, err = http.NewRequest("POST", "/signup", buf)
	w = httptest.NewRecorder()
	handler = http.HandlerFunc(SignUp)
	handler.ServeHTTP(w, req)
	expectedStatus = http.StatusOK
	if w.Code != expectedStatus {
		t.Error(w.Code)
	}

	rows = sqlmock.NewRows([]string{"id", "login", "score"}).
		AddRow(1, "login", 10).
		RowError(1, fmt.Errorf("error"))
	mock.ExpectQuery("SELECT").WillReturnError(fmt.Errorf("some error"))
	mock.ExpectQuery("INSERT").WillReturnError(fmt.Errorf("some error"))
	data, err = json.Marshal(user)
	if err != nil {
		t.Error(err)
	}
	buf = bytes.NewBuffer(data)
	req, err = http.NewRequest("POST", "/signup", buf)
	w = httptest.NewRecorder()
	handler = http.HandlerFunc(SignUp)
	handler.ServeHTTP(w, req)
	expectedStatus = http.StatusInternalServerError
	if w.Code != expectedStatus {
		t.Error(w.Code)
	}
	req, err = http.NewRequest("GET", "/", nil)
	handler = http.HandlerFunc(RootHandler)
	handler.ServeHTTP(w, req)

}

func TestMe(t *testing.T) {
	var user models.User
	user.Login = time.Now().Format("20060102150405") + user.Login
	user.Email = time.Now().Format("20060102150405") + user.Email
	user.Password = "password"
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	dab.SetMock(db)
	rows := sqlmock.NewRows([]string{"id", "login", "score"}).
		AddRow(1, "login", 10).
		RowError(1, fmt.Errorf("error"))
	mock.ExpectQuery("SELECT").WillReturnError(fmt.Errorf("some error"))
	mock.ExpectQuery("INSERT").WillReturnRows(rows)
	data, err := json.Marshal(user)
	if err != nil {
		t.Error(err)
	}
	buf := bytes.NewBuffer(data)
	req, err := http.NewRequest("POST", "/signup", buf)
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(SignUp)
	handler.ServeHTTP(w, req)
	expectedStatus := http.StatusOK
	if w.Code != expectedStatus {
		t.Error(w.Code)
	}

	var ss []string

	s := w.HeaderMap["Set-Cookie"][0]

	ss = strings.Split(s, ";")

	cookieInfo := strings.Split(ss[0], "=")
	cookieExpires := strings.Split(ss[1], "=")
	timeStampString := cookieExpires[1]
	layOut := "Mon, 2 Jan 2006 15:04:05 GMT"
	timeStamp, err := time.Parse(layOut, timeStampString)
	if err != nil {
		t.Error(err)
	}
	cookie := &http.Cookie{
		Name:     cookieInfo[0],
		Value:    cookieInfo[1],
		Expires:  timeStamp,
		HttpOnly: true,
	}
	req, err = http.NewRequest("POST", "/me", buf)
	w = httptest.NewRecorder()
	handler = http.HandlerFunc(Me)
	req.AddCookie(cookie)
	mock.ExpectQuery("SELECT").WillReturnError(fmt.Errorf("some error"))

	handler.ServeHTTP(w, req)

	expectedStatus = http.StatusNotFound
	if w.Code != expectedStatus {
		t.Error(w.Code)
	}
	rows = sqlmock.NewRows([]string{"id", "login", "email", "hashpassword", "score", "name", "games"}).
		AddRow(1, "login", "login@mail.ru", "hdfbkfbdj", 10, "name", "5").
		RowError(1, fmt.Errorf("error"))
	mock.ExpectQuery("SELECT").WillReturnRows(rows)
	handler.ServeHTTP(w, req)
	expectedStatus = http.StatusNotFound
	if w.Code != expectedStatus {
		t.Error(w.Code)
	}

	rows = sqlmock.NewRows([]string{"games", "name", "score"}).
		AddRow(10, "name", 5).
		RowError(1, fmt.Errorf("error"))
	mock.ExpectQuery("UPDATE").WillReturnError(fmt.Errorf("some error"))
	data, err = json.Marshal(user)
	if err != nil {
		t.Error(err)
	}
	buf = bytes.NewBuffer(data)
	req, err = http.NewRequest("POST", "/update", buf)
	w = httptest.NewRecorder()
	handler = http.HandlerFunc(ChangeProfile)
	req.AddCookie(cookie)

	handler.ServeHTTP(w, req)
	expectedStatus = http.StatusOK
	if w.Code != expectedStatus {
		t.Error(w.Code)
	}
}

// func TestLogIn(t *testing.T) {
// 	err := database.Connect()
// 	if err != nil {
// 		helpers.LogMsg("Connection error: ", err)
// 		t.Error(err)
// 	}
// 	defer database.Disconnect()
// 	user, err := GetUserFromJSON("testuser.json")
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	data, err := json.Marshal(user)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	buf := bytes.NewBuffer(data)
// 	req, err := http.NewRequest("POST", "/login", buf)
// 	w := httptest.NewRecorder()
// 	handler := http.HandlerFunc(SignIn)
// 	handler.ServeHTTP(w, req)
// 	expectedStatus := http.StatusOK
// 	if w.Code != expectedStatus {
// 		t.Error(w.Code)
// 	}

// }

// func TestSignUp(t *testing.T) {

// user, err := GetUserFromJSON("wronguser.json")
// if err != nil {
// 	t.Error(err)
// }
// user.Login = time.Now().Format("20060102150405") + user.Login
// user.Email = time.Now().Format("20060102150405") + user.Email
// user.Password = time.Now().Format("20060102150405") + user.Password
// data, err := json.Marshal(user)
// if err != nil {
// 	t.Error(err)
// }
// buf := bytes.NewBuffer(data)
// req, err := http.NewRequest("POST", "/signup", buf)
// w := httptest.NewRecorder()
// handler := http.HandlerFunc(SignUp)
// handler.ServeHTTP(w, req)
// expectedStatus := http.StatusOK
// if w.Code != expectedStatus {
// 	t.Error(w.Code)
// }
// }

// func TestLogInWrongPassword(t *testing.T) {
// 	err := database.Connect()
// 	if err != nil {
// 		helpers.LogMsg("Connection error: ", err)
// 		t.Error(err)
// 	}
// 	defer database.Disconnect()
// 	user, err := GetUserFromJSON("testuser.json")
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	user.Password = "password"
// 	data, err := json.Marshal(user)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	buf := bytes.NewBuffer(data)
// 	req, err := http.NewRequest("POST", "/login", buf)
// 	w := httptest.NewRecorder()
// 	handler := http.HandlerFunc(SignIn)
// 	handler.ServeHTTP(w, req)
// 	expectedStatus := http.StatusForbidden
// 	if w.Code != expectedStatus {
// 		t.Error(w.Code)
// 	}

// }

// func TestEmptyLogIn(t *testing.T) {
// 	err := database.Connect()
// 	if err != nil {
// 		helpers.LogMsg("Connection error: ", err)
// 		t.Error(err)
// 	}
// 	defer database.Disconnect()
// 	buf := bytes.NewBuffer(nil)
// 	req, err := http.NewRequest("POST", "/login", buf)
// 	w := httptest.NewRecorder()
// 	handler := http.HandlerFunc(SignIn)
// 	handler.ServeHTTP(w, req)
// 	expectedStatus := http.StatusInternalServerError
// 	if w.Code != expectedStatus {
// 		t.Error(w.Code)
// 	}

// }

// func TestWrongUserSignIn(t *testing.T) {
// 	err := database.Connect()
// 	if err != nil {
// 		helpers.LogMsg("Connection error: ", err)
// 		t.Error(err)
// 	}
// 	defer database.Disconnect()
// 	user, err := GetUserFromJSON("wronguser.json")
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	data, err := json.Marshal(user)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	buf := bytes.NewBuffer(data)
// 	req, err := http.NewRequest("POST", "/login", buf)
// 	w := httptest.NewRecorder()
// 	handler := http.HandlerFunc(SignIn)
// 	handler.ServeHTTP(w, req)
// 	expectedStatus := http.StatusNotFound
// 	if w.Code != expectedStatus {
// 		t.Error(w.Code)
// 	}

// }

// func TestWrongUserSignOut(t *testing.T) {
// 	err := database.Connect()
// 	if err != nil {
// 		helpers.LogMsg("Connection error: ", err)
// 		t.Error(err)
// 	}
// 	defer database.Disconnect()
// 	user, err := GetUserFromJSON("wronguser.json")
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	data, err := json.Marshal(user)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	buf := bytes.NewBuffer(data)
// 	req, err := http.NewRequest("POST", "/login", buf)
// 	w := httptest.NewRecorder()
// 	handler := http.HandlerFunc(SignIn)
// 	handler.ServeHTTP(w, req)
// 	expectedStatus := http.StatusNotFound
// 	if w.Code != expectedStatus {
// 		t.Error(w.Code)
// 	}

// 	req, err = http.NewRequest("POST", "/signout", buf)
// 	w = httptest.NewRecorder()
// 	handler = http.HandlerFunc(SignOut)
// 	handler.ServeHTTP(w, req)
// 	expectedStatus = http.StatusNotFound
// 	if w.Code != expectedStatus {
// 		t.Error(w.Code)
// 	}

// }

// func TestMe(t *testing.T) {
// 	err := database.Connect()
// 	if err != nil {
// 		helpers.LogMsg("Connection error: ", err)
// 		t.Error(err)
// 	}
// 	defer database.Disconnect()
// 	user, err := GetUserFromJSON("testuser.json")
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	data, err := json.Marshal(user)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	buf := bytes.NewBuffer(data)
// 	req, err := http.NewRequest("POST", "/login", buf)
// 	w := httptest.NewRecorder()
// 	handler := http.HandlerFunc(SignIn)
// 	handler.ServeHTTP(w, req)
// 	expectedStatus := http.StatusOK
// 	if w.Code != expectedStatus {
// 		t.Error(w.Code)
// 	}
// 	var ss []string

// 	s := w.HeaderMap["Set-Cookie"][0]

// 	ss = strings.Split(s, ";")

// 	cookieInfo := strings.Split(ss[0], "=")
// 	cookieExpires := strings.Split(ss[1], "=")
// 	timeStampString := cookieExpires[1]
// 	layOut := "Mon, 2 Jan 2006 15:04:05 GMT"
// 	timeStamp, err := time.Parse(layOut, timeStampString)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	cookie := &http.Cookie{
// 		Name:     cookieInfo[0],
// 		Value:    cookieInfo[1],
// 		Expires:  timeStamp,
// 		HttpOnly: true,
// 	}
// 	req, err = http.NewRequest("POST", "/me", buf)
// 	w = httptest.NewRecorder()
// 	handler = http.HandlerFunc(Me)
// 	req.AddCookie(cookie)
// 	handler.ServeHTTP(w, req)
// 	expectedStatus = http.StatusOK
// 	if w.Code != expectedStatus {
// 		t.Error(w.Code)
// 	}
// }

// func TestSignOut(t *testing.T) {
// 	err := database.Connect()
// 	if err != nil {
// 		helpers.LogMsg("Connection error: ", err)
// 		t.Error(err)
// 	}
// 	defer database.Disconnect()
// 	user, err := GetUserFromJSON("testuser.json")
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	data, err := json.Marshal(user)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	buf := bytes.NewBuffer(data)
// 	req, err := http.NewRequest("POST", "/login", buf)
// 	w := httptest.NewRecorder()
// 	handler := http.HandlerFunc(SignIn)
// 	handler.ServeHTTP(w, req)
// 	expectedStatus := http.StatusOK
// 	if w.Code != expectedStatus {
// 		t.Error(w.Code)
// 	}
// 	var ss []string

// 	s := w.HeaderMap["Set-Cookie"][0]

// 	ss = strings.Split(s, ";")

// 	cookieInfo := strings.Split(ss[0], "=")
// 	cookieExpires := strings.Split(ss[1], "=")
// 	timeStampString := cookieExpires[1]
// 	layOut := "Mon, 2 Jan 2006 15:04:05 GMT"
// 	timeStamp, err := time.Parse(layOut, timeStampString)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	cookie := &http.Cookie{
// 		Name:     cookieInfo[0],
// 		Value:    cookieInfo[1],
// 		Expires:  timeStamp,
// 		HttpOnly: true,
// 	}
// 	req, err = http.NewRequest("POST", "/signout", buf)
// 	w = httptest.NewRecorder()
// 	handler = http.HandlerFunc(SignOut)
// 	req.AddCookie(cookie)
// 	handler.ServeHTTP(w, req)
// 	expectedStatus = http.StatusOK
// 	if w.Code != expectedStatus {
// 		t.Error(w.Code)
// 	}
// }

// func TestSignUpConflict(t *testing.T) {
// 	err := database.Connect()
// 	if err != nil {
// 		helpers.LogMsg("Connection error: ", err)
// 		t.Error(err)
// 	}
// 	defer database.Disconnect()
// 	user, err := GetUserFromJSON("testuser.json")
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	data, err := json.Marshal(user)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	buf := bytes.NewBuffer(data)
// 	req, err := http.NewRequest("POST", "/signup", buf)
// 	w := httptest.NewRecorder()
// 	handler := http.HandlerFunc(SignUp)
// 	handler.ServeHTTP(w, req)
// 	expectedStatus := http.StatusConflict
// 	if w.Code != expectedStatus {
// 		t.Error(w.Code)
// 	}
// }

// func TestSignUpUpdate(t *testing.T) {
// 	err := database.Connect()
// 	if err != nil {
// 		helpers.LogMsg("Connection error: ", err)
// 		t.Error(err)
// 	}
// 	defer database.Disconnect()
// 	user, err := GetUserFromJSON("wronguser.json")
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	user.Login = time.Now().Format("20060102150405") + user.Login + "2"
// 	user.Email = time.Now().Format("20060102150405") + user.Email + "2"
// 	user.Password = time.Now().Format("20060102150405") + user.Password + "2"
// 	data, err := json.Marshal(user)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	buf := bytes.NewBuffer(data)
// 	req, err := http.NewRequest("POST", "/signup", buf)
// 	w := httptest.NewRecorder()
// 	handler := http.HandlerFunc(SignUp)
// 	handler.ServeHTTP(w, req)
// 	expectedStatus := http.StatusOK
// 	if w.Code != expectedStatus {
// 		t.Error(w.Code)
// 	}
// 	var ss []string

// 	s := w.HeaderMap["Set-Cookie"][0]

// 	ss = strings.Split(s, ";")

// 	cookieInfo := strings.Split(ss[0], "=")
// 	cookieExpires := strings.Split(ss[1], "=")
// 	timeStampString := cookieExpires[1]
// 	layOut := "Mon, 2 Jan 2006 15:04:05 GMT"
// 	timeStamp, err := time.Parse(layOut, timeStampString)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	cookie := &http.Cookie{
// 		Name:     cookieInfo[0],
// 		Value:    cookieInfo[1],
// 		Expires:  timeStamp,
// 		HttpOnly: true,
// 	}
// 	user.Login = time.Now().Format("20060102150405") + user.Login
// 	user.Email = time.Now().Format("20060102150405") + user.Email
// 	user.Password = time.Now().Format("20060102150405") + user.Password
// 	data, err = json.Marshal(user)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	buf = bytes.NewBuffer(data)
// 	req, err = http.NewRequest("POST", "/change_profile", buf)
// 	w = httptest.NewRecorder()
// 	handler = http.HandlerFunc(ChangeProfile)
// 	req.AddCookie(cookie)
// 	handler.ServeHTTP(w, req)
// 	expectedStatus = http.StatusOK
// 	if w.Code != expectedStatus {
// 		t.Error(w.Code)
// 	}
// }

// func TestSignInConflict(t *testing.T) {
// 	err := database.Connect()
// 	if err != nil {
// 		helpers.LogMsg("Connection error: ", err)
// 		t.Error(err)
// 	}
// 	defer database.Disconnect()
// 	user, err := GetUserFromJSON("testuser.json")
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	data, err := json.Marshal(user)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	buf := bytes.NewBuffer(data)
// 	req, err := http.NewRequest("POST", "/signin", buf)
// 	w := httptest.NewRecorder()
// 	handler := http.HandlerFunc(SignIn)
// 	handler.ServeHTTP(w, req)
// 	expectedStatus := http.StatusOK
// 	if w.Code != expectedStatus {
// 		t.Error(w.Code)
// 	}
// 	var ss []string

// 	s := w.HeaderMap["Set-Cookie"][0]

// 	ss = strings.Split(s, ";")

// 	cookieInfo := strings.Split(ss[0], "=")
// 	cookieExpires := strings.Split(ss[1], "=")
// 	timeStampString := cookieExpires[1]
// 	layOut := "Mon, 2 Jan 2006 15:04:05 GMT"
// 	timeStamp, err := time.Parse(layOut, timeStampString)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	cookie := &http.Cookie{
// 		Name:     cookieInfo[0],
// 		Value:    cookieInfo[1],
// 		Expires:  timeStamp,
// 		HttpOnly: true,
// 	}
// 	user.Login = "iamfrommoscow"
// 	user.Email = time.Now().Format("20060102150405") + user.Email
// 	user.Password = time.Now().Format("20060102150405") + user.Password
// 	data, err = json.Marshal(user)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	buf = bytes.NewBuffer(data)
// 	req, err = http.NewRequest("POST", "/change_profile", buf)
// 	w = httptest.NewRecorder()
// 	handler = http.HandlerFunc(ChangeProfile)
// 	req.AddCookie(cookie)
// 	handler.ServeHTTP(w, req)
// 	expectedStatus = http.StatusConflict
// 	if w.Code != expectedStatus {
// 		t.Error(w.Code)
// 	}
// }

// func TestHome(t *testing.T) {
// 	req, _ := http.NewRequest("GET", "/", nil)
// 	w := httptest.NewRecorder()
// 	handler := http.HandlerFunc(RootHandler)
// 	handler.ServeHTTP(w, req)
// 	expectedStatus := http.StatusOK
// 	if w.Code != expectedStatus {
// 		t.Error(w.Code)
// 	}
// }

// func TestUploadImageInternalServerError(t *testing.T) {
// 	err := database.Connect()
// 	if err != nil {
// 		helpers.LogMsg("Connection error: ", err)
// 		t.Error(err)
// 	}
// 	defer database.Disconnect()
// 	user, err := GetUserFromJSON("testuser.json")
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	data, err := json.Marshal(user)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	buf := bytes.NewBuffer(data)
// 	req, err := http.NewRequest("POST", "/login", buf)
// 	w := httptest.NewRecorder()
// 	handler := http.HandlerFunc(SignIn)
// 	handler.ServeHTTP(w, req)
// 	expectedStatus := http.StatusOK
// 	if w.Code != expectedStatus {
// 		t.Error(w.Code)
// 	}
// 	var ss []string

// 	s := w.HeaderMap["Set-Cookie"][0]

// 	ss = strings.Split(s, ";")

// 	cookieInfo := strings.Split(ss[0], "=")
// 	cookieExpires := strings.Split(ss[1], "=")
// 	timeStampString := cookieExpires[1]
// 	layOut := "Mon, 2 Jan 2006 15:04:05 GMT"
// 	timeStamp, err := time.Parse(layOut, timeStampString)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	cookie := &http.Cookie{
// 		Name:     cookieInfo[0],
// 		Value:    cookieInfo[1],
// 		Expires:  timeStamp,
// 		HttpOnly: true,
// 	}

// 	if err != nil {
// 		t.Error(err)
// 	}
// 	req, err = http.NewRequest("POST", "/change_profile", nil)
// 	w = httptest.NewRecorder()
// 	handler = http.HandlerFunc(UploadImage)
// 	req.AddCookie(cookie)
// 	handler.ServeHTTP(w, req)
// 	expectedStatus = http.StatusInternalServerError
// 	if w.Code != expectedStatus {
// 		t.Error(w.Code)
// 	}
// }

// func TestStartWS(t *testing.T) {
// 	err := database.Connect()
// 	if err != nil {
// 		helpers.LogMsg("Connection error: ", err)
// 		t.Error(err)
// 	}
// 	defer database.Disconnect()
// 	req, err := http.NewRequest("POST", "/ws", nil)
// 	w := httptest.NewRecorder()
// 	handler := http.HandlerFunc(StartWS)
// 	handler.ServeHTTP(w, req)
// 	expectedStatus := http.StatusUnauthorized
// 	if w.Code != expectedStatus {
// 		t.Error(w.Code)
// 	}
// }
