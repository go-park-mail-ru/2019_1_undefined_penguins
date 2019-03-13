package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetLeaders(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		setupResponse(&w, r)
		w.WriteHeader(200)
		return
	} else {
		fmt.Print(r)
		bytes, err := json.Marshal(users)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"error":"=("}`))
		}
		setupResponse(&w, r)
		w.Write(bytes)
		fmt.Println(bytes)

	}
}
