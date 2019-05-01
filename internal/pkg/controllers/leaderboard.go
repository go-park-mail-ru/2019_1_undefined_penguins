package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	db "2019_1_undefined_penguins/internal/pkg/database"

	"github.com/gorilla/mux"
)

func GetLeaderboardPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {

		w.WriteHeader(http.StatusNotFound)
		return
	}

	users, _ := db.GetLeaders(id)
	//if len(users) == 0 {
	//	w.WriteHeader(http.StatusNotFound)
	//	return
	//}
	if respBody, err := json.Marshal(users); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.Write(respBody)
	}
}
