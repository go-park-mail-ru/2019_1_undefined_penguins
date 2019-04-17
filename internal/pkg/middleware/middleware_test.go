package middleware

import (
	"net/http"
	"testing"
)

func TestCORS(t *testing.T) {
	_ = CORSMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if w.Header().Get("Access-Control-Allow-Headers") != "Content-Type" {
			t.Error("CORS are not working")
		}
	}))
}

func TestAuthMiddleware(t *testing.T) {
	_ = AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	}))
}

func TestLoggingMiddleware(t *testing.T) {
	_ = LoggingMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	}))
}

func TestPanicMiddleware(t *testing.T) {
	_ = PanicMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	}))
}
