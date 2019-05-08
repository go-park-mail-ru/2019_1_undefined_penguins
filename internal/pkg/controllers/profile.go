package controllers

import (
	//"2019_1_undefined_penguins/internal/app/auth"
	"2019_1_undefined_penguins/internal/pkg/helpers"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"

	"net/http"

	//db "2019_1_undefined_penguins/internal/pkg/database"
	"2019_1_undefined_penguins/internal/pkg/models"
	//"github.com/dgrijalva/jwt-go"
)

func Me(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("sessionid")

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

	authManager := models.NewAuthCheckerClient(grcpConn)
	ctx := context.Background()

	user, err := authManager.GetUser(ctx, &models.JWT{Token: cookie.Value})

	fmt.Println(err)
	if err != nil {
		switch errGRPC, _ := status.FromError(err); errGRPC.Code() {
		case 2:
			w.WriteHeader(http.StatusUnauthorized)
			return
		default:
			helpers.LogMsg("Unknown gprc error")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	bytes, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(bytes)
	return
}

func ChangeProfile(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("sessionid")

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
	authManager := models.NewAuthCheckerClient(grcpConn)
	ctx := context.Background()

	user, err := authManager.GetUser(ctx, &models.JWT{Token: cookie.Value})

	fmt.Println(err)

	if err != nil {
		switch errGRPC, _ := status.FromError(err); errGRPC.Code() {
		case 2:
			w.WriteHeader(http.StatusUnauthorized)
			return
		default:
			helpers.LogMsg("Unknown gprc error")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		helpers.LogMsg(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	//var user models.User
	var newUser *models.User
	err = json.Unmarshal(body, &newUser)

	if err != nil {
		helpers.LogMsg(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	newUser.ID = user.ID
	_, _ = authManager.ChangeUser(ctx, newUser)
	bytes, err := json.Marshal(newUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(bytes)
	return
}

func UploadImage(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("sessionid")

	//token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
	//	return SECRET, nil
	//})
	//
	//claims, _ := token.Claims.(jwt.MapClaims)
	//user, err := db.GetUserByEmail(claims["userEmail"].(string))
	//if user == nil {
	//	w.WriteHeader(http.StatusNotFound)
	//	return
	//}

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

	authManager := models.NewAuthCheckerClient(grcpConn)
	ctx := context.Background()

	user, err := authManager.GetUser(ctx, &models.JWT{Token: cookie.Value})

	fmt.Println(err)
	if err != nil {
		switch errGRPC, _ := status.FromError(err); errGRPC.Code() {
		case 2:
			w.WriteHeader(http.StatusUnauthorized)
			return
		default:
			helpers.LogMsg("Unknown gprc error")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	err = r.ParseMultipartForm(5 * 1024 * 1025)
	if err != nil {
		helpers.LogMsg("Ошибка при парсинге тела запроса")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	file, handler, err := r.FormFile("avatar")
	if err != nil {
		helpers.LogMsg("Ошибка при получении файла из тела запроса")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer file.Close()
	extension := filepath.Ext(handler.Filename)
	if extension == "" {
		helpers.LogMsg("Файл не имеет расширения")

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	t := time.Now()

	fileName := user.Login + t.Format("20060102150405") + extension
	fileAndPath := "static/" + fileName
	saveFile, err := os.Create(fileAndPath)
	if err != nil {
		helpers.LogMsg("Create", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer saveFile.Close()

	_, err = io.Copy(saveFile, file)
	if err != nil {
		helpers.LogMsg("Copy", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// u := user
	// err = database.UpdateImage(u.Login, fileName)
	if err != nil {
		helpers.LogMsg("Ошибка при обновлении картинки в базе данных")

		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	return
}
