package middleware

import (
	"fmt"
	//"log"
	"net/http"

	_ "2019_1_undefined_penguins/internal/pkg/controllers"

	"2019_1_undefined_penguins/internal/pkg/helpers"
)

//TODO check
func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")

		responseHeader := w.Header()
		responseHeader.Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		responseHeader.Set("Access-Control-Allow-Credentials", "true")
		responseHeader.Set("Access-Control-Allow-Headers", "Content-Type")
		responseHeader.Set("Access-Control-Allow-Origin", origin)

		next.ServeHTTP(w, r)
	})
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//log.Println(r.Method + r.RequestURI)
		helpers.LogMsg(r.Method + r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func PanicMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
				helpers.LogMsg("Recovered panic: ", err)
				http.Error(w, "Server error", 500)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

//TODO check
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
		next.ServeHTTP(w, r)
	})
}
