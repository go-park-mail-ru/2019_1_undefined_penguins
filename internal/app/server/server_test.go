package server

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestServerFail(t *testing.T) {

	params := Params{Port: os.Getenv("PORT")}
	if params.Port == "" {
		params.Port = "8080"
	}
	err := StartApp(params)
	if err == nil {
		t.Fatal(err)
	}

}

func TestHomepage(t *testing.T) {

	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(RootHandler)

	handler.ServeHTTP(w, req)
	expectedStatus := http.StatusOK
	if w.Code != expectedStatus {
		t.Error(w.Code)
	}

}
