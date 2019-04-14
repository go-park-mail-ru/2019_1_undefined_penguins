package controllers

import (
	"2019_1_undefined_penguins/internal/pkg/helpers"
	"2019_1_undefined_penguins/internal/pkg/models"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func GetUserFromJSON(fileName string) (*models.User, error) {
	dir, err := os.Getwd()
	if err != nil {
		helpers.LogMsg("Getting directory error: ", err)
		return nil, err

	}

	file, err := os.Open(dir + "/configs/" + fileName)
	if err != nil {
		helpers.LogMsg("Open directory error: ", err)
		return nil, err
	}

	body, _ := ioutil.ReadAll(file)
	var user models.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		helpers.LogMsg("Init parse user error: ", err)
		return nil, err
	}
	return &user, nil
}

func TestUserCreate(t *testing.T) {
	user, err := GetUserFromJSON("testuser.json")
	if err != nil {
		t.Fatal(err)
	}
	data, err := json.Marshal(user)
	if err != nil {
		t.Fatal(err)
	}
	buf := bytes.NewBuffer(data)
	req, err := http.NewRequest("POST", "/user/create", buf)
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(SignIn)
	handler.ServeHTTP(w, req)
	expectedStatus := http.StatusCreated
	if w.Code != expectedStatus {
		t.Error(w.Code)
	}
	if w.HeaderMap["Set-Cookies"] != nil {
		t.Error("Response without cookies")
	}
}
