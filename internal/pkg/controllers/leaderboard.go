package controllers

import (
	db "2019_1_undefined_penguins/internal/pkg/database"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func GetLeaderboardPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		id = 1
	}

	users, _ := db.GetLeaders(id)
  //TODO: fix 404!!!!
	fmt.Println("led: ", users)
	if len(users) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if respBody, err := json.Marshal(users); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write(respBody)
	}
}

func GetLeaderboardInfo(w http.ResponseWriter, r *http.Request) {
	info, err := db.UsersCount()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if respBody, err := json.Marshal(info); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write(respBody)
	}
}
