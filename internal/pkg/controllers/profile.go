package controllers

import (
	//"2019_1_undefined_penguins/internal/app/auth"
	"2019_1_undefined_penguins/internal/pkg/helpers"
	"encoding/json"
	"fmt"
	//"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"golang.org/x/net/context"
	"google.golang.org/grpc/status"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"time"

	//db "2019_1_undefined_penguins/internal/pkg/database"
	"2019_1_undefined_penguins/internal/pkg/models"
	//"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/aws"
	//"github.com/dgrijalva/jwt-go"
)

func Me(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("sessionid")

	ctx := context.Background()
	user, err := models.AuthManager.GetUser(ctx, &models.JWT{Token: cookie.Value})

	fmt.Println(err)
	if err != nil {
		switch errGRPC, _ := status.FromError(err); errGRPC.Code() {
		// case 2:
		// 	w.WriteHeader(http.StatusUnauthorized)
		// 	return
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

	ctx := context.Background()
	user, err := models.AuthManager.GetUser(ctx, &models.JWT{Token: cookie.Value})

	fmt.Println(err)

	if err != nil {
		switch errGRPC, _ := status.FromError(err); errGRPC.Code() {
		// case 2:
		// 	w.WriteHeader(http.StatusUnauthorized)
		// 	return
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
	_, _ = models.AuthManager.ChangeUser(ctx, newUser)
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

	ctx := context.Background()
	user, err := models.AuthManager.GetUser(ctx, &models.JWT{Token: cookie.Value})

	fmt.Println(err)
	if err != nil {
		switch errGRPC, _ := status.FromError(err); errGRPC.Code() {
		// case 2:
		// 	w.WriteHeader(http.StatusUnauthorized)
		// 	return
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

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState:session.SharedConfigEnable,
		Config: aws.Config{
			Endpoint:aws.String("http://hb.bizmrg.com"),
		},
	}))
	svc := s3.New(sess)

	f, handler, err := r.FormFile("avatar")
	if err != nil {
		helpers.LogMsg("Ошибка при получении файла из тела запроса")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer f.Close()

	extension := filepath.Ext(handler.Filename)
	t := time.Now()
	fileName := user.Login + t.Format("20060102150405") + extension

	bucket := "penguins_images"
	acl := "public-read"
	key := fileName

	res, err := svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(fileName),
		Body:   f,
	})
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(res)
	}

	params := &s3.PutObjectAclInput{
		Bucket:&bucket,
		ACL:&acl,
		Key:&key,
	}

	_, err = svc.PutObjectAcl(params)
	if err != nil {
		fmt.Println(err)
	}

	newUser := new(models.User)
	newUser = user
	newUser.Picture = "https://hb.bizmrg.com/penguins_images/" + fileName
	newUser.ID = user.ID
	//_, errr := models.AuthManager.ChangeUser(ctx, newUser)
	//fmt.Println(errr)
	bytes, err := json.Marshal(newUser)
	w.Write(bytes)

	// check object grands
	//input := &s3.GetObjectAclInput{
	//	Bucket: aws.String("penguins_images"),
	//	Key: aws.String("user620190521134436.png"),
	//}
	//
	//result1, err := svc.GetObjectAcl(input)
	//if err != nil {
	//	if aerr, ok := err.(awserr.Error); ok {
	//		switch aerr.Code() {
	//		case s3.ErrCodeNoSuchKey:
	//			fmt.Println(s3.ErrCodeNoSuchKey, aerr.Error())
	//		default:
	//			fmt.Println(aerr.Error())
	//		}
	//	} else {
	//		// Print the error, cast err to awserr.Error to get the Code and
	//		// Message from an error.
	//		fmt.Println(err.Error())
	//	}
	//	return
	//}
	//
	//fmt.Println(result1)


	// saving picture in fileserver
	// file, handler, err := r.FormFile("avatar")
	// if err != nil {
	// 	helpers.LogMsg("Ошибка при получении файла из тела запроса")
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }
	// defer file.Close()
	// extension := filepath.Ext(handler.Filename)
	// if extension == "" {
	// 	helpers.LogMsg("Файл не имеет расширения")

	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }

	// t := time.Now()

	// fileName := user.Login + t.Format("20060102150405") + extension
	// fileAndPath := "static/" + fileName
	// saveFile, err := os.Create(fileAndPath)
	// if err != nil {
	// 	helpers.LogMsg("Create", err)
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }
	// defer saveFile.Close()

	// _, err = io.Copy(saveFile, file)
	// if err != nil {
	// 	helpers.LogMsg("Copy", err)
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }

	// // u := user
	// // err = database.UpdateImage(u.Login, fileName)
	// if err != nil {
	// 	helpers.LogMsg("Ошибка при обновлении картинки в базе данных")

	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }
	// return
}
