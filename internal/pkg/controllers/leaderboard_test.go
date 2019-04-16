package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestLeaderboardWrongPage(t *testing.T) {
	r := mux.NewRouter()
	ts := httptest.NewServer(r)
	defer ts.Close()
	url := "http://localhost:8080/leaders/-11"
	resp, err := http.Get(url)
	if err != nil {
		t.Fatal(err)
	}

	if status := resp.StatusCode; status != http.StatusNotFound {
		t.Fatal(status)
	}
}

// func TestLeaderboardOKPage(t *testing.T) {

// 	r := mux.NewRouter()
// 	ts := httptest.NewServer(r)
// 	defer ts.Close()
// 	url := "http://localhost:8080/leaders/1"
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if status := resp.StatusCode; status != http.StatusOK {
// 		t.Fatal(status)
// 	}
// }
