package middleware

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"

	"2019_1_undefined_penguins/internal/pkg/helpers"
)

var SECRET = []byte("myawesomesecret")

func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")

		responseHeader := w.Header()
		responseHeader.Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		responseHeader.Set("Access-Control-Allow-Credentials", "true")
		responseHeader.Set("Access-Control-Allow-Headers", "Content-Type")
		responseHeader.Set("Access-Control-Allow-Origin", origin)

		if r.Method == "OPTIONS" {
			return
		}

		next.ServeHTTP(w, r)
	})
}

func PanicMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				helpers.LogMsg("Recovered panic: ", err)
				http.Error(w, "Server error", 500)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.RequestURI != "/me" {
			next.ServeHTTP(w, r)
			return
		}
		cookie, err := r.Cookie("sessionid")

		if err != nil || cookie.Value == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				w.WriteHeader(http.StatusForbidden)
				helpers.DeleteCookie(&w, cookie)
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return SECRET, nil
		})

		if _, ok := token.Claims.(jwt.MapClaims); !(ok && token.Valid) {
			w.WriteHeader(http.StatusForbidden)
			helpers.DeleteCookie(&w, cookie)
			return
		}
		next.ServeHTTP(w, r)
	})
}
