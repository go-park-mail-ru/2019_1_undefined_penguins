package middleware

import (
	"net/http"

	_ "2019_1_undefined_penguins/internal/pkg/controllers"
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
