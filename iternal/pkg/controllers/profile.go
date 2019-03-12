package controllers

import (
	"encoding/json"
	"net/http"
)

func Me(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("sessionid")
	if err != nil {
		//УТОЧНИТЬ У ФРОНТА КАКОЙ СТАТУС
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("You are not authorized"))
		return
	}
	user, found := sessions[cookie.Value]
	if !found {
		//УТОЧНИТЬ У ФРОНТА КАКОЙ СТАТУС!
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("You are not authorized"))
		return
	}
	bytes, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error"))
	}
	w.Write(bytes)

}
