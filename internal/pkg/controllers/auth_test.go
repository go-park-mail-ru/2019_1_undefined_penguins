package controllers

import (
	"2019_1_undefined_penguins/internal/pkg/models"
	"bytes"
	"encoding/json"
	"strings"
	"time"

	// "fmt"
	// "io/ioutil"
	"net/http"
	"net/http/httptest"

	// "os"

	"testing"
	// "google.golang.org/grpc"
	// "github.com/DATA-DOG/go-sqlmock"
)

func TestRoot(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	handler := http.HandlerFunc(RootHandler)
	handler.ServeHTTP(w, req)
}

func TestSignUp(t *testing.T) {
	data, err := json.Marshal("hello")
	if err != nil {
		t.Error(err)
	}
	buf := bytes.NewBuffer(data)
	req, err := http.NewRequest("POST", "/signup", buf)
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(SignUp)
	handler.ServeHTTP(w, req)
	expectedStatus := http.StatusInternalServerError
	if w.Code != expectedStatus {
		t.Error(w.Code)
	}
}

func TestSignInMeAndSignOut(t *testing.T) {
	var user models.User
	user.Email = "test@test.te"
	user.Password = "testtest"
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

	buf = bytes.NewBuffer(data)
	req, err = http.NewRequest("POST", "/signin", buf)
	w = httptest.NewRecorder()
	handler = http.HandlerFunc(SignIn)
	handler.ServeHTTP(w, req)
	expectedStatus = http.StatusOK
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
	handler.ServeHTTP(w, req)
	expectedStatus = http.StatusOK
	if w.Code != expectedStatus {
		t.Error(w.Code)
	}
	req, _ = http.NewRequest("GET", "/signout", nil)
	req.AddCookie(cookie)
	w = httptest.NewRecorder()
	handler = http.HandlerFunc(SignOut)
	handler.ServeHTTP(w, req)
	expectedStatus = http.StatusOK
	if w.Code != expectedStatus {
		t.Error(w.Code)
	}
}

func TestSignOut(t *testing.T) {

	req, _ := http.NewRequest("GET", "/signout", nil)
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(SignOut)
	handler.ServeHTTP(w, req)
	expectedStatus := 401
	if w.Code != expectedStatus {
		t.Error(w.Code)
	}
}

func TestUpdateUser(t *testing.T) {
	var user models.User
	user.Email = "test@test.te"
	user.Password = "testtest"
	data, err := json.Marshal(user)
	if err != nil {
		t.Error(err)
	}
	buf := bytes.NewBuffer(data)
	req, err := http.NewRequest("POST", "/signin", buf)
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(SignIn)
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
	user.Email = "test@teste.te"
	data, err = json.Marshal(user)
	if err != nil {
		t.Error(err)
	}
	buf = bytes.NewBuffer(data)
	req, err = http.NewRequest("PUT", "/me", buf)
	req.AddCookie(cookie)
	w = httptest.NewRecorder()
	handler = http.HandlerFunc(ChangeProfile)
	handler.ServeHTTP(w, req)
	expectedStatus = http.StatusOK
	if w.Code != expectedStatus {
		t.Error(w.Code)
	}
	user.Email = "test@test.te"
	data, err = json.Marshal(user)
	if err != nil {
		t.Error(err)
	}
	buf = bytes.NewBuffer(data)
	req, err = http.NewRequest("PUT", "/me", buf)
	req.AddCookie(cookie)
	w = httptest.NewRecorder()
	handler = http.HandlerFunc(ChangeProfile)
	handler.ServeHTTP(w, req)
	expectedStatus = http.StatusOK
	if w.Code != expectedStatus {
		t.Error(w.Code)
	}

	handler = http.HandlerFunc(UploadImage)
	req.AddCookie(cookie)
	handler.ServeHTTP(w, req)

}

// func GetUserFromJSON(fileName string) (*models.User, error) {
// 	dir, err := os.Getwd()
// 	if err != nil {
// 		helpers.LogMsg("Getting directory error: ", err)
// 		return nil, err

