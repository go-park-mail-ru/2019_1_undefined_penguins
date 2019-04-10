package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	//"fmt"
	"io/ioutil"
	"net/http"

	//"2019_1_undefined_penguins/internal/pkg/helpers"
	"2019_1_undefined_penguins/internal/pkg/database"
	"2019_1_undefined_penguins/internal/pkg/models"

	db "2019_1_undefined_penguins/internal/pkg/database"

	"github.com/jackc/pgx"
)

func Me(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		return
	}
	cookie, err := r.Cookie("sessionid")
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	email, found := models.Sessions[cookie.Value]
	if !found {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	user, err := db.GetUserByEmail(email)
	if user == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	bytes, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(bytes)
}

func ChangeProfile(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		return
	}
	cookie, err := r.Cookie("sessionid")
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	email, found := models.Sessions[cookie.Value]
	if !found {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	var user models.User
	err = json.Unmarshal(body, &user)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = db.UpdateUser(&user, email)
	if err != nil {
		switch errPgx := err.(pgx.PgError); errPgx.Code {
		case "23505":
			w.WriteHeader(http.StatusConflict)
			return
		default:
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	models.Sessions[cookie.Value] = user.Email
	w.Write(body)
}

var uploadFormTmpl = []byte(`
<html>
	<body>
	<form action="/upload" method="post" enctype="multipart/form-data">
		Image: <input type="file" name="avatar">
		<input type="submit" value="Upload">
	</form>
	</body>
</html>
`)

func UploadPage(w http.ResponseWriter, r *http.Request) {

	w.Write(uploadFormTmpl)

}

func UploadImage(w http.ResponseWriter, r *http.Request) {

	// cookie, err := r.Cookie("sessionid")
	// if err != nil {
	// 	w.WriteHeader(http.StatusForbidden)
	// 	return
	// }
	// email, found := models.Sessions[cookie.Value]
	// if !found {
	// 	w.WriteHeader(http.StatusForbidden)
	// 	return
	// }
	// user, err := db.GetUserByEmail(email)
	// if user == nil {
	// 	w.WriteHeader(http.StatusNotFound)
	// 	return
	// }

	var user models.User
	user.Login = "iamfrommoscow"

	err := r.ParseMultipartForm(5 * 1024 * 1025)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	file, handler, err := r.FormFile("avatar")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer file.Close()
	extension := filepath.Ext(handler.Filename)
	if extension == "" {
		fmt.Println(err)

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	t := time.Now()

	fileName := user.Login + t.Format("20060102150405") + extension
	fileAndPath := "static/" + fileName
	saveFile, err := os.Create(fileAndPath)
	if err != nil {
		fmt.Println("Create", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer saveFile.Close()

	_, err = io.Copy(saveFile, file)
	if err != nil {
		fmt.Println("Copy", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = database.UpdateImage(user.Login, fileName)
	if err != nil {
		fmt.Println(err)

		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
