package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLeaderboard(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/leaders/1", nil)
	handler := http.HandlerFunc(GetLeaderboardPage)
	handler.ServeHTTP(w, req)
}

func TestLeaderboardInfo(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	handler := http.HandlerFunc(GetLeaderboardInfo)
	handler.ServeHTTP(w, req)
}