// 	}
// 	dir = strings.Replace(dir, "/internal/pkg/controllers", "", -1)
// 	file, err := os.Open(dir + "/configs/" + fileName)
// 	if err != nil {
// 		helpers.LogMsg("Open directory error: ", err)
// 		return nil, err
// 	}

// 	body, _ := ioutil.ReadAll(file)
// 	var user models.User
// 	err = json.Unmarshal(body, &user)
// 	if err != nil {
// 		helpers.LogMsg("Init parse user error: ", err)
// 		return nil, err
// 	}
// 	return &user, nil
// }

// func TestSignIn(t *testing.T) {
// 	var user models.User
// 	user.Login = time.Now().Format("20060102150405") + user.Login
// 	user.Email = time.Now().Format("20060102150405") + user.Email
// 	user.Password = "password"
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}
// 	defer db.Close()
// 	dab.SetMock(db)
// 	rows := sqlmock.NewRows([]string{"id", "login", "email", "hashpassword", "score", "name", "games"}).
// 		AddRow(1, "login", "login@mail.ru", "$2a$08$Q1nN3cy96NhOW7jOx31atuzY.QuRXbnWRitfkwZDHbC3dY83bw53i", 10, "name", "5").
// 		RowError(1, fmt.Errorf("error"))
// 	mock.ExpectQuery("SELECT").WillReturnRows(rows)
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

// 	rows = sqlmock.NewRows([]string{"id", "login", "email", "hashpassword", "score", "name", "games"}).
// 		AddRow(1, "login", "login@mail.ru", "$2a$08$.QuRXbnWRitfkwZDHbC3dY83bw53i", 10, "name", "5").
// 		RowError(1, fmt.Errorf("error"))
// 	mock.ExpectQuery("SELECT").WillReturnRows(rows)
// 	data, err = json.Marshal(user)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	buf = bytes.NewBuffer(data)
// 	req, err = http.NewRequest("POST", "/login", buf)
// 	w = httptest.NewRecorder()
// 	handler = http.HandlerFunc(SignIn)
// 	handler.ServeHTTP(w, req)
// 	expectedStatus = http.StatusForbidden
// 	if w.Code != expectedStatus {
// 		t.Error(w.Code)
// 	}

// 	data, err = json.Marshal(user)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	buf = bytes.NewBuffer(data)
// 	req, err = http.NewRequest("POST", "/login", buf)
// 	w = httptest.NewRecorder()
// 	handler = http.HandlerFunc(SignIn)
// 	handler.ServeHTTP(w, req)
// 	expectedStatus = http.StatusNotFound
// 	if w.Code != expectedStatus {
// 		t.Error(w.Code)
// 	}

// 	data, err = json.Marshal("")
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	buf = bytes.NewBuffer(data)
// 	req, err = http.NewRequest("POST", "/login", buf)
// 	w = httptest.NewRecorder()
// 	handler = http.HandlerFunc(SignIn)
// 	handler.ServeHTTP(w, req)
// 	expectedStatus = http.StatusInternalServerError
// 	if w.Code != expectedStatus {
// 		t.Error(w.Code)
// 	}
// }

// func TestSignUp(t *testing.T) {
// 	var user models.User
// 	user.Login = time.Now().Format("20060102150405") + user.Login
// 	user.Email = time.Now().Format("20060102150405") + user.Email
// 	user.Password = "password"
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}
// 	defer db.Close()
// 	dab.SetMock(db)
// 	rows := sqlmock.NewRows([]string{"id", "login", "email", "hashpassword", "score", "name", "games"}).
// 		AddRow(1, "login", "login@mail.ru", "$2a$08$Q1nN3cy96NhOW7jOx31atuzY.QuRXbnWRitfkwZDHbC3dY83bw53i", 10, "name", "5").
// 		RowError(1, fmt.Errorf("error"))
// 	mock.ExpectQuery("SELECT").WillReturnRows(rows)
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

