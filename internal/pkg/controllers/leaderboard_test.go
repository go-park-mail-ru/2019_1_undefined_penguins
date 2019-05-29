package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"google.golang.org/grpc"
	"2019_1_undefined_penguins/internal/pkg/models"

)

func TestLeaderboard(t *testing.T) {

	grcpConn, err := grpc.Dial(
		//authAddress,
		"127.0.0.1:8083",
		grpc.WithInsecure(),
	)
	if err != nil {
		t.Error(err)
	}
	defer grcpConn.Close()

	models.AuthManager = models.NewAuthCheckerClient(grcpConn)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/leaders/1", nil)
	handler := http.HandlerFunc(GetLeaderboardPage)
	handler.ServeHTTP(w, req)
}

func TestLeaderboardInfo(t *testing.T) {

	grcpConn, err := grpc.Dial(
		//authAddress,
		"127.0.0.1:8083",
		grpc.WithInsecure(),
	)
	if err != nil {
		t.Error(err)
	}
	defer grcpConn.Close()

	models.AuthManager = models.NewAuthCheckerClient(grcpConn)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	handler := http.HandlerFunc(GetLeaderboardInfo)
	handler.ServeHTTP(w, req)
}
