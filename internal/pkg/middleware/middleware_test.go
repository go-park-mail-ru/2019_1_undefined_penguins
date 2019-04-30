package middleware

import (
	"net/http"
	"testing"
	"time"

	"github.com/gorilla/mux"
)

func RootTestHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello penguins"))
}

func PanicTestHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello penguins"))
	router := mux.NewRouter()
	err := http.ListenAndServe(":8085", router)
	panic(err)
}

func TestMid(t *testing.T) {
	router := mux.NewRouter()

	router.Use(PanicMiddleware)
	router.Use(CORSMiddleware)
	router.Use(AuthMiddleware)
	router.HandleFunc("/", RootTestHandler).Methods("GET")
	router.HandleFunc("/me", RootTestHandler).Methods("GET", "OPTIONS")
	router.HandleFunc("/panic", PanicTestHandler).Methods("GET")
	go func() {
		http.ListenAndServe(":8085", router)
	}()
	_, err := http.Get("http://127.0.0.1:8085/")
	if err != nil {
		t.Error(err)
	}
	_, err = http.Get("http://127.0.0.1:8085/panic")
	if err != nil {
		t.Error(err)
	}
	_, err = http.Get("http://127.0.0.1:8085/me")
	if err != nil {
		t.Error(err)
	}
	req, err := http.NewRequest("GET", "http://127.0.0.1:8085/me", nil)
	cookie := http.Cookie{Name: "sessionid", Value: "cookie_value"}
	req.AddCookie(&cookie)
	client := &http.Client{Timeout: time.Second * 10}
	_, err = client.Do(req)

	req, err = http.NewRequest("GET", "http://127.0.0.1:8085/me", nil)
	cookie = http.Cookie{Name: "sessionid", Value: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"}
	req.AddCookie(&cookie)
	_, err = client.Do(req)

	req, err = http.NewRequest("OPTIONS", "http://127.0.0.1:8085/me", nil)
	_, err = client.Do(req)
}