// 	data, err = json.Marshal("")
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	buf = bytes.NewBuffer(data)
// 	req, err = http.NewRequest("POST", "/login", buf)
// 	w = httptest.NewRecorder()
// 	handler = http.HandlerFunc(SignUp)
// 	handler.ServeHTTP(w, req)
// 	expectedStatus = http.StatusInternalServerError
// 	if w.Code != expectedStatus {
// 		t.Error(w.Code)
// 	}

// 	rows = sqlmock.NewRows([]string{"id", "login", "score"}).
// 		AddRow(1, "login", 10).
// 		RowError(1, fmt.Errorf("error"))
// 	mock.ExpectQuery("SELECT").WillReturnError(fmt.Errorf("some error"))
// 	mock.ExpectQuery("INSERT").WillReturnRows(rows)
// 	data, err = json.Marshal(user)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	buf = bytes.NewBuffer(data)
// 	req, err = http.NewRequest("POST", "/signup", buf)
// 	w = httptest.NewRecorder()
// 	handler = http.HandlerFunc(SignUp)
// 	handler.ServeHTTP(w, req)
// 	expectedStatus = http.StatusOK
// 	if w.Code != expectedStatus {
// 		t.Error(w.Code)
// 	}

// 	rows = sqlmock.NewRows([]string{"id", "login", "score"}).
// 		AddRow(1, "login", 10).
// 		RowError(1, fmt.Errorf("error"))
// 	mock.ExpectQuery("SELECT").WillReturnError(fmt.Errorf("some error"))
// 	mock.ExpectQuery("INSERT").WillReturnError(fmt.Errorf("some error"))
// 	data, err = json.Marshal(user)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	buf = bytes.NewBuffer(data)
// 	req, err = http.NewRequest("POST", "/signup", buf)
// 	w = httptest.NewRecorder()
// 	handler = http.HandlerFunc(SignUp)
// 	handler.ServeHTTP(w, req)
// 	expectedStatus = http.StatusInternalServerError
// 	if w.Code != expectedStatus {
// 		t.Error(w.Code)
// 	}
// 	req, err = http.NewRequest("GET", "/", nil)
// 	handler = http.HandlerFunc(RootHandler)
// 	handler.ServeHTTP(w, req)

// }

// func TestMe(t *testing.T) {
// 	var user models.User
// 	user.Login = time.Now().Format("20060102150405") + user.Login
// 	user.Email = time.Now().Format("20060102150405") + user.Email
// 	user.Password = "password"
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}
// 	defer db.Close()
// 	dab.SetMock(db)
// 	rows := sqlmock.NewRows([]string{"id", "login", "score"}).
// 		AddRow(1, "login", 10).
// 		RowError(1, fmt.Errorf("error"))
// 	mock.ExpectQuery("SELECT").WillReturnError(fmt.Errorf("some error"))
// 	mock.ExpectQuery("INSERT").WillReturnRows(rows)
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
// 	req, err = http.NewRequest("POST", "/me", buf)
// 	w = httptest.NewRecorder()
// 	handler = http.HandlerFunc(Me)
// 	req.AddCookie(cookie)
// 	mock.ExpectQuery("SELECT").WillReturnError(fmt.Errorf("some error"))

// 	handler.ServeHTTP(w, req)

// 	expectedStatus = http.StatusNotFound
// 	if w.Code != expectedStatus {
// 		t.Error(w.Code)
// 	}
// 	rows = sqlmock.NewRows([]string{"id", "login", "email", "hashpassword", "score", "name", "games"}).
// 		AddRow(1, "login", "login@mail.ru", "hdfbkfbdj", 10, "name", "5").
// 		RowError(1, fmt.Errorf("error"))
// 	mock.ExpectQuery("SELECT").WillReturnRows(rows)
// 	handler.ServeHTTP(w, req)
// 	expectedStatus = http.StatusNotFound
// 	if w.Code != expectedStatus {
// 		t.Error(w.Code)
// 	}

