package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	db "2019_1_undefined_penguins/internal/pkg/database"
	"2019_1_undefined_penguins/internal/pkg/models"

	"github.com/gorilla/mux"
)

func GetLeaders(w http.ResponseWriter, r *http.Request) {

	if r.Method == "OPTIONS" {
		w.WriteHeader(200)
		return
	} else {
		users := []models.User{}
		users, _ = db.GetLeaders(1)
		if len(users) == 0 {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		bytes, err := json.Marshal(users)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"error":"=("}`))
		}
		w.Write(bytes)
	}

}

func GetLeaderboardPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	users, _ := db.GetLeaders(id)
	if len(users) == 0 {

		w.WriteHeader(http.StatusNotFound)
		return
	}
	if respBody, err := json.Marshal(users); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.Write(respBody)
	}
}
