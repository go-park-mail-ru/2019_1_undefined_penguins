package controllers

import (
	"2019_1_undefined_penguins/internal/pkg/helpers"
	//"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"golang.org/x/net/context"
	"google.golang.org/grpc/status"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"time"

	"2019_1_undefined_penguins/internal/pkg/models"
	"github.com/aws/aws-sdk-go/aws"
)

func Me(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("sessionid")
	if cookie == nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	ctx := context.Background()
	user, err := models.AuthManager.GetUser(ctx, &models.JWT{Token: cookie.Value})

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

	easyJsonUser := models.ToEasyJsonUser(user)
	//bytes, err := json.Marshal(user)
	bytes, err := easyJsonUser.MarshalJSON()
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
	easyJsonUser := models.ToEasyJsonUser(newUser)
	//err = json.Unmarshal(body, &easyJsonUser)
	err = easyJsonUser.UnmarshalJSON(body)

	if err != nil {
		helpers.LogMsg(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	newUser = models.ToModelUser(easyJsonUser)
	newUser.ID = user.ID
	_, _ = models.AuthManager.ChangeUser(ctx, newUser)
	easyJsonUser = models.ToEasyJsonUser(newUser)
	//bytes, err := json.Marshal(newUser)
	bytes, err := easyJsonUser.MarshalJSON()
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

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState:session.SharedConfigEnable,
		Config: aws.Config{
			Region:aws.String("ru-msk"),
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

	_, err = svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(fileName),
		Body:   f,
	})
	if err != nil {
		helpers.LogMsg("Ошибка при сохранении файла: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	params := &s3.PutObjectAclInput{
		Bucket:&bucket,
		ACL:&acl,
		Key:&key,
	}

	_, err = svc.PutObjectAcl(params)
	if err != nil {
		helpers.LogMsg("Ошибка при добавлении прав доступа: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	newUser := new(models.User)
	newUser = user
	newUser.Picture = "https://hb.bizmrg.com/penguins_images/" + fileName
	newUser.ID = user.ID
	easyJsonUser := models.ToEasyJsonUser(newUser)
	_, _ = models.AuthManager.ChangeUserPicture(ctx, newUser)
	//bytes, err := json.Marshal(newUser)
	bytes, err := easyJsonUser.MarshalJSON()
	w.Write(bytes)
}