// 	rows = sqlmock.NewRows([]string{"games", "name", "score"}).
// 		AddRow(10, "name", 5).
// 		RowError(1, fmt.Errorf("error"))
// 	mock.ExpectQuery("UPDATE").WillReturnError(fmt.Errorf("some error"))
// 	data, err = json.Marshal(user)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	buf = bytes.NewBuffer(data)
// 	req, err = http.NewRequest("POST", "/update", buf)
// 	w = httptest.NewRecorder()
// 	handler = http.HandlerFunc(ChangeProfile)
// 	req.AddCookie(cookie)
// 	handler.ServeHTTP(w, req)
// 	expectedStatus = http.StatusConflict
// 	if w.Code != expectedStatus {
// 		t.Error(w.Code)
// 	}

// 	rows = sqlmock.NewRows([]string{"games", "name", "score"}).
// 		AddRow(10, "name", 5).
// 		RowError(1, fmt.Errorf("error"))
// 	mock.ExpectQuery("UPDATE").WillReturnRows(rows)
// 	data, err = json.Marshal(user)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	buf = bytes.NewBuffer(data)
// 	req, err = http.NewRequest("POST", "/update", buf)
// 	w = httptest.NewRecorder()
// 	handler = http.HandlerFunc(ChangeProfile)
// 	req.AddCookie(cookie)
// 	handler.ServeHTTP(w, req)
// 	expectedStatus = http.StatusOK
// 	if w.Code != expectedStatus {
// 		t.Error(w.Code)
// 	}

// 	req, err = http.NewRequest("GET", "/signout", nil)
// 	w = httptest.NewRecorder()
// 	handler = http.HandlerFunc(SignOut)
// 	req.AddCookie(cookie)

// 	handler.ServeHTTP(w, req)
// 	expectedStatus = http.StatusOK
// 	if w.Code != expectedStatus {
// 		t.Error(w.Code)
// 	}

// 	newReq, err := http.NewRequest("GET", "/signout", nil)
// 	w = httptest.NewRecorder()
// 	handler = http.HandlerFunc(SignOut)

// 	handler.ServeHTTP(w, newReq)
// 	expectedStatus = http.StatusNotFound
// 	if w.Code != expectedStatus {
// 		t.Error(w.Code)
// 	}
// }

// func TestUpdateImageAndLeaders(t *testing.T) {
// 	var user models.User
// 	user.Login = time.Now().Format("20060102150405") + user.Login
// 	user.Email = time.Now().Format("20060102150405") + user.Email
// 	user.Password = "password"
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}
// 	defer db.Close()
// 	dab.SetMock(db)
// 	rows := sqlmock.NewRows([]string{"id", "login", "score"}).
// 		AddRow(1, "login", 10).
// 		RowError(1, fmt.Errorf("error"))
// 	mock.ExpectQuery("SELECT").WillReturnError(fmt.Errorf("some error"))
// 	mock.ExpectQuery("INSERT").WillReturnRows(rows)
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
// 	req, err = http.NewRequest("POST", "/upload", buf)
// 	w = httptest.NewRecorder()
// 	handler = http.HandlerFunc(UploadImage)
// 	req.AddCookie(cookie)
// 	mock.ExpectQuery("SELECT").WillReturnError(fmt.Errorf("some error"))

// 	handler.ServeHTTP(w, req)

// 	expectedStatus = http.StatusNotFound
// 	if w.Code != expectedStatus {
// 		t.Error(w.Code)
// 	}

// 	req, _ = http.NewRequest("GET", "/leaders", nil)
// 	handler = http.HandlerFunc(GetLeaderboardPage)
// 	handler.ServeHTTP(w, req)
// 	expectedStatus = http.StatusNotFound
// 	if w.Code != expectedStatus {
// 		t.Error(w.Code)
// 	}

// 	rows = sqlmock.NewRows([]string{"login", "score"}).
// 		AddRow("login", 10).
// 		RowError(1, fmt.Errorf("error"))
// 	mock.ExpectBegin()
// 	mock.ExpectQuery("SELECT").WillReturnRows(rows)
// 	mock.ExpectCommit()
// 	handler = http.HandlerFunc(GetLeaderboardPage)
// 	w = httptest.NewRecorder()
// 	handler.ServeHTTP(w, req)
// 	expectedStatus = http.StatusOK
// 	if w.Code != expectedStatus {
// 		t.Error(w.Code)
// 	}
// }
