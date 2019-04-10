package controllers

import (
	"encoding/json"
	"fmt"

	//"fmt"
	"io/ioutil"
	"net/http"

	db "2019_1_undefined_penguins/internal/pkg/database"
	//"2019_1_undefined_penguins/internal/pkg/helpers"
	"2019_1_undefined_penguins/internal/pkg/models"
	"github.com/jackc/pgx"

	jwt "github.com/dgrijalva/jwt-go"
)

func Me(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		return
	}
	cookie, err := r.Cookie("sessionid")
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	//email, _ := models.Sessions[cookie.Value]
	//if !found {
	//	w.WriteHeader(http.StatusForbidden)
	//	return
	//}


	token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return SECRET, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		temp := token.Claims.(jwt.MapClaims)
		fmt.Println(temp)
		user, err := db.GetUserByEmail(claims["username"].(string))
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
		return
	}
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte("not authorized"))
	fmt.Println(err)




	}

func ChangeProfile(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		return
	}
	//cookie, err := r.Cookie("sessionid")
	//if err != nil {
	//	w.WriteHeader(http.StatusForbidden)
	//	return
	//}
	//email, found := models.Sessions[cookie.Value]
	//if !found {
	//	w.WriteHeader(http.StatusForbidden)
	//	return
	//}

	cookie, err := r.Cookie("sessionid")
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	//email, _ := models.Sessions[cookie.Value]
	//if !found {
	//	w.WriteHeader(http.StatusForbidden)
	//	return
	//}


	token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return SECRET, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		temp := token.Claims.(jwt.MapClaims)
		fmt.Println(temp)
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

		err = db.UpdateUser(&user, claims["username"].(string))
		if err != nil {
			switch errPgx := err.(pgx.PgError); errPgx.Code  {
			case "23505":
				w.WriteHeader(http.StatusConflict)
				return
			default:
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
		//models.Sessions[cookie.Value] = user.Email
		w.Write(body)

	}
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte("not authorized"))
	fmt.Println(err)
}
