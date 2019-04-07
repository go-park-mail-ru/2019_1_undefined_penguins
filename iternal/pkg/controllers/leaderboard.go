package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-park-mail-ru/2019_1_undefined_penguins/iternal/pkg/helpers"

	"github.com/gorilla/mux"
)

func GetLeaders(w http.ResponseWriter, r *http.Request) {

	// if r.Method == "OPTIONS" {
	// 	SetupCORS(&w, r)
	// 	w.WriteHeader(200)
	// 	return
	// } else {
	// 	fmt.Print(r)
	// 	bytes, err := json.Marshal(users)
	// 	if err != nil {
	// 		w.WriteHeader(http.StatusInternalServerError)
	// 		w.Write([]byte(`{"error":"=("}`))
	// 	}
	// 	SetupCORS(&w, r)
	// 	w.Write(bytes)
	// 	fmt.Println(bytes)

	// }
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
	users := helpers.GetLeaders(id)
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
