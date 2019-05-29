package middleware

import (
	"2019_1_undefined_penguins/internal/app/metrics"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"

	"2019_1_undefined_penguins/internal/pkg/helpers"
)

type Status struct {
	http.ResponseWriter
	Code int
}

var SECRET = []byte("myawesomesecret")

func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//origin := r.Header.Get("Origin")

		responseHeader := w.Header()
		responseHeader.Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		responseHeader.Set("Access-Control-Allow-Credentials", "true")
		responseHeader.Set("Access-Control-Allow-Headers", "Content-Type")
		responseHeader.Set("Access-Control-Allow-Origin", "https://penguin-wars.sytes.pro")

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

func MonitoringMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		status := GetStatus(w)
		next.ServeHTTP(status, r)

		metrics.Hits.WithLabelValues(
			strconv.FormatInt(int64(status.Code), 10),
			r.URL.String()).Inc()
	})
}

func GetStatus(responseWriter http.ResponseWriter) *Status {
	return &Status{ResponseWriter: responseWriter}
}

func (w *Status) WriteHeader(status int) {
	w.Code = status
	w.ResponseWriter.WriteHeader(status)
}

func (w *Status) Write(data []byte) (int, error) {
	if w.Code == 0 {
		w.Code = 200
	}

	n, err := w.ResponseWriter.Write(data)
	return n, err
}

func (w *Status) Header() http.Header {
	return w.ResponseWriter.Header()
}
